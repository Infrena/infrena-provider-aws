package smithyinvoke

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"github.com/aws/smithy-go/transport/http/protocol/ec2query"
)

// Target is where and as whom to send a request.
type Target struct {
	Endpoint    string // e.g. https://ec2.us-east-1.amazonaws.com, or a fake's URL
	Region      string
	SigningName string // e.g. "ec2"
	Credentials aws.CredentialsProvider
	HTTP        *http.Client
}

// Invoke calls an operation by name with a generic input and returns a generic output. Nothing about the
// operation exists in Go before this call: its schema is built from the model now.
//
// Deliberately bare: no retries, no endpoint resolution, no error-type registry. Those are what a generated
// client's middleware stack adds, and what a production version would have to reassemble.
func Invoke(ctx context.Context, m *Model, t Target, operation string, in map[string]any) (map[string]any, error) {
	op, err := m.Operation(operation)
	if err != nil {
		return nil, err
	}
	if err := Validate(op.Input, in, operation); err != nil {
		return nil, err
	}

	proto := ec2query.New(m.Service())
	req := smithyhttp.NewStackRequest().(*smithyhttp.Request)
	u, err := url.Parse(t.Endpoint)
	if err != nil {
		return nil, err
	}
	req.URL = u
	if err := proto.SerializeRequest(ctx, op, input{schema: op.Input, value: in}, req); err != nil {
		return nil, fmt.Errorf("serialize: %w", err)
	}

	body, err := io.ReadAll(req.GetStream())
	if err != nil {
		return nil, err
	}
	if req, err = req.SetStream(bytes.NewReader(body)); err != nil {
		return nil, err
	}
	httpReq := req.Build(ctx)
	httpReq.ContentLength = int64(len(body))
	sum := sha256.Sum256(body)

	creds, err := t.Credentials.Retrieve(ctx)
	if err != nil {
		return nil, err
	}
	if err := v4.NewSigner().SignHTTP(ctx, creds, httpReq, hex.EncodeToString(sum[:]), t.SigningName, t.Region, time.Now()); err != nil {
		return nil, fmt.Errorf("sign: %w", err)
	}
	client := t.HTTP
	if client == nil {
		client = http.DefaultClient
	}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	out := &output{schema: op.Output}
	if err := proto.DeserializeResponse(ctx, op, &smithy.TypeRegistry{}, &smithyhttp.Response{Response: resp}, out); err != nil {
		return nil, err
	}
	return out.value, nil
}
