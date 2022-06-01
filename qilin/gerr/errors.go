package gerr

import (
	"errors"
	"fmt"
	"strings"

	httpstatus "github.com/yuhu-tech/qilin-sdk-go/gerr/status"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/status"
)

const (
	// UnknownCode is unknown code for error info.
	UnknownCode = 500
	// UnknownReason is unknown reason for error info.
	UnknownReason = ""
	// SupportPackageIsVersion1 this constant should not be referenced by any other code.
	SupportPackageIsVersion1 = true
)

//go:generate protoc -I. --go_out=paths=source_relative:. errors.proto

type GRPCError interface {
	error
	GRPCStatus() *status.Status
}

var _ GRPCError = (*Error)(nil)

const (
	En            = "en"
	ZhCN          = "zh_cn"
	DefaultLocale = ZhCN
)

func New(code int, reason, message string) *Error {
	var err Error
	err.Reason = reason
	err.Code = int32(code)
	err.Message = message
	return &err
}

func Newf(code int, reason, format string, args ...interface{}) *Error {
	var err Error
	err.Reason = reason
	err.Code = int32(code)
	err.Message = fmt.Sprintf(format, args...)
	return &err
}

func (e *Error) Error() string {
	return fmt.Sprintf("error: code = %d reason = %s message = %s metadata = %v", e.Code, e.Reason, e.Message, e.Metadata)
}

// GRPCStatus returns the Status represented by se.
func (e *Error) GRPCStatus() *status.Status {
	domain := ""
	// reason = domain_model_errorinfo
	if reasonElements := strings.Split(e.Reason, "_"); len(reasonElements) > 1 {
		domain = reasonElements[0]
	}
	s, _ := status.New(httpstatus.ToGRPCCode(int(e.Code)), e.Message).
		WithDetails(&errdetails.ErrorInfo{
			Reason:   e.Reason,
			Domain:   domain,
			Metadata: e.Metadata,
		})
	return s
}

// Is matches each error in the chain with the target value.
func (e *Error) Is(err error) bool {
	if se := new(Error); errors.As(err, &se) {
		return se.Reason == e.Reason
	}
	return false
}

// WithMetadata with an MD formed by the mapping of key, value.
func (e *Error) WithMetadata(md map[string]string) *Error {
	err := New(int(e.Code), e.Reason, e.Message)
	err.Metadata = md
	return err
}

// Code returns the http code for a error.
// It supports wrapped errors.
func Code(err error) int {
	if err == nil {
		return 200 //nolint:gomnd
	}
	if se := FromError(err); se != nil {
		return int(se.Code)
	}
	return UnknownCode
}

// Reason returns the reason for a particular error.
// It supports wrapped errors.
func Reason(err error) string {
	if se := FromError(err); se != nil {
		return se.Reason
	}
	return UnknownReason
}

// FromError try to convert an error to *Error.
// It supports wrapped errors.
func FromError(err error) *Error {
	if err == nil {
		return nil
	}
	if se := new(Error); errors.As(err, &se) {
		return se
	}
	gs, ok := status.FromError(err)
	if ok {
		ret := New(
			httpstatus.FromGRPCCode(gs.Code()),
			UnknownReason,
			gs.Message(),
		)
		for _, detail := range gs.Details() {
			switch d := detail.(type) {
			case *errdetails.ErrorInfo:
				ret.Reason = d.Reason
				return ret.WithMetadata(d.Metadata)
			}
		}
		return ret
	}
	return New(UnknownCode, UnknownReason, err.Error())
}

func ClearErrorCause(err error) error {
	if e, ok := status.FromError(err); ok {
		details := e.Details()
		if len(details) > 0 {
			detail := details[0]
			if d, ok := detail.(*Error); ok {
				d.Message = ""
				// clear detail
				proto := e.Proto()
				proto.Details = proto.Details[:0]
				e = status.FromProto(proto)
				e, _ := e.WithDetails(d)
				return e.Err()
			}
		}
	}
	return err
}
