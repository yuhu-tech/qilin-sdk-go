package storage

import (
	"bytes"
	"context"
	"errors"
	"io"
	"mime/multipart"
	"sort"
	"strings"

	qhttp "github.com/yuhu-tech/qilin-sdk-go/qilin/transport/http"
)

const ServiceName = "storage"
const ServiceAPIVersion = "2022-06-02"

type Client struct {
	cc *qhttp.Client
}

type Config struct {
	AK string
	SK string

	Endpoint string
}

func NewClient(ctx context.Context, cfg *Config) (*Client, error) {
	if cfg.AK == "" || cfg.SK == "" || cfg.Endpoint == "" {
		return nil, errors.New("cfg ak,sk,endpoint can not be empty")
	}
	auth, err := qhttp.NewAuthenticator(cfg.AK, cfg.SK)
	if err != nil {
		return nil, err
	}
	c, err := qhttp.NewClient(ctx, qhttp.WithEndpoint(cfg.Endpoint), qhttp.WithAuth(auth))
	if err != nil {
		return nil, err
	}
	return &Client{cc: c}, nil
}

// func NewClientFormConfig(qilin.Config)

func (c *Client) UploadFiles(ctx context.Context, in *UploadFilesInput, opts ...qhttp.CallOption) (*UploadFilesOutPut, error) {
	var out UploadFilesOutPut
	pattern := "/v1/app/storage/{tenant_id}/files"
	path := "/v1/app/storage/" + in.tenant_id + "/files"
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
	b.WriteString("&tenant_id=")
	b.WriteString(u.tenant_id)

	// options
	if u.folder_id != "" {
		b.WriteString("&folder_id=")
		b.WriteString(u.folder_id)
	}
	if u.folder_digest != "" {
		b.WriteString("&folder_digest=")
		b.WriteString(u.folder_digest)
	}

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

type UploadFilesOutPut struct {
}
