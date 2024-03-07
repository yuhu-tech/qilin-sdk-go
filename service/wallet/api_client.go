/**
 * Created by zhouwenzhe on 2024/1/15
 */

package wallet

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/yuhu-tech/qilin-sdk-go/qilin/gerr"
	qhttp "github.com/yuhu-tech/qilin-sdk-go/qilin/transport/http"
)

const ServiceName = "wallet"
const ServiceAPIVersion = "2024-01-31"

var ErrDuplicateWallet = errors.New("wallet has been created")
var ErrNotFoundWallet = errors.New("wallet not found")

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

type CreateWalletRequest struct {
	// 必填：外部用户编号
	SignUserId string `protobuf:"bytes,1,opt,name=sign_user_id,json=signUserId,proto3" json:"sign_user_id,omitempty"`
	// 必填：租户id
	TenantId string `protobuf:"bytes,2,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
	// 必填：链节点实例编号
	ChainInstanceId string `protobuf:"bytes,3,opt,name=chain_instance_id,json=chainInstanceId,proto3" json:"chain_instance_id,omitempty"`
}
type CreateWalletResponse struct {
	// 用户编号
	SignUserId string `protobuf:"bytes,1,opt,name=sign_user_id,json=signUserId,proto3" json:"sign_user_id,omitempty"`
	// 钱包ID
	WalletId string `protobuf:"bytes,2,opt,name=wallet_id,json=walletId,proto3" json:"wallet_id,omitempty"`
	// 账户地址
	Address string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
}
type GetWalletRequest struct {
	// 外部唯一标识
	OutUserId string `protobuf:"bytes,1,opt,name=out_user_id,json=outUserId,proto3" json:"out_user_id,omitempty"`
	// 链节点实例编号
	ChainInstanceId string `protobuf:"bytes,2,opt,name=chain_instance_id,json=chainInstanceId,proto3" json:"chain_instance_id,omitempty"`
	// 租户编号
	TenantId string `protobuf:"bytes,3,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
}
type GetWalletResponse struct {
	// 钱包ID
	WalletId string `protobuf:"bytes,1,opt,name=wallet_id,json=walletId,proto3" json:"wallet_id,omitempty"`
	// 钱包地址
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	// 租户ID
	TenantId string `protobuf:"bytes,3,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
	// 算法
	Algorithm string `protobuf:"bytes,4,opt,name=algorithm,proto3" json:"algorithm,omitempty"`
	// 私钥（明文）
	PrivateKey string `protobuf:"bytes,5,opt,name=private_key,json=privateKey,proto3" json:"private_key,omitempty"`
}

func (r *CreateWalletRequest) Payload() string {
	b := strings.Builder{}
	s1 := fmt.Sprintf("sign_user_id=\"%s\"", r.SignUserId)
	s2 := fmt.Sprintf("tenant_id=\"%s\"", r.TenantId)
	s3 := fmt.Sprintf("chain_instance_id=\"%s\"", r.ChainInstanceId)
	s := strings.Join([]string{s3, s1, s2}, "&")
	b.WriteString(s)
	return b.String()
}
func (r *GetWalletRequest) Payload() string {
	b := strings.Builder{}
	s1 := fmt.Sprintf("out_user_id=\"%s\"", r.OutUserId)
	s2 := fmt.Sprintf("tenant_id=\"%s\"", r.TenantId)
	s3 := fmt.Sprintf("chain_instance_id=\"%s\"", r.ChainInstanceId)
	s := strings.Join([]string{s3, s1, s2}, "&")
	b.WriteString(s)
	return b.String()
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

func (c *Client) CreateWallet(ctx context.Context, in *CreateWalletRequest, opts ...qhttp.CallOption) (*CreateWalletResponse, error) {
	out := new(CreateWalletResponse)
	pattern := "/v1/app/wallets/create"
	path := "/v1/app/wallets/create"

	opts = append(opts, qhttp.Operation("qilin.api.wallet.CreateWallet"))
	opts = append(opts, qhttp.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		if gerr.Reason(err) == "WalletServiceError_DuplicateWallet" {
			return nil, ErrDuplicateWallet
		}
		return nil, err
	}
	return out, nil
}

func (c *Client) GetWallet(ctx context.Context, in *GetWalletRequest, opts ...qhttp.CallOption) (*GetWalletResponse, error) {
	out := new(GetWalletResponse)
	pattern := "/v1/app/wallets/{out_user_id}/{tenant_id}/{chain_instance_id}"
	path := fmt.Sprintf("/v1/app/wallets/%s/%s/%s", in.OutUserId, in.TenantId, in.ChainInstanceId)

	opts = append(opts, qhttp.Operation("qilin.api.wallet.GetWallet"))
	opts = append(opts, qhttp.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, in, &out, opts...)
	if err != nil {
		if gerr.Reason(err) == "Qilin_Basic_Wallet_V1_NotFound" {
			return nil, ErrNotFoundWallet
		}
		return nil, err
	}
	return out, nil
}
