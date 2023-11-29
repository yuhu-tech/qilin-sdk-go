package evidence

import (
	"context"
	"errors"
	"fmt"
	"strings"

	qhttp "github.com/yuhu-tech/qilin-sdk-go/qilin/transport/http"
)

const ServiceName = "evidence"
const ServiceAPIVersion = "2022-06-02"

type Client struct {
	cc       *qhttp.Client
	tenantId string
}

type Config struct {
	AK string
	SK string

	TenantId string
	Endpoint string
}

func NewClient(ctx context.Context, cfg *Config) (*Client, error) {
	if cfg.AK == "" || cfg.SK == "" || cfg.Endpoint == "" || cfg.TenantId == "" {
		return nil, errors.New("cfg ak,sk,tenantId,endpoint can not be empty")
	}
	auth, err := qhttp.NewAuthenticator(cfg.AK, cfg.SK)
	if err != nil {
		return nil, err
	}
	c, err := qhttp.NewClient(ctx, qhttp.WithEndpoint(cfg.Endpoint), qhttp.WithAuth(auth))
	if err != nil {
		return nil, err
	}
	return &Client{cc: c, tenantId: cfg.TenantId}, nil
}

type CreateStructuredEvidenceRequest struct {
	// model_id 和 model_name 选填其一，都填写则使用 model id
	ModelId   string `json:"model_id,omitempty"`
	ModelName string `json:"model_name,omitempty"`
	// 必填：模型数据
	NameValues []*KeyValuePair `json:"name_values,omitempty"`
	TenantId   string          `json:"tenant_id,omitempty"`
}

type KeyValuePair struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

func (r *CreateStructuredEvidenceRequest) Payload() string {
	b := strings.Builder{}

	if r.ModelId != "" {
		b.WriteString("model_id=")
		b.WriteString("\"" + r.ModelId + "\"")
	} else if r.ModelName != "" {
		b.WriteString("model_name=")
		b.WriteString("\"" + r.ModelName + "\"")
	}

	b.WriteString("&name_values=")
	json := "["
	for i, v := range r.NameValues {
		json += "{"
		json += fmt.Sprintf(`"key":"%s","value":"%s"`, v.Key, v.Value)
		json += "}"

		if i < len(r.NameValues)-1 {
			json += ","
		}
	}
	json += "]"
	b.WriteString(json)

	b.WriteString("&tenant_id=")
	b.WriteString("\"" + r.TenantId + "\"")

	return b.String()
}

type CreateStructuredEvidenceResponse struct {
	// 存证编号
	EvidenceId string `json:"evidence_id,omitempty"`
}

func (c *Client) CreateStructuredEvidence(ctx context.Context, in *CreateStructuredEvidenceRequest, opts ...qhttp.CallOption) (*CreateStructuredEvidenceResponse, error) {
	var out CreateStructuredEvidenceResponse
	pattern := "/v1/app/structured_evidences"
	path := "/v1/app/structured_evidences"

	opts = append(opts, qhttp.Operation("qilin.api.evidence.CreateStructuredEvidence"))
	opts = append(opts, qhttp.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
