// nolint:gomnd
package gerr

type IDomain interface {
	GetDomain() string
}

// 通用的错误
// 满足下面的场景就可以使用通用的错误
// - 如果可以通过 http code 确定错误（reason、format, args... 只是用于显示错误信息），那么无需在 proto 里自定义错误

// TODO 每个方法前面加New方便提示
// BadRequest new BadRequest error that is mapped to a 400 response.
func InvalidArgument(d IDomain, format string, args ...interface{}) *Error {
	return Newf(400, d.GetDomain()+"_InvalidArgument", format, args...)
}

// IsBadRequest determines if err is an error which indicates a BadRequest error.
// It supports wrapped errors.
func IsInvalidArgument(err error) bool {
	return Code(err) == 400
}

// Unauthorized new Unauthorized error that is mapped to a 401 response.
func Unauthorized(d IDomain, format string, args ...interface{}) *Error {
	return Newf(401, d.GetDomain()+"_Unauthorized", format, args...)
}

// IsUnauthorized determines if err is an error which indicates a Unauthorized error.
// It supports wrapped errors.
func IsUnauthorized(err error) bool {
	return Code(err) == 401
}

// Forbidden new Forbidden error that is mapped to a 403 response.
func PermissionDenied(d IDomain, format string, args ...interface{}) *Error {
	return Newf(403, d.GetDomain()+"_PermissionDenied", format, args...)
}

// IsForbidden determines if err is an error which indicates a Forbidden error.
// It supports wrapped errors.
func IsPermissionDenied(err error) bool {
	return Code(err) == 403
}

// NotFound new NotFound error that is mapped to a 404 response.
func NotFound(d IDomain, format string, args ...interface{}) *Error {
	return Newf(404, d.GetDomain()+"_NotFound", format, args...)
}

// IsNotFound determines if err is an error which indicates an NotFound error.
// It supports wrapped errors.
func IsNotFound(err error) bool {
	return Code(err) == 404
}

// Conflict new Conflict error that is mapped to a 409 response.
func AlreadyExists(d IDomain, format string, args ...interface{}) *Error {
	return Newf(409, d.GetDomain()+"_AlreadyExists", format, args...)
}

// IsConflict determines if err is an error which indicates a Conflict error.
// It supports wrapped errors.
func IsAlreadyExists(err error) bool {
	return Code(err) == 409
}

// InternalServer new InternalServer error that is mapped to a 500 response.
func Internal(d IDomain, format string, args ...interface{}) *Error {
	return Newf(500, d.GetDomain()+"_Internal", format, args...)
}

// IsInternalServer determines if err is an error which indicates an Internal error.
// It supports wrapped errors.
func IsInternal(err error) bool {
	return Code(err) == 500
}

// ServiceUnavailable new ServiceUnavailable error that is mapped to a HTTP 503 response.
func ServiceUnavailable(d IDomain, format string, args ...interface{}) *Error {
	return Newf(503, d.GetDomain()+"_ServiceUnavailable", format, args...)
}

// IsServiceUnavailable determines if err is an error which indicates a Unavailable error.
// It supports wrapped errors.
func IsServiceUnavailable(err error) bool {
	return Code(err) == 503
}

// GatewayTimeout new GatewayTimeout error that is mapped to a HTTP 504 response.
func GatewayTimeout(d IDomain, format string, args ...interface{}) *Error {
	return Newf(504, d.GetDomain()+"_GatewayTimeout", format, args...)
}

// IsGatewayTimeout determines if err is an error which indicates a GatewayTimeout error.
// It supports wrapped errors.
func IsGatewayTimeout(err error) bool {
	return Code(err) == 504
}

// ClientClosed new ClientClosed error that is mapped to a HTTP 499 response.
func ClientClosed(d IDomain, format string, args ...interface{}) *Error {
	return Newf(499, d.GetDomain()+"_ClientClosed", format, args...)
}

// IsClientClosed determines if err is an error which indicates a IsClientClosed error.
// It supports wrapped errors.
func IsClientClosed(err error) bool {
	return Code(err) == 499
}
