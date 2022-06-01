package gerr

import "testing"

type domain string

func (d domain) GetDomain() string {
	return string(d)
}

func TestTypes(t *testing.T) {

	var (
		input = []error{
			InvalidArgument(domain("domain"), "reason_400"),
			Unauthorized(domain("domain"), "reason_401"),
			PermissionDenied(domain("domain"), "reason_403"),
			NotFound(domain("domain"), "reason_404"),
			AlreadyExists(domain("domain"), "reason_409"),
			Internal(domain("domain"), "reason_500"),
			ServiceUnavailable(domain("domain"), "reason_503"),
		}
		output = []func(error) bool{
			IsInvalidArgument,
			IsUnauthorized,
			IsPermissionDenied,
			IsNotFound,
			IsAlreadyExists,
			IsInternal,
			IsServiceUnavailable,
		}
	)

	for i, in := range input {
		if !output[i](in) {
			t.Errorf("not expect: %v", in)
		}
	}
}
