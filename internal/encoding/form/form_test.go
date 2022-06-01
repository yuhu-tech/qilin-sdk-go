package form

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yuhu-tech/qilin-sdk-go/internal/encoding"
)

type LoginRequest struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

type TestModel struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

const contentType = "x-www-form-urlencoded"

func TestFormCodecMarshal(t *testing.T) {
	req := &LoginRequest{
		Username: "kratos",
		Password: "kratos_pwd",
	}
	content, err := encoding.GetCodec(contentType).Marshal(req)
	require.NoError(t, err)
	require.Equal(t, []byte("password=kratos_pwd&username=kratos"), content)

	req = &LoginRequest{
		Username: "kratos",
		Password: "",
	}
	content, err = encoding.GetCodec(contentType).Marshal(req)
	require.NoError(t, err)
	require.Equal(t, []byte("username=kratos"), content)

	m := &TestModel{
		ID:   1,
		Name: "kratos",
	}
	content, err = encoding.GetCodec(contentType).Marshal(m)
	t.Log(string(content))
	require.NoError(t, err)
	require.Equal(t, []byte("id=1&name=kratos"), content)
}

func TestFormCodecUnmarshal(t *testing.T) {
	req := &LoginRequest{
		Username: "kratos",
		Password: "kratos_pwd",
	}
	content, err := encoding.GetCodec(contentType).Marshal(req)
	require.NoError(t, err)

	bindReq := new(LoginRequest)
	err = encoding.GetCodec(contentType).Unmarshal(content, bindReq)
	require.NoError(t, err)
	require.Equal(t, "kratos", bindReq.Username)
	require.Equal(t, "kratos_pwd", bindReq.Password)
}

func TestProtoEncodeDecode(t *testing.T) {
	// in := &complex.Complex{
	// 	Id:      2233,
	// 	NoOne:   "2233",
	// 	Simple:  &complex.Simple{Component: "5566"},
	// 	Simples: []string{"3344", "5566"},
	// 	B:       true,
	// 	Sex:     complex.Sex_woman,
	// 	Age:     18,
	// 	A:       19,
	// 	Count:   3,
	// 	Price:   11.23,
	// 	D:       22.22,
	// 	Byte:    []byte("123"),
	// 	Map:     map[string]string{"kratos": "https://go-kratos.dev/"},

	// 	Timestamp: &timestamppb.Timestamp{Seconds: 20, Nanos: 2},
	// 	Duration:  &durationpb.Duration{Seconds: 120, Nanos: 22},
	// 	Field:     &fieldmaskpb.FieldMask{Paths: []string{"1", "2"}},
	// 	Double:    &wrapperspb.DoubleValue{Value: 12.33},
	// 	Float:     &wrapperspb.FloatValue{Value: 12.34},
	// 	Int64:     &wrapperspb.Int64Value{Value: 64},
	// 	Int32:     &wrapperspb.Int32Value{Value: 32},
	// 	Uint64:    &wrapperspb.UInt64Value{Value: 64},
	// 	Uint32:    &wrapperspb.UInt32Value{Value: 32},
	// 	Bool:      &wrapperspb.BoolValue{Value: false},
	// 	String_:   &wrapperspb.StringValue{Value: "go-kratos"},
	// 	Bytes:     &wrapperspb.BytesValue{Value: []byte("123")},
	// }
	// content, err := encoding.GetCodec(contentType).Marshal(in)
	// require.NoError(t, err)
	// require.Equal(t, "a=19&age=18&b=true&bool=false&byte=MTIz&bytes=MTIz&count=3&d=22.22&double=12.33&duration="+
	// 	"2m0.000000022s&field=1%2C2&float=12.34&id=2233&int32=32&int64=64&map%5Bkratos%5D=https%3A%2F%2Fgo-kratos.dev%2F&"+
	// 	"numberOne=2233&price=11.23&sex=woman&simples=3344&simples=5566&string=go-kratos"+
	// 	"&timestamp=1970-01-01T00%3A00%3A20.000000002Z&uint32=32&uint64=64&very_simple.component=5566", string(content))
	// in2 := &complex.Complex{}
	// err = encoding.GetCodec(contentType).Unmarshal(content, in2)
	// require.NoError(t, err)
	// require.Equal(t, int64(2233), in2.Id)
	// require.Equal(t, "2233", in2.NoOne)
	// require.NotEmpty(t, in2.Simple)
	// require.Equal(t, "5566", in2.Simple.Component)
	// require.NotEmpty(t, in2.Simples)
	// require.Len(t, in2.Simples, 2)
	// require.Equal(t, "3344", in2.Simples[0])
	// require.Equal(t, "5566", in2.Simples[1])
}
