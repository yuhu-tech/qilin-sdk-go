package storage

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"sort"
	"strings"

	qhttp "github.com/yuhu-tech/qilin-sdk-go/qilin/transport/http"
)

const ServiceName = "storage"
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

// func NewClientFormConfig(qilin.Config)

func (c *Client) UploadFiles(ctx context.Context, in *UploadFilesInput, opts ...qhttp.CallOption) (*UploadFilesOutPut, error) {
	var out UploadFilesOutPut
	pattern := "/v1/app/storage/{tenant_id}/files"
	path := "/v1/app/storage/" + c.tenantId + "/files"
	// TODO 通用解析
	// path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, qhttp.Operation("qilin.api.storage.UploadFiles"))
	opts = append(opts, qhttp.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

type Option func(*UploadFilesInput)

func WithFolderId(folder_id string) Option {
	return func(ufi *UploadFilesInput) {
		ufi.folder_id = folder_id
	}
}

func WithFolderDigest(folder_digest string) Option {
	return func(ufi *UploadFilesInput) {
		ufi.folder_id = folder_digest
	}
}

var _ qhttp.FormDataParser = (*UploadFilesInput)(nil)

type UploadFilesInput struct {
	tenant_id     string
	folder_id     string
	folder_digest string
	files         []Files
}
type UploadFilesOutPut struct {
}
type Files struct {
	FileName string
	Data     io.Reader
}

func NewUploadFilesInput(tenant_id string, fs []Files, opts ...Option) (*UploadFilesInput, error) {
	input := &UploadFilesInput{
		tenant_id: tenant_id,
		files:     fs,
	}
	for _, o := range opts {
		o(input)
	}

	if input.tenant_id == "" {
		return nil, errors.New("need set tenant_id")
	}
	if len(input.files) == 0 {
		return nil, errors.New("need with files")
	}

	return input, nil
}

func (u *UploadFilesInput) Payload() string {
	b := strings.Builder{}
	// must have
	var fileNames string
	if len(u.files) == 1 {
		fileNames = u.files[0].FileName
	} else {
		names := make([]string, 0, len(u.files))
		for _, f := range u.files {
			names = append(names, f.FileName)
		}
		sort.Strings(names)
		fileNames = strings.Join(names, "|")
	}
	b.WriteString("files=")
	b.WriteString(fileNames)

	if u.folder_digest != "" {
		b.WriteString("&folder_digest=")
		b.WriteString(u.folder_digest)
	}
	if u.folder_id != "" {
		b.WriteString("&folder_id=")
		b.WriteString(u.folder_id)
	}

	b.WriteString("&tenant_id=")
	b.WriteString(u.tenant_id)

	return b.String()
}

// to parse form data
func (u *UploadFilesInput) ParseFormData() (data io.Reader, contentType string, err error) {
	// TODO set to buffer pool
	bodyBuffer := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuffer)
	defer bodyWriter.Close()

	for _, f := range u.files {
		w, err := bodyWriter.CreateFormFile("files", f.FileName)
		if err != nil {
			return nil, "", err
		}
		if _, err := io.Copy(w, f.Data); err != nil {
			return nil, "", err
		}

	}
	if err := bodyWriter.WriteField("tenant_id", u.tenant_id); err != nil {
		return nil, "", err
	}
	if u.folder_id != "" {
		if err := bodyWriter.WriteField("folder_id", u.folder_id); err != nil {
			return nil, "", err
		}
	}
	if u.folder_digest != "" {
		if err := bodyWriter.WriteField("folder_digest", u.folder_digest); err != nil {
			return nil, "", err
		}
	}

	return bodyBuffer, bodyWriter.FormDataContentType(), nil
}

type CreateFolderRequest struct {
	// 必填 文件夹名称
	FolderName string `json:"folder_name,omitempty"`
	// 下面两个参数用来指定文件夹的父目录，优先级 parent_id > folder_path, 都不传则使用root
	// 选填 父id，不存在报错
	ParentId string `json:"parent_id,omitempty"`
	// 选填 父路径，如果路径不存在则创建
	FolderPath string `json:"folder_path,omitempty"`
}
type CreateFolderResponse struct {
	FolderId        string `json:"folder_id,omitempty"`
	FolderDigest    string `json:"folder_digest,omitempty"`
	FolderUri       string `json:"folder_uri,omitempty"`
	FolderDigestUri string `json:"folder_digest_uri,omitempty"`
}

func (r *CreateFolderRequest) Payload() string {
	b := strings.Builder{}

	b.WriteString("folder_name=")
	b.WriteString("\"" + r.FolderName + "\"")
	// parent set
	if r.FolderPath != "" {
		b.WriteString("&folder_path=")
		b.WriteString("\"" + r.FolderPath + "\"")

	}
	if r.ParentId != "" {
		b.WriteString("&parent_id=")
		b.WriteString("\"" + r.ParentId + "\"")
	}

	fmt.Println(b.String())
	return b.String()
}

func (c *Client) CreateFolder(ctx context.Context, in *CreateFolderRequest, opts ...qhttp.CallOption) (*CreateFolderResponse, error) {
	var out CreateFolderResponse
	pattern := "/v1/app/storage/{tenant_id}/folder"
	path := "/v1/app/storage/" + c.tenantId + "/folder"

	opts = append(opts, qhttp.Operation("qilin.api.storage.CreateFolder"))
	opts = append(opts, qhttp.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

type DownloadFileRequest struct {
	FolderSearchKey string `json:"folder_search_key,omitempty"`
	FileName        string `json:"file_name,omitempty"`
}

// 实现 LoadContent接口, 参考 aws 和 文档
type DownloadFileResponse struct {
	buf bytes.Buffer
}

func (r *DownloadFileRequest) Payload() string {
	b := strings.Builder{}

	b.WriteString("file_name=")
	b.WriteString("\"" + r.FileName + "\"")
	b.WriteString("&folder_search_key=")
	b.WriteString("\"" + r.FolderSearchKey + "\"")

	return b.String()
}

func (c *Client) DownloadFile(ctx context.Context, in *DownloadFileRequest, opts ...qhttp.CallOption) (*UploadFilesOutPut, error) {
	var out UploadFilesOutPut
	pattern := "/v1/app/storage/{tenant_id}/{folder_search_key}/{file_name}"
	path := "/v1/app/storage/" + c.tenantId + "/" + in.FolderSearchKey + "/" + in.FileName

	opts = append(opts, qhttp.Operation("qilin.api.storage.DownloadFile"))
	opts = append(opts, qhttp.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
