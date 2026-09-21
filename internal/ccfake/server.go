// Package ccfake is an in-process stand-in for AWS Cloud Control API (and STS AssumeRole). Tests drive the real SDK
// against it, so the SDK is the judge of the wire format.
package ccfake

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"regexp"
	"slices"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

// AssumedAccessKey is what AssumeRole hands out, so a test can tell assumed-role calls from static-key calls.
const AssumedAccessKey = "ASIAFAKEASSUMED"

// TypeConfig describes how the fake behaves for one resource type.
type TypeConfig struct {
	TypeName   string            // AWS::EC2::VPC
	Identifier string            // the property the fake assigns on create, e.g. VpcId
	IDPrefix   string            // e.g. "vpc-"
	ReadOnly   map[string]string // properties added on create; "{id}" is replaced by the identifier
	Defaults   map[string]any    // properties AWS picks when the desired state omits them
	CreateOnly []string          // patching these is refused with NotUpdatableException
	WriteOnly  []string          // accepted, never returned
	OnRead     func(props map[string]any)
}

// Fault makes the Nth call of Action fail with an HTTP error.
type Fault struct {
	Action  string
	Nth     int
	Status  int
	Code    string
	Message string
}

type failure struct {
	code, message string
	keep          bool
}

type request struct {
	token, operation, typeName, identifier, region string
	polls                                          int
	failCode, failMessage                          string
}

// Server is the fake.
type Server struct {
	*httptest.Server
	PollsToComplete int
	PageSize        int

	mu        sync.Mutex
	next      int
	types     map[string]TypeConfig
	resources map[string]map[string]any // key: region|type|identifier
	writeOnly map[string]map[string]any
	requests  map[string]*request
	byToken   map[string]string // client token -> request token
	faults    []Fault
	failNext  map[string]failure
	hidden    map[string]int
	calls     map[string]int
	tokens    map[string][]string
	keys      []string
	lastPatch []map[string]any

	// What EC2 says AWS itself owns. Cloud Control cannot answer it, so discovery asks EC2, and the fake answers
	// about the very resources it already holds: one cloud seen through two APIs.
	defaultVPCs    map[string]string          // region -> the VPC EC2 calls that region's default
	defaultSubnets map[string]map[string]bool // region -> subnets EC2 calls default for their availability zone
	ec2Filters     map[string][]string        // EC2 action -> the filters each call carried
}

// New starts a fake.
func New() *Server {
	s := &Server{
		PollsToComplete: 1, types: map[string]TypeConfig{}, resources: map[string]map[string]any{},
		writeOnly: map[string]map[string]any{}, requests: map[string]*request{}, byToken: map[string]string{},
		failNext: map[string]failure{}, hidden: map[string]int{}, calls: map[string]int{}, tokens: map[string][]string{},
		defaultVPCs: map[string]string{}, defaultSubnets: map[string]map[string]bool{}, ec2Filters: map[string][]string{},
	}
	s.Server = httptest.NewServer(http.HandlerFunc(s.serve))
	return s
}

var scope = regexp.MustCompile(`Credential=([^/]+)/[^/]+/([^/]+)/`)

func key(region, typeName, id string) string { return region + "|" + typeName + "|" + id }

func (s *Server) serve(w http.ResponseWriter, r *http.Request) {
	accessKey, region := "", ""
	if m := scope.FindStringSubmatch(r.Header.Get("Authorization")); m != nil {
		accessKey, region = m[1], m[2]
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	s.keys = append(s.keys, accessKey)

	// No X-Amz-Target means one of the form-encoded protocols: STS's, or EC2's query protocol.
	target := r.Header.Get("X-Amz-Target")
	if target == "" {
		_ = r.ParseForm()
		action := r.Form.Get("Action")
		s.calls[action]++
		if f, hit := s.fault(action); hit {
			writeEC2Error(w, f.Status, f.Code, f.Message)
			return
		}
		switch action {
		case "DescribeVpcs":
			s.describeVpcs(w, region, r.Form)
		case "DescribeSubnets":
			s.describeSubnets(w, region, r.Form)
		default:
			s.assumeRole(w)
		}
		return
	}
	action := strings.TrimPrefix(target, "CloudApiService.")
	s.calls[action]++
	if f, hit := s.fault(action); hit {
		writeError(w, f.Status, f.Code, f.Message)
		return
	}
	body, _ := io.ReadAll(r.Body)
	var in map[string]any
	_ = json.Unmarshal(body, &in)
	str := func(k string) string { v, _ := in[k].(string); return v }
	if tok := str("ClientToken"); tok != "" {
		s.tokens[action] = append(s.tokens[action], tok)
	}

	switch action {
	case "CreateResource":
		s.create(w, region, str("TypeName"), str("DesiredState"), str("ClientToken"))
	case "GetResource":
		s.get(w, region, str("TypeName"), str("Identifier"))
	case "UpdateResource":
		s.update(w, region, str("TypeName"), str("Identifier"), str("PatchDocument"), str("ClientToken"))
	case "DeleteResource":
		s.delete(w, region, str("TypeName"), str("Identifier"), str("ClientToken"))
	case "ListResources":
		s.list(w, region, str("TypeName"), str("NextToken"))
	case "GetResourceRequestStatus":
		s.status(w, str("RequestToken"))
	default:
		writeError(w, 400, "UnsupportedActionException", "ccfake does not implement "+action)
	}
}

// fault takes the queued fault for this call of action, if a test queued one.
func (s *Server) fault(action string) (Fault, bool) {
	for i, f := range s.faults {
		if f.Action == action && f.Nth == s.calls[action] {
			s.faults = append(s.faults[:i], s.faults[i+1:]...)
			return f, true
		}
	}
	return Fault{}, false
}

func (s *Server) typeConfig(w http.ResponseWriter, typeName string) (TypeConfig, bool) {
	tc, ok := s.types[typeName]
	if !ok {
		writeError(w, 404, "TypeNotFoundException", "The type "+typeName+" cannot be found")
	}
	return tc, ok
}

func (s *Server) newRequest(region, operation, typeName, identifier, clientToken string) (*request, bool) {
	if clientToken != "" {
		if tok, seen := s.byToken[operation+"|"+clientToken]; seen {
			return s.requests[tok], true
		}
	}
	s.next++
	req := &request{token: fmt.Sprintf("req-%d", s.next), operation: operation, typeName: typeName, identifier: identifier, region: region}
	if f, ok := s.failNext[operation]; ok {
		delete(s.failNext, operation)
		req.failCode, req.failMessage = f.code, f.message
	}
	s.requests[req.token] = req
	if clientToken != "" {
		s.byToken[operation+"|"+clientToken] = req.token
	}
	return req, false
}

func (s *Server) create(w http.ResponseWriter, region, typeName, desired, clientToken string) {
	tc, ok := s.typeConfig(w, typeName)
	if !ok {
		return
	}
	keep := s.failNext["CREATE"].keep
	req, repeat := s.newRequest(region, "CREATE", typeName, "", clientToken)
	if !repeat && (req.failCode == "" || keep) {
		var props map[string]any
		if err := json.Unmarshal([]byte(desired), &props); err != nil {
			writeError(w, 400, "InvalidRequestException", "DesiredState is not a JSON object")
			return
		}
		// A name the caller chose (a bucket or role name) is the identifier, as in AWS; otherwise the fake makes one.
		id, named := props[tc.Identifier].(string)
		if !named || id == "" {
			s.next++
			id = fmt.Sprintf("%s%017x", tc.IDPrefix, s.next)
		}
		if _, taken := s.resources[key(region, typeName, id)]; taken {
			req.failCode, req.failMessage = "AlreadyExists", id+" already exists"
			writeJSON(w, map[string]any{"ProgressEvent": s.event(req)})
			return
		}
		props[tc.Identifier] = id
		for p, v := range tc.Defaults {
			if _, set := props[p]; !set {
				props[p] = v
			}
		}
		for p, tmpl := range tc.ReadOnly {
			props[p] = strings.ReplaceAll(tmpl, "{id}", id)
		}
		wo := map[string]any{}
		for _, p := range tc.WriteOnly {
			if v, set := props[p]; set {
				wo[p] = v
				delete(props, p)
			}
		}
		k := key(region, typeName, id)
		s.resources[k], s.writeOnly[k] = props, wo
		req.identifier = id
	}
	writeJSON(w, map[string]any{"ProgressEvent": s.event(req)})
}

func (s *Server) get(w http.ResponseWriter, region, typeName, id string) {
	if _, ok := s.typeConfig(w, typeName); !ok {
		return
	}
	props, ok := s.resources[key(region, typeName, id)]
	if s.hidden[id] > 0 {
		s.hidden[id]--
		ok = false
	}
	if !ok {
		writeError(w, 404, "ResourceNotFoundException", typeName+" Handler returned status FAILED: "+id+" does not exist (HandlerErrorCode: NotFound)")
		return
	}
	out := cloneMap(props)
	if tc := s.types[typeName]; tc.OnRead != nil {
		tc.OnRead(out)
	}
	raw, _ := json.Marshal(out)
	writeJSON(w, map[string]any{"TypeName": typeName, "ResourceDescription": map[string]any{"Identifier": id, "Properties": string(raw)}})
}

func (s *Server) update(w http.ResponseWriter, region, typeName, id, patch, clientToken string) {
	tc, ok := s.typeConfig(w, typeName)
	if !ok {
		return
	}
	props, exists := s.resources[key(region, typeName, id)]
	if !exists {
		writeError(w, 404, "ResourceNotFoundException", id+" does not exist")
		return
	}
	var ops []map[string]any
	if err := json.Unmarshal([]byte(patch), &ops); err != nil {
		writeError(w, 400, "InvalidRequestException", "PatchDocument is not a JSON Patch array")
		return
	}
	for _, op := range ops {
		path, _ := op["path"].(string)
		name := strings.TrimPrefix(path, "/")
		if strings.Contains(name, "/") {
			writeError(w, 400, "InvalidRequestException", "ccfake supports top-level patch paths only: "+path)
			return
		}
		for _, co := range tc.CreateOnly {
			if co == name {
				writeError(w, 400, "NotUpdatableException", "Invalid patch update: createOnlyProperties ["+path+"] cannot be updated")
				return
			}
		}
	}
	req, repeat := s.newRequest(region, "UPDATE", typeName, id, clientToken)
	if !repeat && req.failCode == "" {
		s.lastPatch = ops
		for _, op := range ops {
			name := strings.TrimPrefix(op["path"].(string), "/")
			switch op["op"] {
			case "add", "replace":
				if isWriteOnly(tc, name) {
					s.writeOnly[key(region, typeName, id)][name] = op["value"]
				} else {
					props[name] = op["value"]
				}
			case "remove":
				delete(props, name)
			}
		}
	}
	writeJSON(w, map[string]any{"ProgressEvent": s.event(req)})
}

func (s *Server) delete(w http.ResponseWriter, region, typeName, id, clientToken string) {
	if _, ok := s.typeConfig(w, typeName); !ok {
		return
	}
	req, repeat := s.newRequest(region, "DELETE", typeName, id, clientToken)
	k := key(region, typeName, id)
	if _, exists := s.resources[k]; !exists && req.failCode == "" {
		req.failCode, req.failMessage = "NotFound", id+" does not exist"
	}
	if !repeat && req.failCode == "" {
		delete(s.resources, k)
		delete(s.writeOnly, k)
	}
	writeJSON(w, map[string]any{"ProgressEvent": s.event(req)})
}

func (s *Server) list(w http.ResponseWriter, region, typeName, token string) {
	tc, ok := s.typeConfig(w, typeName)
	if !ok {
		return
	}
	ids := s.identifiers(region, typeName)
	start, _ := strconv.Atoi(token)
	end, next := len(ids), ""
	if s.PageSize > 0 && start+s.PageSize < len(ids) {
		end, next = start+s.PageSize, strconv.Itoa(start+s.PageSize)
	}
	var descs []map[string]any
	for _, id := range ids[min(start, len(ids)):end] {
		raw, _ := json.Marshal(map[string]any{tc.Identifier: id}) // like the real API: often only the identifier
		descs = append(descs, map[string]any{"Identifier": id, "Properties": string(raw)})
	}
	out := map[string]any{"TypeName": typeName, "ResourceDescriptions": descs}
	if next != "" {
		out["NextToken"] = next
	}
	writeJSON(w, out)
}

func (s *Server) status(w http.ResponseWriter, token string) {
	req, ok := s.requests[token]
	if !ok {
		writeError(w, 404, "RequestTokenNotFoundException", "no request "+token)
		return
	}
	req.polls++
	writeJSON(w, map[string]any{"ProgressEvent": s.event(req)})
}

func (s *Server) event(req *request) map[string]any {
	ev := map[string]any{
		"TypeName": req.typeName, "RequestToken": req.token, "Operation": req.operation,
		"EventTime": float64(time.Now().UnixMilli()) / 1000, "OperationStatus": "IN_PROGRESS",
	}
	if req.identifier != "" {
		ev["Identifier"] = req.identifier
	}
	if req.polls >= s.PollsToComplete {
		ev["OperationStatus"] = "SUCCESS"
		if req.failCode != "" {
			ev["OperationStatus"], ev["ErrorCode"], ev["StatusMessage"] = "FAILED", req.failCode, req.failMessage
		}
	}
	return ev
}

func (s *Server) assumeRole(w http.ResponseWriter) {
	type creds struct {
		AccessKeyID     string `xml:"AccessKeyId"`
		SecretAccessKey string `xml:"SecretAccessKey"`
		SessionToken    string `xml:"SessionToken"`
		Expiration      string `xml:"Expiration"`
	}
	var out struct {
		XMLName xml.Name `xml:"AssumeRoleResponse"`
		Result  struct {
			Credentials creds `xml:"Credentials"`
		} `xml:"AssumeRoleResult"`
	}
	out.Result.Credentials = creds{AssumedAccessKey, "assumed-secret", "assumed-token", "2099-01-01T00:00:00Z"}
	w.Header().Set("Content-Type", "text/xml")
	_, _ = w.Write([]byte(xml.Header))
	_ = xml.NewEncoder(w).Encode(out)
}

// identifiers lists what the fake holds of a type in a region, in a stable order.
func (s *Server) identifiers(region, typeName string) []string {
	var ids []string
	prefix := region + "|" + typeName + "|"
	for k := range s.resources {
		if id, ok := strings.CutPrefix(k, prefix); ok {
			ids = append(ids, id)
		}
	}
	sort.Strings(ids)
	return ids
}

// ---- EC2, the one API discovery asks besides Cloud Control ----

// describeVpcs answers from the VPCs the fake already holds: a VPC is the default only if a test said so with
// SetDefaultVPC, so nothing here can flag a resource AWS did not make.
func (s *Server) describeVpcs(w http.ResponseWriter, region string, form url.Values) {
	want, err := parseEC2Filters(form, "is-default")
	if err != nil {
		writeEC2Error(w, 400, "InvalidParameterValue", err.Error())
		return
	}
	s.ec2Filters["DescribeVpcs"] = append(s.ec2Filters["DescribeVpcs"], want.String())
	var items []string
	for _, id := range s.identifiers(region, "AWS::EC2::VPC") {
		isDefault := s.defaultVPCs[region] == id
		if !want.matches("is-default", strconv.FormatBool(isDefault)) {
			continue
		}
		items = append(items, fmt.Sprintf("<item><vpcId>%s</vpcId><isDefault>%t</isDefault></item>", xmlText(id), isDefault))
	}
	writeXML(w, "DescribeVpcsResponse", "<vpcSet>"+strings.Join(items, "")+"</vpcSet>")
}

func (s *Server) describeSubnets(w http.ResponseWriter, region string, form url.Values) {
	want, err := parseEC2Filters(form, "default-for-az")
	if err != nil {
		writeEC2Error(w, 400, "InvalidParameterValue", err.Error())
		return
	}
	s.ec2Filters["DescribeSubnets"] = append(s.ec2Filters["DescribeSubnets"], want.String())
	var items []string
	for _, id := range s.identifiers(region, "AWS::EC2::Subnet") {
		isDefault := s.defaultSubnets[region][id]
		if !want.matches("default-for-az", strconv.FormatBool(isDefault)) {
			continue
		}
		items = append(items, fmt.Sprintf("<item><subnetId>%s</subnetId><defaultForAz>%t</defaultForAz></item>", xmlText(id), isDefault))
	}
	writeXML(w, "DescribeSubnetsResponse", "<subnetSet>"+strings.Join(items, "")+"</subnetSet>")
}

// ec2Filter is one Filter.N.Name and its Filter.N.Value.M list, as the SDK form-encodes them.
type ec2Filter map[string][]string

// parseEC2Filters reads the filters a request carried, REFUSING any the fake does not implement: quietly ignoring one
// would answer a question nobody asked, and a caller that forgot its filter would look like it worked.
func parseEC2Filters(form url.Values, supported ...string) (ec2Filter, error) {
	out := ec2Filter{}
	for key, vals := range form {
		n, ok := strings.CutPrefix(key, "Filter.")
		if !ok || len(vals) == 0 {
			continue
		}
		n, ok = strings.CutSuffix(n, ".Name")
		if !ok {
			continue
		}
		name := vals[0]
		if !slices.Contains(supported, name) {
			return nil, fmt.Errorf("ccfake does not implement the filter %q", name)
		}
		var values []string
		for i := 1; ; i++ {
			v := form.Get(fmt.Sprintf("Filter.%s.Value.%d", n, i))
			if v == "" {
				break
			}
			values = append(values, v)
		}
		out[name] = values
	}
	return out, nil
}

// matches reports whether a resource whose value for name is value survives the filters. A filter nobody sent
// excludes nothing.
func (f ec2Filter) matches(name, value string) bool {
	values, filtered := f[name]
	return !filtered || slices.Contains(values, value)
}

func (f ec2Filter) String() string {
	out := make([]string, 0, len(f))
	for name, values := range f {
		out = append(out, name+"="+strings.Join(values, ","))
	}
	sort.Strings(out)
	return strings.Join(out, " ")
}

func writeXML(w http.ResponseWriter, root, body string) {
	w.Header().Set("Content-Type", "text/xml")
	fmt.Fprintf(w, `%s<%s xmlns="http://ec2.amazonaws.com/doc/2016-11-15/"><requestId>fake-request</requestId>%s</%s>`,
		xml.Header, root, body, root)
}

// writeEC2Error is the shape aws/protocol/ec2query parses: Code and Message under Errors>Error, RequestID beside it.
func writeEC2Error(w http.ResponseWriter, status int, code, message string) {
	w.Header().Set("Content-Type", "text/xml")
	w.WriteHeader(status)
	fmt.Fprintf(w, `%s<Response><Errors><Error><Code>%s</Code><Message>%s</Message></Error></Errors><RequestID>fake-request-err</RequestID></Response>`,
		xml.Header, xmlText(code), xmlText(message))
}

func xmlText(s string) string {
	var buf bytes.Buffer
	_ = xml.EscapeText(&buf, []byte(s))
	return buf.String()
}

func writeJSON(w http.ResponseWriter, body any) {
	w.Header().Set("Content-Type", "application/x-amz-json-1.0")
	_ = json.NewEncoder(w).Encode(body)
}

func writeError(w http.ResponseWriter, status int, code, message string) {
	w.Header().Set("Content-Type", "application/x-amz-json-1.0")
	w.Header().Set("X-Amzn-ErrorType", code)
	w.Header().Set("X-Amzn-Requestid", "fake-request-err") // the SDK reads it into ServiceRequestID
	w.WriteHeader(status)
	// The SDK's schema-based deserializer matches a known error shape's member name exactly ("Message", not
	// "message"): the generic ProtocolErrorInfo fallback is case-insensitive, but a registered type like
	// InvalidRequestException is not. Verified against service/cloudcontrol@v1.38.0's schemas (AddMember("Message", ...)).
	_ = json.NewEncoder(w).Encode(map[string]any{"__type": code, "Message": message})
}

func isWriteOnly(tc TypeConfig, name string) bool {
	for _, p := range tc.WriteOnly {
		if p == name {
			return true
		}
	}
	return false
}

func cloneMap(m map[string]any) map[string]any {
	raw, _ := json.Marshal(m)
	var out map[string]any
	_ = json.Unmarshal(raw, &out)
	return out
}

// ---- test controls ----

// Register makes the fake serve a type.
func (s *Server) Register(tc TypeConfig) { s.mu.Lock(); defer s.mu.Unlock(); s.types[tc.TypeName] = tc }

// Inject queues an HTTP fault.
func (s *Server) Inject(f Fault) { s.mu.Lock(); defer s.mu.Unlock(); s.faults = append(s.faults, f) }

// FailNext makes the next async request of operation end FAILED with a handler error code.
func (s *Server) FailNext(operation, handlerCode, message string, keepResource bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.failNext[operation] = failure{code: handlerCode, message: message, keep: keepResource}
}

// HideFromGet makes the next `times` GetResource calls for identifier answer ResourceNotFoundException.
func (s *Server) HideFromGet(identifier string, times int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.hidden[identifier] = times
}

// Calls counts requests for an action.
func (s *Server) Calls(action string) int { s.mu.Lock(); defer s.mu.Unlock(); return s.calls[action] }

// Tokens lists the client tokens an action was called with.
func (s *Server) Tokens(action string) []string {
	s.mu.Lock()
	defer s.mu.Unlock()
	return append([]string(nil), s.tokens[action]...)
}

// AccessKeys lists the access key ID each request was signed with.
func (s *Server) AccessKeys() []string {
	s.mu.Lock()
	defer s.mu.Unlock()
	return append([]string(nil), s.keys...)
}

// Resource snapshots one resource's stored properties (write-only ones excluded).
func (s *Server) Resource(region, typeName, identifier string) (map[string]any, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	p, ok := s.resources[key(region, typeName, identifier)]
	if !ok {
		return nil, false
	}
	return cloneMap(p), true
}

// Put stores a resource behind infrena's back: pre-existing infrastructure, or drift.
func (s *Server) Put(region, typeName, identifier string, props map[string]any) {
	s.mu.Lock()
	defer s.mu.Unlock()
	k := key(region, typeName, identifier)
	s.resources[k] = cloneMap(props)
	if s.writeOnly[k] == nil {
		s.writeOnly[k] = map[string]any{}
	}
}

// SetDefaultVPC makes EC2 report identifier as the region's default VPC, as a real account has one per region.
func (s *Server) SetDefaultVPC(region, identifier string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.defaultVPCs[region] = identifier
}

// SetDefaultSubnet makes EC2 report identifier as the default subnet for its availability zone.
func (s *Server) SetDefaultSubnet(region, identifier string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.defaultSubnets[region] == nil {
		s.defaultSubnets[region] = map[string]bool{}
	}
	s.defaultSubnets[region][identifier] = true
}

// Filters lists the filters each call of an EC2 action carried, as "name=value,value".
func (s *Server) Filters(action string) []string {
	s.mu.Lock()
	defer s.mu.Unlock()
	return append([]string(nil), s.ec2Filters[action]...)
}

// Resources snapshots every resource of a type in a region.
func (s *Server) Resources(region, typeName string) map[string]map[string]any {
	s.mu.Lock()
	defer s.mu.Unlock()
	out := map[string]map[string]any{}
	prefix := region + "|" + typeName + "|"
	for k, p := range s.resources {
		if id, ok := strings.CutPrefix(k, prefix); ok {
			out[id] = cloneMap(p)
		}
	}
	return out
}

// LastPatch is the most recent patch document applied.
func (s *Server) LastPatch() []map[string]any { s.mu.Lock(); defer s.mu.Unlock(); return s.lastPatch }
