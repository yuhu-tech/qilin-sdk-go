/**
 * Created by zhouwenzhe on 2024/1/15
 */

package assets

// import (
// 	"testing"
// )

// import (
// 	"context"
// )

// const (
// 	TestTenant = "tid-yuhu1"
// 	Ak         = "test-ak"
// 	Sk         = "test-sk"
// 	Endpoint   = "119.3.106.151:10100"
// 	//Endpoint = "127.0.0.1:20000"
// )

// var cli *Client
// var ctx context.Context

// func TestMain(m *testing.M) {
// 	client, err := NewClient(context.Background(), &Config{
// 		AK: Ak, SK: Sk, TenantId: TestTenant, Endpoint: Endpoint,
// 	})
// 	if err != nil {
// 		panic(err)
// 	}
// 	cli = client
// 	ctx = context.Background()
// 	m.Run()
// }

// func TestClient_CreateArtwork(t *testing.T) {
// 	createArtworkResponse, err := cli.CreateArtwork(ctx, &CreateArtworkRequest{
// 		Name:       "name7",
// 		Symbol:     "symbol7",
// 		ArtworkUrl: "http://sdnft2/",
// 		Digest:     "03f2c34ce3d4a350fab2adad881aa5fff99cc4c366bab345455b068fc9a21f22",
// 		MaxSupply:  "1000",
// 		TenantId:   "tid-yuhu1",
// 		Signer: &Signer{
// 			// WalletId:     "wid-nZJKYzZ3K55v",
// 			SignedUserId: "clkbzvx3ahw1d0767id4kjhrb",
// 		},
// 		RequestId: "20240221",
// 	})
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	t.Log(createArtworkResponse)
// }

// func TestClient_GetArtworkResult(t *testing.T) {
// 	getArtworkResultResponse, err := cli.GetArtworkResult(ctx, &GetArtworkResultRequest{
// 		Txhash:   "0151b169d1ab4fc18c93d609dfdbfdccc81aaf1d82eb48a0af8b37fc22f28c4f",
// 		TenantId: "tid-yuhu1",
// 	})
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	t.Log(getArtworkResultResponse)
// }

// func TestClient_GetMintNFTResult(t *testing.T) {
// 	getMintNFTResultResponse, err := cli.GetMintNFTResult(ctx, &GetMintNFTResultRequest{
// 		TxHash:   "8518e4a79ce346838ab3dc61829ac5d4197a865fb8654e4e84b19eca852d308c",
// 		TenantId: "tid-yuhu1",
// 	})
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	t.Log(getMintNFTResultResponse)
// }

// func TestClient_GetTransferNFTResult(t *testing.T) {
// 	getTransferNFTResultResponse, err := cli.GetTransferNFTResult(ctx, &GetTransferNFTResultRequest{
// 		TxHash:   "d1edde52dd8840d9905818df22561f6fa9a50a9063d54f94bb74e377dca8f025",
// 		TenantId: "tid-yuhu1",
// 	})
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	t.Log(getTransferNFTResultResponse)
// }

// func TestClient_MintNFT(t *testing.T) {
// 	mintNFTResponse, err := cli.MintNFT(ctx, &MintNFTRequest{
// 		ContractAddress: "90c2a20fe818736f89f17b8db7933ae40a4d5fe8",
// 		ReceiverAddress: "538c0edebebf19b4b30680f8d88b8f5fc4bf4993",
// 		Signer: &Signer{
// 			// WalletId:     "",
// 			SignedUserId: "clkbzvx3ahw1d0767id4kjhrb",
// 		},
// 		TenantId:  "tid-yuhu1",
// 		Amount:    "100",
// 		RequestId: "202402192",
// 	})
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	t.Log(mintNFTResponse)
// }

// func TestClient_TransferNFT(t *testing.T) {
// 	transferNFTResponse, err := cli.TransferNFT(ctx, &TransferNFTRequest{
// 		ReceiverAddress: "538c0edebebf19b4b30680f8d88b8f5fc4bf4993",
// 		ContractAddress: "5972275e95921084d75106554de1ab25a7c0459f",
// 		TokenId:         "100",
// 		Signer: &Signer{
// 			WalletId:     "wid-YKxpROYVDJKo",
// 			SignedUserId: "cl11xbqem6md90764cnm3vs2u",
// 		},
// 		TenantId:  "tid-yuhu1",
// 		RequestId: "202402191",
// 	})
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	t.Log(transferNFTResponse)
// }
