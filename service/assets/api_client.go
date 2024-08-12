/**
 * Created by zhouwenzhe on 2024/1/15
 */

package assets

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/yuhu-tech/qilin-sdk-go/pkg/util/stringutil"
	"github.com/yuhu-tech/qilin-sdk-go/qilin/transport/http"
	qhttp "github.com/yuhu-tech/qilin-sdk-go/qilin/transport/http"
)

const ServiceName = "assets"
const ServiceAPIVersion = "2024-01-15"

var (
	_ http.PayloadMaker = (*CreateArtworkRequest)(nil)
	_ http.PayloadMaker = (*GetArtworkResultRequest)(nil)
	_ http.PayloadMaker = (*GetMintNFTResultRequest)(nil)
	_ http.PayloadMaker = (*GetTransferNFTResultRequest)(nil)
	_ http.PayloadMaker = (*MintNFTRequest)(nil)
	_ http.PayloadMaker = (*TransferNFTRequest)(nil)
	_ http.PayloadMaker = (*GetBatchTransferNFTResultRequest)(nil)
	_ http.PayloadMaker = (*BatchTransferNFTRequest)(nil)
	_ http.PayloadMaker = (*ListWalletNFTHoldingRequest)(nil)
	_ http.PayloadMaker = (*ListWalletTokenHoldingRequest)(nil)
	_ http.PayloadMaker = (*CreateDigitalIPRequest)(nil)
	_ http.PayloadMaker = (*SetDigitalIPBaseURIRequest)(nil)
	_ http.PayloadMaker = (*GetDigitalIPInfoRequest)(nil)
	_ http.PayloadMaker = (*GetDigitalIPNFTInfoRequest)(nil)
	_ http.PayloadMaker = (*GetArtworkInfoRequest)(nil)
	_ http.PayloadMaker = (*GetArtworkNFTInfoRequest)(nil)
)
var _ AssetsServiceClient = (*Client)(nil)

type AssetsServiceClient interface {
	CreateArtwork(ctx context.Context, in *CreateArtworkRequest, opts ...qhttp.CallOption) (*CreateArtworkResponse, error)
	MintNFT(ctx context.Context, in *MintNFTRequest, opts ...qhttp.CallOption) (*MintNFTResponse, error)
	TransferNFT(ctx context.Context, in *TransferNFTRequest, opts ...qhttp.CallOption) (*TransferNFTResponse, error)
	GetArtworkResult(ctx context.Context, in *GetArtworkResultRequest, opts ...qhttp.CallOption) (*GetArtworkResultResponse, error)
	GetMintNFTResult(ctx context.Context, in *GetMintNFTResultRequest, opts ...qhttp.CallOption) (*GetMintNFTResultResponse, error)
	GetTransferNFTResult(ctx context.Context, in *GetTransferNFTResultRequest, opts ...qhttp.CallOption) (*GetTransferNFTResultResponse, error)
	GetBatchTransferNFTResult(ctx context.Context, in *GetBatchTransferNFTResultRequest, opts ...qhttp.CallOption) (*GetBatchTransferNFTResultResponse, error)
	BatchTransferNFT(ctx context.Context, in *BatchTransferNFTRequest, opts ...qhttp.CallOption) (*BatchTransferNFTResponse, error)
	ListWalletTokenHolding(ctx context.Context, in *ListWalletTokenHoldingRequest, opts ...qhttp.CallOption) (*ListWalletTokenHoldingResponse, error)
	ListWalletNFTHolding(ctx context.Context, in *ListWalletNFTHoldingRequest, opts ...qhttp.CallOption) (*ListWalletNFTHoldingResponse, error)
	CreateDigitalIP(ctx context.Context, in *CreateDigitalIPRequest, opts ...qhttp.CallOption) (*CreateDigitalIPResponse, error)
	SetDigitalIPBaseURI(ctx context.Context, in *SetDigitalIPBaseURIRequest, opts ...qhttp.CallOption) (*SetDigitalIPBaseURIResponse, error)
	GetDigitalIPInfo(ctx context.Context, in *GetDigitalIPInfoRequest, opts ...qhttp.CallOption) (*GetDigitalIPInfoResponse, error)
	GetDigitalIPNFTInfo(ctx context.Context, in *GetDigitalIPNFTInfoRequest, opts ...qhttp.CallOption) (*GetDigitalIPNFTInfoResponse, error)
	GetArtworkInfo(ctx context.Context, in *GetArtworkInfoRequest, opts ...qhttp.CallOption) (*GetArtworkInfoResponse, error)
	GetArtworkNFTInfo(ctx context.Context, in *GetArtworkNFTInfoRequest, opts ...qhttp.CallOption) (*GetArtworkNFTInfoResponse, error)
}
type GetArtworkNFTInfoRequest struct {

	// 合约地址
	ContractAddress string `protobuf:"bytes,1,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty"`
	// token id
	TokenId string `protobuf:"bytes,2,opt,name=token_id,json=tokenId,proto3" json:"token_id,omitempty"`
	// 租户id
	TenantId string `protobuf:"bytes,3,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
}

// Payload implements http.PayloadMaker.
func (r *GetArtworkNFTInfoRequest) Payload() string {
	arr := []string{
		fmt.Sprintf("contract_address=\"%s\"", r.ContractAddress),
		fmt.Sprintf("tenant_id=\"%s\"", r.TenantId),
		fmt.Sprintf("token_id=\"%s\"", r.TokenId),
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i][0] < arr[j][0]
	})
	b := strings.Builder{}
	b.WriteString(strings.Join(arr, "&"))
	return b.String()
}

type GetArtworkNFTInfoResponse struct {

	// owner拥有者
	Owner string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
}

type GetArtworkInfoRequest struct {

	// 合约地址
	ContractAddress string `protobuf:"bytes,1,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty"`
	// 租户id
	TenantId string `protobuf:"bytes,2,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
}

// Payload implements http.PayloadMaker.
func (r *GetArtworkInfoRequest) Payload() string {
	arr := []string{
		fmt.Sprintf("contract_address=\"%s\"", r.ContractAddress),
		fmt.Sprintf("tenant_id=\"%s\"", r.TenantId),
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i][0] < arr[j][0]
	})
	b := strings.Builder{}
	b.WriteString(strings.Join(arr, "&"))
	return b.String()
}

type GetArtworkInfoResponse struct {

	// 艺术家姓名
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// 艺术品名称
	Symbol string `protobuf:"bytes,2,opt,name=symbol,proto3" json:"symbol,omitempty"`
	// 艺术品url
	ArtworkUrl string `protobuf:"bytes,3,opt,name=artwork_url,json=artworkUrl,proto3" json:"artwork_url,omitempty"`
	// 摘要
	Digest string `protobuf:"bytes,4,opt,name=digest,proto3" json:"digest,omitempty"`
	// 最大发行量
	MaxSupply string `protobuf:"bytes,5,opt,name=max_supply,json=maxSupply,proto3" json:"max_supply,omitempty"`
	// 合约地址
	ContractAddress string `protobuf:"bytes,6,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty"`
}

type GetDigitalIPNFTInfoRequest struct {

	// 合约地址
	ContractAddress string `protobuf:"bytes,1,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty"`
	// token id
	TokenId string `protobuf:"bytes,2,opt,name=token_id,json=tokenId,proto3" json:"token_id,omitempty"`
	// 租户id
	TenantId string `protobuf:"bytes,3,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
}

// Payload implements http.PayloadMaker.
func (r *GetDigitalIPNFTInfoRequest) Payload() string {
	arr := []string{
		fmt.Sprintf("contract_address=\"%s\"", r.ContractAddress),
		fmt.Sprintf("tenant_id=\"%s\"", r.TenantId),
		fmt.Sprintf("token_id=\"%s\"", r.TokenId),
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i][0] < arr[j][0]
	})
	b := strings.Builder{}
	b.WriteString(strings.Join(arr, "&"))
	return b.String()
}

type GetDigitalIPNFTInfoResponse struct {

	// owner拥有者
	Owner string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	// token uri
	TokenUri string `protobuf:"bytes,2,opt,name=token_uri,json=tokenUri,proto3" json:"token_uri,omitempty"`
}

type GetDigitalIPInfoRequest struct {

	// 合约地址
	ContractAddress string `protobuf:"bytes,1,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty"`
	// 租户id
	TenantId string `protobuf:"bytes,2,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
}

// Payload implements http.PayloadMaker.
func (r *GetDigitalIPInfoRequest) Payload() string {
	arr := []string{
		fmt.Sprintf("contract_address=\"%s\"", r.ContractAddress),
		fmt.Sprintf("tenant_id=\"%s\"", r.TenantId),
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i][0] < arr[j][0]
	})
	b := strings.Builder{}
	b.WriteString(strings.Join(arr, "&"))
	return b.String()
}

type GetDigitalIPInfoResponse struct {

	// 艺术家姓名
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// 艺术品名称
	Symbol string `protobuf:"bytes,2,opt,name=symbol,proto3" json:"symbol,omitempty"`
	// 最大发行量
	MaxSupply string `protobuf:"bytes,5,opt,name=max_supply,json=maxSupply,proto3" json:"max_supply,omitempty"`
	// 合约地址
	ContractAddress string `protobuf:"bytes,6,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty"`
	// base uri
	BaseUri string `protobuf:"bytes,7,opt,name=base_uri,json=baseUri,proto3" json:"base_uri,omitempty"`
}

type SetDigitalIPBaseURIRequest struct {
	// 合约地址
	ContractAddress string `protobuf:"bytes,1,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty"`
	// base uri
	BaseUri string `protobuf:"bytes,2,opt,name=base_uri,json=baseUri,proto3" json:"base_uri,omitempty"`
	// 租户id
	TenantId string `protobuf:"bytes,3,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
	// 请求id
	RequestId string `protobuf:"bytes,4,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	// 签名者
	Signer *Signer `protobuf:"bytes,5,opt,name=signer,proto3" json:"signer,omitempty"`
}

// Payload implements http.PayloadMaker.
func (r *SetDigitalIPBaseURIRequest) Payload() string {
	arr := []string{
		signerString(r.Signer),
		fmt.Sprintf("tenant_id=\"%s\"", r.TenantId),
		fmt.Sprintf("request_id=\"%s\"", r.RequestId),
		fmt.Sprintf("contract_address=\"%s\"", r.ContractAddress),
		fmt.Sprintf("base_uri=\"%s\"", r.BaseUri),
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i][0] < arr[j][0]
	})
	b := strings.Builder{}
	b.WriteString(strings.Join(arr, "&"))
	return b.String()
}

type SetDigitalIPBaseURIResponse struct {
	// 交易哈希
	TxHash string `protobuf:"bytes,2,opt,name=txHash,proto3" json:"txHash,omitempty"`
	// 交易状态
	Status string `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
}
type CreateDigitalIPRequest struct {
	// 作者名称
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// 藏品名称
	Symbol string `protobuf:"bytes,2,opt,name=symbol,proto3" json:"symbol,omitempty"`
	// 藏品最大铸造量
	MaxSupply string `protobuf:"bytes,3,opt,name=max_supply,json=maxSupply,proto3" json:"max_supply,omitempty"`
	// 签名者
	Signer *Signer `protobuf:"bytes,4,opt,name=signer,proto3" json:"signer,omitempty"`
	// 租户id
	TenantId string `protobuf:"bytes,5,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
	// 请求id
	RequestId string `protobuf:"bytes,6,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
}

func (r *CreateDigitalIPRequest) Payload() string {
	arr := []string{
		fmt.Sprintf("max_supply=\"%s\"", r.MaxSupply),
		fmt.Sprintf("name=\"%s\"", r.Name),
		fmt.Sprintf("request_id=\"%s\"", r.RequestId),
		signerString(r.Signer),
		fmt.Sprintf("symbol=\"%s\"", r.Symbol),
		fmt.Sprintf("tenant_id=\"%s\"", r.TenantId),
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i][0] < arr[j][0]
	})
	b := strings.Builder{}
	b.WriteString(strings.Join(arr, "&"))
	return b.String()
}

type CreateDigitalIPResponse struct {
	// 合约地址
	ContractAddress string `protobuf:"bytes,1,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty"`
	// 交易哈希
	Txhash string `protobuf:"bytes,2,opt,name=txhash,proto3" json:"txhash,omitempty"`
	// 交易状态
	Status string `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
}

type ListWalletTokenHoldingRequest struct {
	// 合约列表
	ContractAddressList []string `protobuf:"bytes,1,rep,name=contract_address_list,json=contractAddressList,proto3" json:"contract_address_list,omitempty"`
	// 钱包地址
	WalletAddress string `protobuf:"bytes,2,opt,name=wallet_address,json=walletAddress,proto3" json:"wallet_address,omitempty"`
	// 租户id
	TenantId string `protobuf:"bytes,3,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
}

// Payload implements http.PayloadMaker.
func (r *ListWalletTokenHoldingRequest) Payload() string {
	b := strings.Builder{}
	s1 := fmt.Sprintf("contract_address_list=[%s]", stringutil.StringJoinWithOvercoat("\"", "\"", ",", r.ContractAddressList...))
	s2 := fmt.Sprintf("wallet_address=\"%s\"", r.WalletAddress)
	s7 := fmt.Sprintf("tenant_id=\"%s\"", r.TenantId)
	s := strings.Join([]string{s1, s7, s2}, "&")
	b.WriteString(s)
	return b.String()
}

type ListWalletTokenHoldingResponse_WalletTokenHolding struct {
	// 合约地址
	ContractAddress string `protobuf:"bytes,1,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty"`
	// 数量
	Num string `protobuf:"bytes,2,opt,name=num,proto3" json:"num,omitempty"`
}

type ListWalletTokenHoldingResponse struct {
	// 合约数量列表
	WalletTokenHoldingList []*ListWalletTokenHoldingResponse_WalletTokenHolding `protobuf:"bytes,1,rep,name=wallet_token_holding_list,json=walletTokenHoldingList,proto3" json:"wallet_token_holding_list,omitempty"`
}

type ListWalletNFTHoldingRequest struct {
	// 合约列表
	ContractAddressList []string `protobuf:"bytes,1,rep,name=contract_address_list,json=contractAddressList,proto3" json:"contract_address_list,omitempty"`
	// 钱包地址
	WalletAddress string `protobuf:"bytes,2,opt,name=wallet_address,json=walletAddress,proto3" json:"wallet_address,omitempty"`
	// 限制
	Limit uint32 `protobuf:"bytes,3,opt,name=limit,proto3" json:"limit,omitempty"`
	// 游标
	Cursor string `protobuf:"bytes,4,opt,name=cursor,proto3" json:"cursor,omitempty"`
	// 偏移
	Offset uint32 `protobuf:"bytes,5,opt,name=offset,proto3" json:"offset,omitempty"`
	// 是否倒序
	IsReversed bool `protobuf:"bytes,6,opt,name=is_reversed,json=isReversed,proto3" json:"is_reversed,omitempty"`
	// 租户id
	TenantId string `protobuf:"bytes,7,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
}

// Payload implements http.PayloadMaker.
func (r *ListWalletNFTHoldingRequest) Payload() string {
	b := strings.Builder{}
	s1 := fmt.Sprintf("contract_address_list=[%s]", stringutil.StringJoinWithOvercoat("\"", "\"", ",", r.ContractAddressList...))
	s2 := fmt.Sprintf("wallet_address=\"%s\"", r.WalletAddress)
	s3 := fmt.Sprintf("limit=%d", r.Limit)
	s4 := fmt.Sprintf("cursor=\"%s\"", r.Cursor)
	s5 := fmt.Sprintf("offset=%d", r.Offset)
	s6 := fmt.Sprintf("is_reversed=%t", r.IsReversed)
	s7 := fmt.Sprintf("tenant_id=\"%s\"", r.TenantId)
	s := strings.Join([]string{s1, s4, s6, s3, s5, s7, s2}, "&")
	b.WriteString(s)
	return b.String()
}

type ListWalletNFTHoldingResponse struct {
	// nft列表
	NftHolding []*ListWalletNFTHoldingResponse_NFT `protobuf:"bytes,1,rep,name=nft_holding,json=nftHolding,proto3" json:"nft_holding,omitempty"`
	// 总数
	Count string `protobuf:"bytes,2,opt,name=count,proto3" json:"count,omitempty"`
	// 是否有下页数
	HasNextPage bool `protobuf:"bytes,3,opt,name=has_next_page,json=hasNextPage,proto3" json:"has_next_page,omitempty"`
	// 末尾游标
	EndCursor string `protobuf:"bytes,4,opt,name=end_cursor,json=endCursor,proto3" json:"end_cursor,omitempty"`
}
type ListWalletNFTHoldingResponse_NFT struct {
	// 艺术家姓名
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// 艺术品名称
	Symbol string `protobuf:"bytes,2,opt,name=symbol,proto3" json:"symbol,omitempty"`
	// 艺术品url
	ArtworkUrl string `protobuf:"bytes,3,opt,name=artwork_url,json=artworkUrl,proto3" json:"artwork_url,omitempty"`
	// 摘要
	Digest string `protobuf:"bytes,4,opt,name=digest,proto3" json:"digest,omitempty"`
	// 最大发行量
	MaxSupply string `protobuf:"bytes,5,opt,name=max_supply,json=maxSupply,proto3" json:"max_supply,omitempty"`
	// 合约地址
	ContractAddress string `protobuf:"bytes,6,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty"`
	// token_id
	TokenId string `protobuf:"bytes,7,opt,name=token_id,json=tokenId,proto3" json:"token_id,omitempty"`
	// id
	BalanceTokenId string `protobuf:"bytes,8,opt,name=balance_token_id,json=balanceTokenId,proto3" json:"balance_token_id,omitempty"`
	// 交易哈希
	TxHash string `protobuf:"bytes,5,opt,name=tx_hash,json=txHash,proto3" json:"tx_hash,omitempty"`
}

type BatchTransferNFTRequest struct {
	// 新所有者
	ReceiverAddress string `protobuf:"bytes,1,opt,name=receiver_address,json=receiverAddress,proto3" json:"receiver_address,omitempty"`
	// 合约地址
	ContractAddress string `protobuf:"bytes,2,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty"`
	// 数量
	Amount uint64 `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount,omitempty"`
	// 签名者
	Signer *Signer `protobuf:"bytes,4,opt,name=signer,proto3" json:"signer,omitempty"`
	// 租户id
	TenantId string `protobuf:"bytes,6,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
	// 请求id
	RequestId string `protobuf:"bytes,7,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
}

// Payload implements http.PayloadMaker.
func (r *BatchTransferNFTRequest) Payload() string {
	b := strings.Builder{}
	s1 := fmt.Sprintf("receiver_address=\"%s\"", r.ReceiverAddress)
	s2 := fmt.Sprintf("contract_address=\"%s\"", r.ContractAddress)
	s3 := fmt.Sprintf("amount=%d", r.Amount)
	s4 := fmt.Sprintf("tenant_id=\"%s\"", r.TenantId)
	s5 := fmt.Sprintf("request_id=\"%s\"", r.RequestId)
	s6 := signerString(r.Signer)
	s := strings.Join([]string{s3, s2, s1, s5, s6, s4}, "&")
	b.WriteString(s)
	return b.String()
}

func signerString(signer *Signer) string {
	lst := make([]string, 0)
	if signer.SignedUserId != "" {
		lst = append(lst, fmt.Sprintf("\"signed_user_id\":\"%s\"", signer.SignedUserId))
	}
	if signer.WalletId != "" {
		lst = append(lst, fmt.Sprintf("\"wallet_id\":\"%s\"", signer.WalletId))
	}
	return fmt.Sprintf("signer={%s}", strings.Join(lst, ","))
}

type BatchTransferNFTResponse struct {
	// 交易哈希
	Txhash string `protobuf:"bytes,1,opt,name=txhash,proto3" json:"txhash,omitempty"`
	// 交易状态
	Status string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
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
type GetBatchTransferNFTResultRequest struct {
	// 交易哈希
	TxHash string `protobuf:"bytes,1,opt,name=tx_hash,json=txHash,proto3" json:"tx_hash,omitempty"`
	// 租户id
	TenantId string `protobuf:"bytes,2,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
}

// Payload implements http.PayloadMaker.
func (r *GetBatchTransferNFTResultRequest) Payload() string {
	b := strings.Builder{}
	s1 := fmt.Sprintf("tx_hash=\"%s\"", r.TxHash)
	s2 := fmt.Sprintf("tenant_id=\"%s\"", r.TenantId)
	s := strings.Join([]string{s2, s1}, "&")
	b.WriteString(s)
	return b.String()
}

type GetBatchTransferNFTResultResponse struct {
	// 状态
	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	// token列表
	TokenIdList []string `protobuf:"bytes,2,rep,name=token_id_list,json=tokenIdList,proto3" json:"token_id_list,omitempty"`
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
	s8 := signerString(r.Signer)
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
	s6 := signerString(r.Signer)
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
	s6 := signerString(r.Signer)
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

type Config struct {
	AK string
	SK string

	TenantId string
	Endpoint string
}
type Client struct {
	cc       *qhttp.Client
	tenantId string
}

// GetArtworkNFTInfo implements AssetsServiceClient.
func (c *Client) GetArtworkNFTInfo(ctx context.Context, in *GetArtworkNFTInfoRequest, opts ...qhttp.CallOption) (*GetArtworkNFTInfoResponse, error) {
	out := new(GetArtworkNFTInfoResponse)
	pattern := "/v1/app/artworks/nft_info"
	path := fmt.Sprintf("/v1/app/artworks/nft_info?contract_address=%s&token_id=%s&tenant_id=%s", in.ContractAddress, in.TokenId, in.TenantId)

	opts = append(opts, qhttp.Operation("qilin.api.assets.GetArtworkNFTInfo"))
	opts = append(opts, qhttp.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetArtworkInfo implements AssetsServiceClient.
func (c *Client) GetArtworkInfo(ctx context.Context, in *GetArtworkInfoRequest, opts ...qhttp.CallOption) (*GetArtworkInfoResponse, error) {
	out := new(GetArtworkInfoResponse)
	pattern := "/v1/app/artworks/info"
	path := fmt.Sprintf("/v1/app/artworks/info?contract_address=%s&tenant_id=%s", in.ContractAddress, in.TenantId)

	opts = append(opts, qhttp.Operation("qilin.api.assets.GetArtworkInfo"))
	opts = append(opts, qhttp.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetDigitalIPNFTInfo implements AssetsServiceClient.
func (c *Client) GetDigitalIPNFTInfo(ctx context.Context, in *GetDigitalIPNFTInfoRequest, opts ...qhttp.CallOption) (*GetDigitalIPNFTInfoResponse, error) {
	out := new(GetDigitalIPNFTInfoResponse)
	pattern := "/v1/app/digital_ips/nft_info"
	path := fmt.Sprintf("/v1/app/digital_ips/nft_info?contract_address=%s&token_id=%s&tenant_id=%s", in.ContractAddress, in.TokenId, in.TenantId)

	opts = append(opts, qhttp.Operation("qilin.api.assets.GetDigitalIPNFTInfo"))
	opts = append(opts, qhttp.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetDigitalIPInfo implements AssetsServiceClient.
func (c *Client) GetDigitalIPInfo(ctx context.Context, in *GetDigitalIPInfoRequest, opts ...qhttp.CallOption) (*GetDigitalIPInfoResponse, error) {
	out := new(GetDigitalIPInfoResponse)
	pattern := "/v1/app/digital_ips/info"
	path := fmt.Sprintf("/v1/app/digital_ips/info?contract_address=%s&tenant_id=%s", in.ContractAddress, in.TenantId)

	opts = append(opts, qhttp.Operation("qilin.api.assets.GetDigitalIPInfo"))
	opts = append(opts, qhttp.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SetDigitalIPBaseURI implements AssetsServiceClient.
func (c *Client) SetDigitalIPBaseURI(ctx context.Context, in *SetDigitalIPBaseURIRequest, opts ...qhttp.CallOption) (*SetDigitalIPBaseURIResponse, error) {
	out := new(SetDigitalIPBaseURIResponse)
	pattern := "/v1/app/digital_ips/base_uri"
	path := "/v1/app/digital_ips/base_uri"

	opts = append(opts, qhttp.Operation("qilin.api.assets.SetDigitalIPBaseURI"))
	opts = append(opts, qhttp.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CreateDigitalIP implements AssetsServiceClient.
func (c *Client) CreateDigitalIP(ctx context.Context, in *CreateDigitalIPRequest, opts ...qhttp.CallOption) (*CreateDigitalIPResponse, error) {
	out := new(CreateDigitalIPResponse)
	pattern := "/v1/app/digital_ips"
	path := "/v1/app/digital_ips"

	opts = append(opts, qhttp.Operation("qilin.api.assets.CreateDigitalIP"))
	opts = append(opts, qhttp.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ListWalletNFTHolding implements AssetsServiceClient.
func (c *Client) ListWalletNFTHolding(ctx context.Context, in *ListWalletNFTHoldingRequest, opts ...qhttp.CallOption) (*ListWalletNFTHoldingResponse, error) {
	out := new(ListWalletNFTHoldingResponse)
	pattern := "/v1/app/nfts"
	path := fmt.Sprintf("/v1/app/nfts?%s&wallet_address=%s&limit=%d&cursor=%s&offset=%d&is_reversed=%t&tenant_id=%s",
		stringutil.StringJoinWithOvercoat("contract_address_list=", "", "&", in.ContractAddressList...),
		in.WalletAddress, in.Limit, in.Cursor, in.Offset, in.IsReversed, in.TenantId)

	opts = append(opts, qhttp.Operation("qilin.api.assets.ListWalletNFTHolding"))
	opts = append(opts, qhttp.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ListWalletTokenHolding implements AssetsServiceClient.
func (c *Client) ListWalletTokenHolding(ctx context.Context, in *ListWalletTokenHoldingRequest, opts ...qhttp.CallOption) (*ListWalletTokenHoldingResponse, error) {
	out := new(ListWalletTokenHoldingResponse)
	pattern := "/v1/app/nfts"
	path := fmt.Sprintf("/v1/app/nfts/num?%s&wallet_address=%s&tenant_id=%s",
		stringutil.StringJoinWithOvercoat("contract_address_list=", "", "&", in.ContractAddressList...),
		in.WalletAddress, in.TenantId)

	opts = append(opts, qhttp.Operation("qilin.api.assets.ListWalletNFTHolding"))
	opts = append(opts, qhttp.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BatchTransferNFT implements AssetsServiceClient.
func (c *Client) BatchTransferNFT(ctx context.Context, in *BatchTransferNFTRequest, opts ...qhttp.CallOption) (*BatchTransferNFTResponse, error) {
	out := new(BatchTransferNFTResponse)
	pattern := "/v1/app/nfts:batch_transfer"
	path := "/v1/app/nfts:batch_transfer"

	opts = append(opts, qhttp.Operation("qilin.api.assets.BatchTransferNFT"))
	opts = append(opts, qhttp.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetBatchTransferNFTResult implements AssetsServiceClient.
func (c *Client) GetBatchTransferNFTResult(ctx context.Context, in *GetBatchTransferNFTResultRequest, opts ...qhttp.CallOption) (*GetBatchTransferNFTResultResponse, error) {
	out := new(GetBatchTransferNFTResultResponse)
	pattern := "/v1/app/nfts:batch_stransfer/result"
	path := fmt.Sprintf("/v1/app/nfts:batch_stransfer/result?tx_hash=%s&tenant_id=%s", in.TxHash, in.TenantId)

	opts = append(opts, qhttp.Operation("qilin.api.assets.GetBatchTransferNFTResult"))
	opts = append(opts, qhttp.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
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
