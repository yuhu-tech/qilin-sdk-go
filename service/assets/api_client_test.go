/**
 * Created by zhouwenzhe on 2024/1/15
 */

package assets

// import (
// 	"context"
// 	"testing"

// 	_ "github.com/go-sql-driver/mysql"
// )

// const (
// 	TestTenant = "tid-yuhu1"
// 	Ak         = "test-ak"
// 	Sk         = "test-sk"
// 	Endpoint   = "122.112.237.200:10100"
// 	// Endpoint = "127.0.0.1:10000"
// )

// var cli *Client
// var ctx context.Context
// var signer = &Signer{
// 	WalletId:     "wid-065gmNjp8G5v",
// 	SignedUserId: "zwztest",
// }

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
// 		Name:       "name",
// 		Symbol:     "symbol",
// 		ArtworkUrl: "http://sdnft2/",
// 		Digest:     "03f2c34ce3d4a350fab2adad881aa5fff99cc4c366bab345455b068fc9a21f22",
// 		MaxSupply:  "10000",
// 		TenantId:   "tid-yuhu1",
// 		Signer: &Signer{
// 			SignedUserId: "zwz2024061801",
// 		},
// 		RequestId: "2024072301",
// 	})
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	t.Logf("%+v", createArtworkResponse)
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
// 		ContractAddress: "4eb05e7cd013b7386aa31c7eb83b6a08e29f5bba",
// 		ReceiverAddress: "0f97914bd90feca31d91acc8d3d7ef58d27d4033",
// 		Signer:          &Signer{WalletId: "wid-rLJynMvP6K03", SignedUserId: "test-user1"},
// 		TenantId:        "tid-yuhu1",
// 		Amount:          "2",
// 		RequestId:       "202408071",
// 	})
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	t.Logf("%+v", mintNFTResponse)
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
// 		ReceiverAddress: "95a36924dab56aee99e8622c4b8f2e9511b8c8c1",
// 		ContractAddress: "0b15a33813365464bc0f57165d7af692ee1046a5",
// 		Amount:          100,
// 		Signer: &Signer{
// 			WalletId:     "wid-rLJynMvP6K03",
// 			SignedUserId: "test-user1",
// 		},
// 		TenantId:  "tid-yuhu1",
// 		RequestId: "1001011111",
// 	})
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	t.Log(batchTransferNFTResponse)
// }

// func TestClient_ListWalletNFTHolding(t *testing.T) {
// 	res, err := cli.ListWalletNFTHolding(ctx, &ListWalletNFTHoldingRequest{
// 		// ContractAddressList: []string{"09f90e04378166f7b69bd63d7ee772b675e1bc30"},
// 		WalletAddress: "538c0edebebf19b4b30680f8d88b8f5fc4bf4993",
// 		Limit:         10,
// 		Cursor:        "",
// 		Offset:        0,
// 		IsReversed:    false,
// 		TenantId:      "tid-yuhu1",
// 	})
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	t.Logf("%+v", res)
// }

// func TestClient_ListWalletTokenHolding(t *testing.T) {
// 	res, err := cli.ListWalletTokenHolding(ctx, &ListWalletTokenHoldingRequest{
// 		ContractAddressList: []string{"09f90e04378166f7b69bd63d7ee772b675e1bc30"},
// 		WalletAddress:       "cb393a59cfcae34f6cfd9c629fe122d93a3041d1",
// 		TenantId:            "tid-yuhu1",
// 	})
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	t.Log(res)
// 	for _, ele := range res.WalletTokenHoldingList {
// 		t.Log(ele)
// 	}
// }

// func TestClient_CreateDigitalIP(t *testing.T) {
// 	res, err := cli.CreateDigitalIP(ctx, &CreateDigitalIPRequest{
// 		Name:      "name1",
// 		Symbol:    "symbol1",
// 		MaxSupply: "100",
// 		Signer:    &Signer{WalletId: "wid-rLJynMvP6K03", SignedUserId: "test-user1"},
// 		TenantId:  TestTenant,
// 		RequestId: "201408061",
// 	})
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	t.Logf("%+v", res)
// }

// func TestClient_SetDigitalIPBaseURI(t *testing.T) {
// 	res, err := cli.SetDigitalIPBaseURI(ctx, &SetDigitalIPBaseURIRequest{
// 		ContractAddress: "4eb05e7cd013b7386aa31c7eb83b6a08e29f5bba",
// 		BaseUri:         "https://console.yuhu.tech/api/v1/app/storage/tid-yuhu1/278ae27255de65c7ecfdaa67006284d0cbfaa6ca282046fef0bf99fa6ab9a504/",
// 		Signer:          &Signer{WalletId: "wid-rLJynMvP6K03", SignedUserId: "test-user1"},
// 		TenantId:        TestTenant,
// 		RequestId:       "2014080611",
// 	})
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	t.Log(res)
// }

// func TestClient_GetDigitalIPInfo(t *testing.T) {
// 	res, err := cli.GetDigitalIPInfo(ctx, &GetDigitalIPInfoRequest{
// 		ContractAddress: "4eb05e7cd013b7386aa31c7eb83b6a08e29f5bba",
// 		TenantId:        TestTenant,
// 	})
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	t.Logf("%+v", res)
// }

// func TestClient_GetDigitalIPNFTInfo(t *testing.T) {
// 	res, err := cli.GetDigitalIPNFTInfo(ctx, &GetDigitalIPNFTInfoRequest{
// 		ContractAddress: "4eb05e7cd013b7386aa31c7eb83b6a08e29f5bba",
// 		TokenId:         "0",
// 		TenantId:        TestTenant,
// 	})
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	t.Logf("%+v", res)
// }

// func TestClient_GetArtworkInfo(t *testing.T) {
// 	res, err := cli.GetArtworkInfo(ctx, &GetArtworkInfoRequest{
// 		ContractAddress: "1d0fb72b31cb4932bde70a99ac73a2a763ab38aa",
// 		TenantId:        TestTenant,
// 	})
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	t.Logf("%+v", res)
// }

// func TestClient_GetArtworkNFTInfo(t *testing.T) {
// 	res, err := cli.GetArtworkNFTInfo(ctx, &GetArtworkNFTInfoRequest{
// 		ContractAddress: "74a55fb59f51faba6fdc8ac94e1706680cb7b622",
// 		TokenId:         "1",
// 		TenantId:        TestTenant,
// 	})
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	t.Logf("%+v", res)
// }
