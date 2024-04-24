/**
 * Created by zhouwenzhe on 2024/1/15
 */

package assets

// import (
// 	"context"
// 	"testing"
// )

// const (
// 	TestTenant = "tid-yuhu1"
// 	Ak         = "test-ak"
// 	Sk         = "test-sk"
// 	Endpoint   = "119.3.106.151:10100"
// 	// Endpoint = "127.0.0.1:10000"
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
// 		RequestId: "20240308",
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
// 		TxHash:   "a44ddad075674590a33b65f04e8c78901c50334bf97e4aa1a7dd4546d0c78ced",
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
// 		ContractAddress: "09f90e04378166f7b69bd63d7ee772b675e1bc30",
// 		ReceiverAddress: "0868a3f91d94683060eb29c41970c320e9254cda",
// 		Signer: &Signer{
// 			// WalletId:     "",
// 			SignedUserId: "clkbzvx3ahw1d0767id4kjhrb",
// 		},
// 		TenantId:  "tid-yuhu1",
// 		Amount:    "100",
// 		RequestId: "20240328",
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
// 		TokenId:         "1",
// 		Signer: &Signer{
// 			WalletId:     "wid-YKxpROYVDJKo",
// 			SignedUserId: "cl11xbqem6md90764cnm3vs2u",
// 		},
// 		TenantId:  "tid-yuhu1",
// 		RequestId: "20240328",
// 	})
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	t.Log(transferNFTResponse)
// }

// func TestClient_GetBatchTransferNFTResult(t *testing.T) {
// 	getBatchTransferNFTResult, err := cli.GetBatchTransferNFTResult(ctx, &GetBatchTransferNFTResultRequest{
// 		TxHash:   "792ef4ece5234526849d1e9249d3cb1538aaae0e95c44dc6b6651ad86bf0111e",
// 		TenantId: "tid-yuhu1",
// 	})
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	t.Log(getBatchTransferNFTResult)
// }

// func TestClient_BatchTransferNFT(t *testing.T) {
// 	batchTransferNFTResponse, err := cli.BatchTransferNFT(ctx, &BatchTransferNFTRequest{
// 		ReceiverAddress: "538c0edebebf19b4b30680f8d88b8f5fc4bf4993",
// 		ContractAddress: "4c147d903517bcb76f21aeaf255eb38e20c96018",
// 		Amount:          2,
// 		Signer: &Signer{
// 			SignedUserId: "yuhu1",
// 		},
// 		TenantId:  "tid-yuhu1",
// 		RequestId: "202404031",
// 	})
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	t.Log(batchTransferNFTResponse)
// }

// func TestClient_ListWalletNFTHolding(t *testing.T) {
// 	res, err := cli.ListWalletNFTHolding(ctx, &ListWalletNFTHoldingRequest{
// 		ContractAddressList: []string{"74a55fb59f51faba6fdc8ac94e1706680cb7b622", "159014b2d449396ba7d1178678cea3076f7dec2c"},
// 		WalletAddress:       "0x7dbd5d3efb0c583257167b1efd58af562053b16c",
// 		Limit:               100,
// 		Cursor:              "",
// 		Offset:              0,
// 		IsReversed:          false,
// 		TenantId:            "tid-yuhu1",
// 	})
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	t.Logf("%+v", res)
// }

// func TestClient_ListWalletTokenHolding(t *testing.T) {
// 	res, err := cli.ListWalletTokenHolding(ctx, &ListWalletTokenHoldingRequest{
// 		ContractAddressList: []string{"74a55fb59f51faba6fdc8ac94e1706680cb7b622", "159014b2d449396ba7d1178678cea3076f7dec2c"},
// 		WalletAddress:       "0x7dbd5d3efb0c583257167b1efd58af562053b16c",
// 		TenantId:            "tid-yuhu1",
// 	})
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	for _, ele := range res.WalletTokenHoldingList {
// 		t.Log(ele)
// 	}
// }
