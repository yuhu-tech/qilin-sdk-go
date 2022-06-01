package status

import (
	"net/http"

	"google.golang.org/grpc/codes"
)

const (
	// ClientClosed is non-standard http status code,
	// which defined by nginx.
	// https://httpstatus.in/499/
	ClientClosed      = 499
	OK                = http.StatusOK
	InvalidArgument   = http.StatusBadRequest
	Unauthenticated   = http.StatusUnauthorized
	PermissionDenied  = http.StatusForbidden
	NotFound          = http.StatusNotFound
	ResourceExhausted = http.StatusTooManyRequests
	Aborted           = http.StatusConflict
	Internal          = http.StatusInternalServerError
	Unimplemented     = http.StatusNotImplemented
	Unavailable       = http.StatusServiceUnavailable
	DeadlineExceeded  = http.StatusGatewayTimeout
	AlreadyExists     = http.StatusConflict
)

type Converter interface {
	// ToGRPCCode converts an HTTP error code into the corresponding gRPC response status.
	ToGRPCCode(code int) codes.Code

	// FromGRPCCode converts a gRPC error code into the corresponding HTTP response status.
	FromGRPCCode(code codes.Code) int
}

type statusConverter struct{}

var DefaultConverter Converter = statusConverter{}

// ToGRPCCode converts a HTTP error code into the corresponding gRPC response status.
// See: https://github.com/googleapis/googleapis/blob/master/google/rpc/code.proto
func (c statusConverter) ToGRPCCode(code int) codes.Code {
	switch code {
	case OK:
		return codes.OK
	case InvalidArgument:
		return codes.InvalidArgument
	case Unauthenticated:
		return codes.Unauthenticated
	case PermissionDenied:
		return codes.PermissionDenied
	case NotFound:
		return codes.NotFound
	case Aborted:
		return codes.Aborted
	case ResourceExhausted:
		return codes.ResourceExhausted
	case Internal:
		return codes.Internal
	case Unimplemented:
		return codes.Unimplemented
	case Unavailable:
		return codes.Unavailable
	case DeadlineExceeded:
		return codes.DeadlineExceeded
	case ClientClosed:
		return codes.Canceled
	}
	return codes.Unknown
}

// FromGRPCCode converts a gRPC error code into the corresponding HTTP response status.
// See: https://github.com/googleapis/googleapis/blob/master/google/rpc/code.proto
func (c statusConverter) FromGRPCCode(code codes.Code) int {
	switch code {
	case codes.OK:
		return OK
	case codes.Canceled:
		return ClientClosed
	case codes.Unknown:
		return Internal
	case codes.InvalidArgument:
		return InvalidArgument
	case codes.DeadlineExceeded:
		return DeadlineExceeded
	case codes.NotFound:
		return NotFound
	case codes.AlreadyExists:
		return AlreadyExists
	case codes.PermissionDenied:
		return PermissionDenied
	case codes.Unauthenticated:
		return Unauthenticated
	case codes.ResourceExhausted:
		return ResourceExhausted
	case codes.FailedPrecondition:
		return InvalidArgument
	case codes.Aborted:
		return Aborted
	case codes.OutOfRange:
		return InvalidArgument
	case codes.Unimplemented:
		return Unimplemented
	case codes.Internal:
		return Internal
	case codes.Unavailable:
		return Unavailable
	case codes.DataLoss:
		return Internal
	}
	return Internal
}

// ToGRPCCode converts an HTTP error code into the corresponding gRPC response status.
func ToGRPCCode(code int) codes.Code {
	return DefaultConverter.ToGRPCCode(code)
}

// FromGRPCCode converts a gRPC error code into the corresponding HTTP response status.
func FromGRPCCode(code codes.Code) int {
	return DefaultConverter.FromGRPCCode(code)
}
