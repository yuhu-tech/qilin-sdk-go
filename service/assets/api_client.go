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
	WalletId string `protobuf:"bytes,1,opt,name=wallet_id,json=walletId,proto3" json:"wallet_id"`
	// 签名用户id
	SignedUserId string `protobuf:"bytes,2,opt,name=signed_user_id,json=signedUserId,proto3" json:"signed_user_id"`
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
type GetMintNFTResultRequest struct {
	// 交易哈希
	TxHash string `protobuf:"bytes,2,opt,name=tx_hash,json=txHash,proto3" json:"tx_hash,omitempty"`
	// 租户id
	TenantId string `protobuf:"bytes,3,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
}
type GetMintNFTResultResponse struct {
	// token id
	TokenIds []string `protobuf:"bytes,1,rep,name=token_ids,json=tokenIds,proto3" json:"token_ids,omitempty"`
	// 钱包地址
	WalletAddress string `protobuf:"bytes,2,opt,name=wallet_address,json=walletAddress,proto3" json:"wallet_address,omitempty"`
	// 状态
	Status string `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
}
type GetTransferNFTResultRequest struct {
	// 交易哈希
	TxHash string `protobuf:"bytes,1,opt,name=tx_hash,json=txHash,proto3" json:"tx_hash,omitempty"`
	// 租户id
	TenantId string `protobuf:"bytes,2,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
}
type GetTransferNFTResultResponse struct {
	// 状态
	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}
type MintNFTRequest struct {
	// 合约地址
	ContractAddress string `protobuf:"bytes,1,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty"`
	// 所有者
	ReceiverAddress string `protobuf:"bytes,2,opt,name=receiver_address,json=receiverAddress,proto3" json:"receiver_address,omitempty"`
	// 签名者
	Signer *Signer `protobuf:"bytes,3,opt,name=signer,proto3" json:"signer,omitempty"`
	// 租户id
	TenantId string `protobuf:"bytes,4,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
	// 数量
	Amount string `protobuf:"bytes,5,opt,name=amount,proto3" json:"amount,omitempty"`
	// 请求id
	RequestId string `protobuf:"bytes,6,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
}
type MintNFTResponse struct {
	// token id
	TokenIds []string `protobuf:"bytes,1,rep,name=token_ids,json=tokenIds,proto3" json:"token_ids,omitempty"`
	// 交易哈希
	TxHash string `protobuf:"bytes,2,opt,name=tx_hash,json=txHash,proto3" json:"txhash,omitempty"`
	// 交易状态
	Status string `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	// 钱包地址
	ContractAddress string `protobuf:"bytes,4,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty"`
}
type TransferNFTRequest struct {
	// 新所有者
	ReceiverAddress string `protobuf:"bytes,2,opt,name=receiver_address,json=receiverAddress,proto3" json:"receiver_address,omitempty"`
	// 合约地址
	ContractAddress string `protobuf:"bytes,3,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty"`
	// token id
	TokenId string `protobuf:"bytes,4,opt,name=token_id,json=tokenId,proto3" json:"token_id,omitempty"`
	// 签名者
	Signer *Signer `protobuf:"bytes,5,opt,name=signer,proto3" json:"signer,omitempty"`
	// 租户id
	TenantId string `protobuf:"bytes,6,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
	// 请求id
	RequestId string `protobuf:"bytes,7,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
}
type TransferNFTResponse struct {
	// 交易哈希
	TxHash string `protobuf:"bytes,1,opt,name=tx_hash,json=txHash,proto3" json:"txhash,omitempty"`
	// 交易状态
	Status string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
}

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

func (r *GetMintNFTResultRequest) Payload() string {
	b := strings.Builder{}
	s1 := fmt.Sprintf("tx_hash=\"%s\"", r.TxHash)
	s2 := fmt.Sprintf("tenant_id=\"%s\"", r.TenantId)
	s := strings.Join([]string{s2, s1}, "&")
	b.WriteString(s)
	return b.String()
}

func (r *GetTransferNFTResultRequest) Payload() string {
	b := strings.Builder{}
	s1 := fmt.Sprintf("tx_hash=\"%s\"", r.TxHash)
	s2 := fmt.Sprintf("tenant_id=\"%s\"", r.TenantId)
	s := strings.Join([]string{s2, s1}, "&")
	b.WriteString(s)
	return b.String()
}

func (r *MintNFTRequest) Payload() string {
	b := strings.Builder{}
	s1 := fmt.Sprintf("receiver_address=\"%s\"", r.ReceiverAddress)
	s2 := fmt.Sprintf("contract_address=\"%s\"", r.ContractAddress)
	s3 := fmt.Sprintf("amount=\"%s\"", r.Amount)
	s4 := fmt.Sprintf("tenant_id=\"%s\"", r.TenantId)
	s5 := fmt.Sprintf("request_id=\"%s\"", r.RequestId)
	s6 := fmt.Sprintf("signer={\"signed_user_id\":\"%s\",\"wallet_id\":\"%s\"}", r.Signer.SignedUserId, r.Signer.WalletId)
	s := strings.Join([]string{s3, s2, s1, s5, s6, s4}, "&")
	b.WriteString(s)
	return b.String()
}
func (r *TransferNFTRequest) Payload() string {
	b := strings.Builder{}
	s1 := fmt.Sprintf("receiver_address=\"%s\"", r.ReceiverAddress)
	s2 := fmt.Sprintf("contract_address=\"%s\"", r.ContractAddress)
	s3 := fmt.Sprintf("token_id=\"%s\"", r.TokenId)
	s4 := fmt.Sprintf("tenant_id=\"%s\"", r.TenantId)
	s5 := fmt.Sprintf("request_id=\"%s\"", r.RequestId)
	s6 := fmt.Sprintf("signer={\"signed_user_id\":\"%s\",\"wallet_id\":\"%s\"}", r.Signer.SignedUserId, r.Signer.WalletId)
	s := strings.Join([]string{s2, s1, s5, s6, s4, s3}, "&")
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
	path := fmt.Sprintf("/v1/app/artworks/result?tenant_id=%s&txhash=%s", in.TenantId, in.Txhash)

	opts = append(opts, qhttp.Operation("qilin.api.assets.GetArtworkResult"))
	opts = append(opts, qhttp.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *Client) GetMintNFTResult(ctx context.Context, in *GetMintNFTResultRequest, opts ...qhttp.CallOption) (*GetMintNFTResultResponse, error) {
	out := new(GetMintNFTResultResponse)
	pattern := "/v1/app/nfts:mint/result"
	path := fmt.Sprintf("/v1/app/nfts:mint/result?tx_hash=%s&tenant_id=%s", in.TxHash, in.TenantId)

	opts = append(opts, qhttp.Operation("qilin.api.assets.GetMintNFTResult"))
	opts = append(opts, qhttp.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *Client) GetTransferNFTResult(ctx context.Context, in *GetTransferNFTResultRequest, opts ...qhttp.CallOption) (*GetTransferNFTResultResponse, error) {
	out := new(GetTransferNFTResultResponse)
	pattern := "/v1/app/nfts:transfer/result"
	path := fmt.Sprintf("/v1/app/nfts:transfer/result?tx_hash=%s&tenant_id=%s", in.TxHash, in.TenantId)

	opts = append(opts, qhttp.Operation("qilin.api.assets.GetMintNFTResult"))
	opts = append(opts, qhttp.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *Client) MintNFT(ctx context.Context, in *MintNFTRequest, opts ...qhttp.CallOption) (*MintNFTResponse, error) {
	out := new(MintNFTResponse)
	pattern := "/v1/app/nfts:mint"
	path := "/v1/app/nfts:mint"

	opts = append(opts, qhttp.Operation("qilin.api.assets.MintNFT"))
	opts = append(opts, qhttp.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *Client) TransferNFT(ctx context.Context, in *TransferNFTRequest, opts ...qhttp.CallOption) (*TransferNFTResponse, error) {
	out := new(TransferNFTResponse)
	pattern := "/v1/app/nfts:transfer"
	path := "/v1/app/nfts:transfer"

	opts = append(opts, qhttp.Operation("qilin.api.assets.TransferNFT"))
	opts = append(opts, qhttp.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
