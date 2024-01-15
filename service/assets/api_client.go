/**
 * Created by zhouwenzhe on 2024/1/15
 */

package assets

import (
	"context"
	"errors"
	"fmt"
	"strings"

	qhttp "github.com/yuhu-tech/qilin-sdk-go/qilin/transport/http"
)

const ServiceName = "assets"
const ServiceAPIVersion = "2024-01-15"

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

type Signer struct {
	// 钱包id
	WalletId string `protobuf:"bytes,1,opt,name=wallet_id,json=walletId,proto3" json:"wallet_id,omitempty"`
	// 签名用户id
	SignedUserId string `protobuf:"bytes,2,opt,name=signed_user_id,json=signedUserId,proto3" json:"signed_user_id,omitempty"`
}

type CreateArtworkRequest struct {
	// 作者名称
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// 藏品名称
	Symbol string `protobuf:"bytes,2,opt,name=symbol,proto3" json:"symbol,omitempty"`
	// 藏品url
	ArtworkUrl string `protobuf:"bytes,3,opt,name=artwork_url,json=artworkUrl,proto3" json:"artwork_url,omitempty"`
	// 藏品摘要
	Digest string `protobuf:"bytes,4,opt,name=digest,proto3" json:"digest,omitempty"`
	// 藏品最大铸造量
	MaxSupply string `protobuf:"bytes,5,opt,name=max_supply,json=maxSupply,proto3" json:"max_supply,omitempty"`
	// 租户id
	TenantId string `protobuf:"bytes,6,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
	// 签名者
	Signer *Signer `protobuf:"bytes,8,opt,name=signer,proto3" json:"signer,omitempty"`
	// 请求id
	RequestId string `protobuf:"bytes,10,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
}

type CreateArtworkResponse struct {
	// 合约地址
	ContractAddress string `protobuf:"bytes,1,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty"`
	// 交易哈希
	Txhash string `protobuf:"bytes,2,opt,name=txhash,proto3" json:"txhash,omitempty"`
	// 交易状态
	Status string `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
}

type GetArtworkResultRequest struct {
	// 交易哈希
	Txhash string `protobuf:"bytes,1,opt,name=txhash,proto3" json:"txhash,omitempty"`
	// 租户id
	TenantId string `protobuf:"bytes,3,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
}

type GetArtworkResultResponse struct {
	// 状态
	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	// 合约地址
	ContractAddress string `protobuf:"bytes,2,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty"`
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

func (c *Client) CreateArtwork(ctx context.Context, in *CreateArtworkRequest, opts ...qhttp.CallOption) (*CreateArtworkResponse, error) {
	out := new(CreateArtworkResponse)
	pattern := "/v1/app/artworks"
	path := "/v1/app/artworks"

	opts = append(opts, qhttp.Operation("qilin.api.assets.CreateArtwork"))
	opts = append(opts, qhttp.PathTemplate(pattern))

	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *Client) GetArtworkResult(ctx context.Context, in *GetArtworkResultRequest, opts ...qhttp.CallOption) (*GetArtworkResultResponse, error) {
	out := new(GetArtworkResultResponse)
	pattern := "/v1/app/artworks/result"
	path := "/v1/app/artworks/result"

	opts = append(opts, qhttp.Operation("qilin.api.assets.GetArtworkResult"))
	opts = append(opts, qhttp.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// model_name="jingkuang-nft"&name_values=[{"key":"contractaddress","value":"0x13FF557D9578f2Fd78fF47383184A70e685b600f"},
// {"key":"zichanmingcheng","value":"雨水卡"},{"key"huangzuozhe","value":"版易官方平台-004"},
// {"key":"faxingliang","value":"100"},{"key":"faxingshijian","value":"2023-07-17 18:55:15"},
// {"key":"liulanqidizhi","value":"httpslorer.jingkuang.info/address/0x13FF557D9578f2Fd78fF47383184A70e685b600f"}]&tenant_id="tid-yuhu1"
func (r *CreateArtworkRequest) Payload() string {
	b := strings.Builder{}
	s1 := fmt.Sprintf("name=\"%s\"", r.Name)
	s2 := fmt.Sprintf("symbol=\"%s\"", r.Symbol)
	s3 := fmt.Sprintf("artwork_url=\"%s\"", r.ArtworkUrl)
	s4 := fmt.Sprintf("digest=\"%s\"", r.Digest)
	s5 := fmt.Sprintf("max_supply=\"%s\"", r.MaxSupply)
	s6 := fmt.Sprintf("tenant_id=\"%s\"", r.TenantId)
	s7 := fmt.Sprintf("request_id=\"%s\"", r.RequestId)
	s8 := fmt.Sprintf("signer={\"signed_user_id\":\"%s\",\"wallet_id\":\"%s\"}", r.Signer.SignedUserId, r.Signer.WalletId)
	s := strings.Join([]string{s3, s4, s5, s1, s7, s8, s2, s6}, "&")
	b.WriteString(s)
	return b.String()
}

func (r *GetArtworkResultRequest) Payload() string {
	b := strings.Builder{}
	s1 := fmt.Sprintf("txhash=\"%s\"", r.Txhash)
	s2 := fmt.Sprintf("tenant_id=\"%s\"", r.TenantId)
	s := strings.Join([]string{s2, s1}, "&")
	b.WriteString(s)
	return b.String()
}
