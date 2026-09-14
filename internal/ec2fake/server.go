// Package ec2fake is an in-process stand-in for the parts of EC2 and STS this plugin calls. It speaks
// the query protocol, so tests drive the real AWS SDK against it: serialisation, error decoding,
// retries and endpoint overrides are all the SDK's own.
package ec2fake

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
)

const (
	// Owner is the account every resource in the fake belongs to.
	Owner = "123456789012"
	// AssumedAccessKey is the access key ID AssumeRole hands out, so a test can tell calls signed with
	// an assumed role from calls signed with the static test keys.
	AssumedAccessKey = "ASIAFAKEASSUMED"
)

// VPC is a snapshot of one fake VPC.
type VPC struct {
	ID, Region, CIDR, Owner string
	Default                 bool
	Tags                    map[string]string
}

// Subnet is a snapshot of one fake subnet.
type Subnet struct {
	ID, Region, VPCID, CIDR, AZ, Owner string
	MapPublicIP                        bool
	Tags                               map[string]string
}

// Fault makes the Nth call of Action (counting from 1) fail. With Drop, the action takes effect and
// the connection is closed without a response: a request that was processed but whose answer was lost.
type Fault struct {
	Action  string
	Nth     int
	Status  int
	Code    string
	Message string
	Drop    bool
}

// Server is the fake. Resources are kept per region, taken from the SigV4 credential scope.
type Server struct {
	*httptest.Server
	// PageSize limits Describe* pages; 0 means everything in one page.
	PageSize int

	mu      sync.Mutex
	next    int
	vpcs    map[string]*VPC
	subnets map[string]*Subnet
	faults  []Fault
	hidden  map[string]int
	calls   map[string]int
	keys    []string
}

// New starts a fake. Close it when done.
func New() *Server {
	s := &Server{vpcs: map[string]*VPC{}, subnets: map[string]*Subnet{}, hidden: map[string]int{}, calls: map[string]int{}}
	s.Server = httptest.NewServer(http.HandlerFunc(s.serve))
	return s
}

var scope = regexp.MustCompile(`Credential=([^/]+)/[^/]+/([^/]+)/`)

func (s *Server) serve(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		writeError(w, 400, "MalformedQueryString", err.Error())
		return
	}
	action := r.Form.Get("Action")
	key, region := "", ""
	if m := scope.FindStringSubmatch(r.Header.Get("Authorization")); m != nil {
		key, region = m[1], m[2]
	}

	s.mu.Lock()
	defer s.mu.Unlock()
	s.calls[action]++
	s.keys = append(s.keys, key)
	fault, faulted := s.faultFor(action)
	if faulted && !fault.Drop {
		writeError(w, fault.Status, fault.Code, fault.Message)
		return
	}
	status, body := s.handle(action, region, r)
	if faulted && fault.Drop {
		if hj, ok := w.(http.Hijacker); ok {
			if conn, _, err := hj.Hijack(); err == nil {
				_ = conn.Close()
				return
			}
		}
	}
	if e, isErr := body.(errorResponse); isErr {
		writeError(w, status, e.Code, e.Message)
		return
	}
	writeXML(w, status, body)
}

func (s *Server) faultFor(action string) (Fault, bool) {
	for i, f := range s.faults {
		if f.Action == action && f.Nth == s.calls[action] {
			s.faults = append(s.faults[:i], s.faults[i+1:]...)
			return f, true
		}
	}
	return Fault{}, false
}

func notFound(code, id string) (int, any) {
	return 400, errorResponse{Code: code, Message: fmt.Sprintf("The ID '%s' does not exist", id)}
}

func ok(action string) returnResponse {
	return returnResponse{XMLName: xmlName(action + "Response"), Xmlns: ec2NS, RequestID: "fake-request", Return: true}
}

func (s *Server) handle(action, region string, r *http.Request) (int, any) {
	f := r.Form
	switch action {
	case "AssumeRole":
		var out assumeRoleResponse
		out.Result.Credentials.AccessKeyID = AssumedAccessKey
		out.Result.Credentials.SecretAccessKey = "assumed-secret"
		out.Result.Credentials.SessionToken = "assumed-token"
		out.Result.Credentials.Expiration = "2099-01-01T00:00:00Z"
		return 200, out

	case "CreateVpc":
		id := s.newID("vpc")
		v := &VPC{ID: id, Region: region, CIDR: f.Get("CidrBlock"), Owner: Owner, Tags: tagSpec(f)}
		s.vpcs[id] = v
		return 200, createVpcResponse{Xmlns: ec2NS, RequestID: "fake-request", Vpc: vpcXML(v)}

	case "DescribeVpcs":
		ids := indexed(f, "VpcId")
		var out []xmlVPC
		for _, id := range ids {
			v, found := s.vpcs[id]
			if !found || v.Region != region || s.hide(id) {
				return notFound("InvalidVpcID.NotFound", id)
			}
			out = append(out, vpcXML(v))
		}
		if len(ids) == 0 {
			for _, id := range sortedKeys(s.vpcs) {
				if v := s.vpcs[id]; v.Region == region && !s.hide(id) {
					out = append(out, vpcXML(v))
				}
			}
		}
		page, next := s.paginate(len(out), f.Get("NextToken"))
		return 200, describeVpcsResponse{Xmlns: ec2NS, RequestID: "fake-request", Vpcs: out[page[0]:page[1]], NextToken: next}

	case "DeleteVpc":
		id := f.Get("VpcId")
		v, found := s.vpcs[id]
		if !found || v.Region != region {
			return notFound("InvalidVpcID.NotFound", id)
		}
		for _, sub := range s.subnets {
			if sub.VPCID == id {
				return 400, errorResponse{Code: "DependencyViolation", Message: "The vpc '" + id + "' has dependencies and cannot be deleted."}
			}
		}
		delete(s.vpcs, id)
		return 200, ok(action)

	case "CreateSubnet":
		vpcID := f.Get("VpcId")
		if v, found := s.vpcs[vpcID]; !found || v.Region != region || s.hide(vpcID) {
			return notFound("InvalidVpcID.NotFound", vpcID)
		}
		id := s.newID("subnet")
		sub := &Subnet{ID: id, Region: region, VPCID: vpcID, CIDR: f.Get("CidrBlock"), AZ: f.Get("AvailabilityZone"), Owner: Owner, Tags: tagSpec(f)}
		s.subnets[id] = sub
		return 200, createSubnetResponse{Xmlns: ec2NS, RequestID: "fake-request", Subnet: subnetXML(sub)}

	case "DescribeSubnets":
		ids := indexed(f, "SubnetId")
		var out []xmlSubnet
		for _, id := range ids {
			sub, found := s.subnets[id]
			if !found || sub.Region != region || s.hide(id) {
				return notFound("InvalidSubnetID.NotFound", id)
			}
			out = append(out, subnetXML(sub))
		}
		if len(ids) == 0 {
			for _, id := range sortedKeys(s.subnets) {
				if sub := s.subnets[id]; sub.Region == region && !s.hide(id) {
					out = append(out, subnetXML(sub))
				}
			}
		}
		page, next := s.paginate(len(out), f.Get("NextToken"))
		return 200, describeSubnetsResponse{Xmlns: ec2NS, RequestID: "fake-request", Subnets: out[page[0]:page[1]], NextToken: next}

	case "DeleteSubnet":
		id := f.Get("SubnetId")
		if sub, found := s.subnets[id]; !found || sub.Region != region {
			return notFound("InvalidSubnetID.NotFound", id)
		}
		delete(s.subnets, id)
		return 200, ok(action)

	case "ModifySubnetAttribute":
		id := f.Get("SubnetId")
		sub, found := s.subnets[id]
		if !found || sub.Region != region || s.hide(id) {
			return notFound("InvalidSubnetID.NotFound", id)
		}
		if v := f.Get("MapPublicIpOnLaunch.Value"); v != "" {
			sub.MapPublicIP = v == "true"
		}
		return 200, ok(action)

	case "CreateTags", "DeleteTags":
		for _, id := range indexed(f, "ResourceId") {
			tags := s.tagsOf(id, region)
			if tags == nil {
				return notFound("InvalidID", id)
			}
			for i := 1; f.Has(fmt.Sprintf("Tag.%d.Key", i)); i++ {
				k := f.Get(fmt.Sprintf("Tag.%d.Key", i))
				if action == "CreateTags" {
					tags[k] = f.Get(fmt.Sprintf("Tag.%d.Value", i))
				} else {
					delete(tags, k)
				}
			}
		}
		return 200, ok(action)
	}
	return 400, errorResponse{Code: "InvalidAction", Message: "ec2fake does not implement " + action}
}

func (s *Server) newID(prefix string) string {
	s.next++
	return fmt.Sprintf("%s-%017x", prefix, s.next)
}

func (s *Server) hide(id string) bool {
	if s.hidden[id] > 0 {
		s.hidden[id]--
		return true
	}
	return false
}

// paginate returns the [start, end) window for this page and the token for the next.
func (s *Server) paginate(n int, token string) ([2]int, string) {
	start, _ := strconv.Atoi(token)
	if s.PageSize <= 0 || start+s.PageSize >= n {
		return [2]int{min(start, n), n}, ""
	}
	return [2]int{start, start + s.PageSize}, strconv.Itoa(start + s.PageSize)
}

func (s *Server) tagsOf(id, region string) map[string]string {
	if v, found := s.vpcs[id]; found && v.Region == region {
		if v.Tags == nil {
			v.Tags = map[string]string{}
		}
		return v.Tags
	}
	if sub, found := s.subnets[id]; found && sub.Region == region {
		if sub.Tags == nil {
			sub.Tags = map[string]string{}
		}
		return sub.Tags
	}
	return nil
}

// indexed reads a flattened query list: Name.1, Name.2, …
func indexed(f map[string][]string, name string) []string {
	var out []string
	for i := 1; ; i++ {
		v, found := f[fmt.Sprintf("%s.%d", name, i)]
		if !found {
			return out
		}
		out = append(out, v[0])
	}
}

// tagSpec reads TagSpecification.1.Tag.N.Key/Value.
func tagSpec(f map[string][]string) map[string]string {
	tags := map[string]string{}
	for i := 1; ; i++ {
		k, found := f[fmt.Sprintf("TagSpecification.1.Tag.%d.Key", i)]
		if !found {
			return tags
		}
		tags[k[0]] = strings.Join(f[fmt.Sprintf("TagSpecification.1.Tag.%d.Value", i)], "")
	}
}

func vpcXML(v *VPC) xmlVPC {
	return xmlVPC{VpcID: v.ID, CidrBlock: v.CIDR, OwnerID: v.Owner, State: "available", IsDefault: v.Default, Tags: tagsXML(v.Tags)}
}

func subnetXML(sub *Subnet) xmlSubnet {
	return xmlSubnet{
		SubnetID: sub.ID, VpcID: sub.VPCID, CidrBlock: sub.CIDR, AvailabilityZone: sub.AZ,
		MapPublicIPOnLaunch: sub.MapPublicIP, OwnerID: sub.Owner, State: "available",
		SubnetArn: fmt.Sprintf("arn:aws:ec2:%s:%s:subnet/%s", sub.Region, sub.Owner, sub.ID), Tags: tagsXML(sub.Tags),
	}
}

func tagsXML(tags map[string]string) []xmlTag {
	var out []xmlTag
	for _, k := range sortedKeys(tags) {
		out = append(out, xmlTag{Key: k, Value: tags[k]})
	}
	return out
}

func sortedKeys[V any](m map[string]V) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func copyTags(tags map[string]string) map[string]string {
	out := make(map[string]string, len(tags))
	for k, v := range tags {
		out[k] = v
	}
	return out
}

// ---- test controls ----

// Inject queues a fault.
func (s *Server) Inject(f Fault) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.faults = append(s.faults, f)
}

// HideFromDescribe makes the next `times` lookups of id miss, as a resource EC2 has not propagated does.
func (s *Server) HideFromDescribe(id string, times int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.hidden[id] = times
}

// Calls counts requests for an action, faults included.
func (s *Server) Calls(action string) int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.calls[action]
}

// AccessKeys lists the access key ID each request was signed with, in order.
func (s *Server) AccessKeys() []string {
	s.mu.Lock()
	defer s.mu.Unlock()
	return append([]string(nil), s.keys...)
}

// AddVPC creates a VPC behind infrata's back, as infrastructure that predates it would be.
func (s *Server) AddVPC(region, cidr string, tags map[string]string) string {
	s.mu.Lock()
	defer s.mu.Unlock()
	id := s.newID("vpc")
	s.vpcs[id] = &VPC{ID: id, Region: region, CIDR: cidr, Owner: Owner, Tags: copyTags(tags)}
	return id
}

// SetTag changes a tag behind infrata's back.
func (s *Server) SetTag(id, key, val string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if tags := s.tagsOf(id, s.regionOf(id)); tags != nil {
		tags[key] = val
	}
}

// SetCIDR changes a CIDR behind infrata's back — something AWS itself never allows, which is the point.
func (s *Server) SetCIDR(id, cidr string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if v, found := s.vpcs[id]; found {
		v.CIDR = cidr
	}
	if sub, found := s.subnets[id]; found {
		sub.CIDR = cidr
	}
}

// Remove deletes a resource behind infrata's back.
func (s *Server) Remove(id string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.vpcs, id)
	delete(s.subnets, id)
}

func (s *Server) regionOf(id string) string {
	if v, found := s.vpcs[id]; found {
		return v.Region
	}
	if sub, found := s.subnets[id]; found {
		return sub.Region
	}
	return ""
}

// VPCs snapshots every VPC, sorted by ID.
func (s *Server) VPCs() []VPC {
	s.mu.Lock()
	defer s.mu.Unlock()
	var out []VPC
	for _, id := range sortedKeys(s.vpcs) {
		v := *s.vpcs[id]
		v.Tags = copyTags(v.Tags)
		out = append(out, v)
	}
	return out
}

// Subnets snapshots every subnet, sorted by ID.
func (s *Server) Subnets() []Subnet {
	s.mu.Lock()
	defer s.mu.Unlock()
	var out []Subnet
	for _, id := range sortedKeys(s.subnets) {
		sub := *s.subnets[id]
		sub.Tags = copyTags(sub.Tags)
		out = append(out, sub)
	}
	return out
}
