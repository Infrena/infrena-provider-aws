package ccprov

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"

	"github.com/aws/aws-sdk-go-v2/aws/retry"
	awshttp "github.com/aws/aws-sdk-go-v2/aws/transport/http"
	"github.com/aws/smithy-go"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"github.com/infrena/infrena-provider-aws/internal/catalog"
	"github.com/infrena/infrena/pkg/provider"
)

// HandlerError is a request Cloud Control accepted whose resource handler ended FAILED, or that was cancelled.
type HandlerError struct {
	Operation  string // CREATE, UPDATE or DELETE
	Code       string // a handler error code, e.g. NotStabilized
	Message    string
	Identifier string // set when AWS had assigned one
	Token      string // the request token, for `aws cloudcontrol get-resource-request-status`
}

func (e *HandlerError) Error() string {
	return fmt.Sprintf("%s request %s ended %s: %s", e.Operation, e.Token, e.Code, e.Message)
}

var (
	// Refused before acting.
	safeCodes = map[string]bool{"ConcurrentOperationException": true, "ResourceConflictException": true}
	// A server-side failure that may have acted.
	maybeCodes = map[string]bool{
		"HandlerInternalFailureException": true, "HandlerFailureException": true, "ServiceInternalErrorException": true,
		"NetworkFailureException": true, "NotStabilizedException": true, "ConcurrentModificationException": true,
	}
	// Handler outcomes worth another attempt, though the failed request may have changed something.
	maybeHandlerCodes = map[string]bool{
		"NotStabilized": true, "ServiceInternalError": true, "InternalFailure": true, "NetworkFailure": true,
		"ServiceTimeout": true, "Throttling": true, "ResourceConflict": true,
	}
)

// classify answers infrena's question, how dangerous is another attempt, from the error alone. It must be a pure
// function: the plugin SDK may ask any configured instance, not the one that failed.
//
// Order matters: a throttle is checked before the HTTP status, and a refused dial before the generic network error
// (a *url.Error is a net.Error).
func classify(err error) provider.Retryability {
	if err == nil || errors.Is(err, context.Canceled) {
		return provider.NotSafeToRetry
	}
	var he *HandlerError
	if errors.As(err, &he) {
		if maybeHandlerCodes[he.Code] {
			return provider.ConditionallyRetryable
		}
		return provider.NotSafeToRetry
	}
	var apiErr smithy.APIError
	if errors.As(err, &apiErr) {
		code := apiErr.ErrorCode()
		if _, throttled := retry.DefaultThrottleErrorCodes[code]; throttled || safeCodes[code] {
			return provider.SafeToRetry
		}
		if maybeCodes[code] {
			return provider.ConditionallyRetryable
		}
	}
	var opErr *net.OpError
	if errors.As(err, &opErr) && opErr.Op == "dial" {
		return provider.SafeToRetry // no connection, so nothing was sent
	}
	var respErr *smithyhttp.ResponseError
	if errors.As(err, &respErr) && respErr.HTTPStatusCode() >= 500 {
		return provider.ConditionallyRetryable
	}
	var netErr net.Error
	if errors.Is(err, context.DeadlineExceeded) || errors.Is(err, io.ErrUnexpectedEOF) || errors.Is(err, io.EOF) || errors.As(err, &netErr) {
		return provider.ConditionallyRetryable // sent, answer lost
	}
	return provider.NotSafeToRetry
}

// errorCode is the AWS error code or handler error code an error carries, or "".
func errorCode(err error) string {
	var he *HandlerError
	if errors.As(err, &he) {
		return he.Code
	}
	var apiErr smithy.APIError
	if errors.As(err, &apiErr) {
		return apiErr.ErrorCode()
	}
	return ""
}

func isNotFound(err error) bool {
	switch errorCode(err) {
	case "ResourceNotFoundException", "NotFound":
		return true
	}
	return false
}

// apiFailure is the message a user reads in a failed apply, with the cause kept underneath for classification.
type apiFailure struct {
	msg string
	err error
}

func (e *apiFailure) Error() string { return e.msg }
func (e *apiFailure) Unwrap() error { return e.err }

// failure says which instance, which call, which type, where, what AWS said and the ID AWS support asks for. Never
// anything from the request itself, which may hold a secret.
func failure(instance, action string, t *catalog.Type, where string, err error) error {
	msg := fmt.Sprintf("aws instance %q: Cloud Control %s of %s (%s) at %s failed: ", instance, action, t.Name, t.CFN, where)
	var he *HandlerError
	var apiErr smithy.APIError
	switch {
	case errors.As(err, &he):
		msg += he.Code + ": " + he.Message + " (request token " + he.Token + ")"
	case errors.As(err, &apiErr):
		msg += apiErr.ErrorCode() + ": " + apiErr.ErrorMessage()
	default:
		msg += err.Error()
	}
	var re *awshttp.ResponseError
	if errors.As(err, &re) && re.ServiceRequestID() != "" {
		msg += " (request ID " + re.ServiceRequestID() + ")"
	}
	switch errorCode(err) {
	case "AccessDenied", "AccessDeniedException", "UnauthorizedTaggingOperation":
		msg += fmt.Sprintf("\ncheck the IAM permissions of the credentials instance %q uses", instance)
	case "InvalidCredentials", "InvalidCredentialsException", "UnrecognizedClientException", "ExpiredTokenException":
		msg += fmt.Sprintf("\nthe credentials instance %q uses were refused: check `profile` or `assume_role_arn`", instance)
	case "TypeNotFoundException":
		msg += "\nCloud Control does not offer " + t.CFN + " in this region"
	case "AlreadyExists", "AlreadyExistsException":
		msg += "\nsomething with that name already exists: adopt it with `infrena import`, or choose another name"
	}
	return &apiFailure{msg: msg, err: err}
}
