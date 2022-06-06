module github.com/yuhu-tech/qilin-sdk-go

go 1.16

replace google.golang.org/protobuf => google.golang.org/protobuf v1.25.0

replace github.com/golang/protobuf => github.com/golang/protobuf v1.3.3

replace google.golang.org/grpc => google.golang.org/grpc v1.29.1

replace google.golang.org/genproto => google.golang.org/genproto v0.0.0-20191009194640-548a555dbc03

require (
	github.com/go-playground/form/v4 v4.2.0
	github.com/golang/protobuf v1.5.2
	github.com/google/go-cmp v0.5.7 // indirect
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.7.1
	golang.org/x/sys v0.0.0-20220412211240-33da011f77ad // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013
	google.golang.org/grpc v1.46.2
	google.golang.org/protobuf v1.28.0
	gopkg.in/yaml.v3 v3.0.1
)
