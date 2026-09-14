package ec2fake

import (
	"encoding/xml"
	"net/http"
)

const ec2NS = "http://ec2.amazonaws.com/doc/2016-11-15/"

type xmlTag struct {
	Key   string `xml:"key"`
	Value string `xml:"value"`
}

type xmlVPC struct {
	VpcID     string   `xml:"vpcId"`
	CidrBlock string   `xml:"cidrBlock"`
	OwnerID   string   `xml:"ownerId"`
	State     string   `xml:"state"`
	IsDefault bool     `xml:"isDefault"`
	Tags      []xmlTag `xml:"tagSet>item"`
}

type xmlSubnet struct {
	SubnetID            string   `xml:"subnetId"`
	VpcID               string   `xml:"vpcId"`
	CidrBlock           string   `xml:"cidrBlock"`
	AvailabilityZone    string   `xml:"availabilityZone"`
	MapPublicIPOnLaunch bool     `xml:"mapPublicIpOnLaunch"`
	SubnetArn           string   `xml:"subnetArn"`
	OwnerID             string   `xml:"ownerId"`
	State               string   `xml:"state"`
	Tags                []xmlTag `xml:"tagSet>item"`
}

type createVpcResponse struct {
	XMLName   xml.Name `xml:"CreateVpcResponse"`
	Xmlns     string   `xml:"xmlns,attr"`
	RequestID string   `xml:"requestId"`
	Vpc       xmlVPC   `xml:"vpc"`
}

type describeVpcsResponse struct {
	XMLName   xml.Name `xml:"DescribeVpcsResponse"`
	Xmlns     string   `xml:"xmlns,attr"`
	RequestID string   `xml:"requestId"`
	Vpcs      []xmlVPC `xml:"vpcSet>item"`
	NextToken string   `xml:"nextToken,omitempty"`
}

type createSubnetResponse struct {
	XMLName   xml.Name  `xml:"CreateSubnetResponse"`
	Xmlns     string    `xml:"xmlns,attr"`
	RequestID string    `xml:"requestId"`
	Subnet    xmlSubnet `xml:"subnet"`
}

type describeSubnetsResponse struct {
	XMLName   xml.Name    `xml:"DescribeSubnetsResponse"`
	Xmlns     string      `xml:"xmlns,attr"`
	RequestID string      `xml:"requestId"`
	Subnets   []xmlSubnet `xml:"subnetSet>item"`
	NextToken string      `xml:"nextToken,omitempty"`
}

// returnResponse is every EC2 action here that answers only success.
type returnResponse struct {
	XMLName   xml.Name
	Xmlns     string `xml:"xmlns,attr"`
	RequestID string `xml:"requestId"`
	Return    bool   `xml:"return"`
}

type assumeRoleResponse struct {
	XMLName xml.Name `xml:"AssumeRoleResponse"`
	Result  struct {
		Credentials struct {
			AccessKeyID     string `xml:"AccessKeyId"`
			SecretAccessKey string `xml:"SecretAccessKey"`
			SessionToken    string `xml:"SessionToken"`
			Expiration      string `xml:"Expiration"`
		} `xml:"Credentials"`
	} `xml:"AssumeRoleResult"`
}

// errorResponse is EC2's query-protocol error body: Errors>Error>Code, Errors>Error>Message, RequestID
// (aws-sdk-go-v2 aws/protocol/ec2query/error_utils.go).
type errorResponse struct {
	XMLName   xml.Name `xml:"Response"`
	Code      string   `xml:"Errors>Error>Code"`
	Message   string   `xml:"Errors>Error>Message"`
	RequestID string   `xml:"RequestID"`
}

func xmlName(local string) xml.Name { return xml.Name{Local: local} }

func writeXML(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "text/xml;charset=UTF-8")
	w.WriteHeader(status)
	_, _ = w.Write([]byte(xml.Header))
	_ = xml.NewEncoder(w).Encode(body)
}

func writeError(w http.ResponseWriter, status int, code, message string) {
	writeXML(w, status, errorResponse{Code: code, Message: message, RequestID: "fake-request-err"})
}
