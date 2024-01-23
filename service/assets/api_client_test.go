/**
 * Created by zhouwenzhe on 2024/1/15
 */

package assets

//
//import (
//	"testing"
//)
//
//import (
//	"context"
//)
//
//const (
//	TestTenant = "tid-yuhu1"
//	Ak         = "test-ak"
//	Sk         = "test-sk"
//	Endpoint   = "119.3.106.151:10100"
//	//Endpoint = "127.0.0.1:20000"
//)
//
//var cli *Client
//var ctx context.Context
//
//func TestMain(m *testing.M) {
//	client, err := NewClient(context.Background(), &Config{
//		AK: Ak, SK: Sk, TenantId: TestTenant, Endpoint: Endpoint,
//	})
//	if err != nil {
//		panic(err)
//	}
//	cli = client
//	ctx = context.Background()
//	m.Run()
//}
//
//func TestClient_CreateArtwork(t *testing.T) {
//	createArtworkResponse, err := cli.CreateArtwork(ctx, &CreateArtworkRequest{
//		Name:       "name7",
//		Symbol:     "symbol7",
//		ArtworkUrl: "http://sdnft2/",
//		Digest:     "03f2c34ce3d4a350fab2adad881aa5fff99cc4c366bab345455b068fc9a21f22",
//		MaxSupply:  "1000",
//		TenantId:   "tid-yuhu1",
//		Signer: &Signer{
//			"wid-llw34jmEw8vE",
//			"yuhu1",
//		},
//		RequestId: "request7",
//	})
//	if err != nil {
//		t.Error(err)
//	}
//	t.Log(createArtworkResponse)
//}
//
//func TestClient_GetArtworkResult(t *testing.T) {
//	getArtworkResultResponse, err := cli.GetArtworkResult(ctx, &GetArtworkResultRequest{
//		Txhash:   "872fdbcfd27f44758385705ee7ffb5784afffdd468e44bbda2714a07c1759b0d",
//		TenantId: "tid-yuhu1",
//	})
//	if err != nil {
//		t.Error(err)
//	}
//	t.Log(getArtworkResultResponse)
//}
//
//func TestClient_GetMintNFTResult(t *testing.T) {
//	getMintNFTResultResponse, err := cli.GetMintNFTResult(ctx, &GetMintNFTResultRequest{
//		TxHash:   "8518e4a79ce346838ab3dc61829ac5d4197a865fb8654e4e84b19eca852d308c",
//		TenantId: "tid-yuhu1",
//	})
//	if err != nil {
//		t.Error(err)
//	}
//	t.Log(getMintNFTResultResponse)
//}
//
//func TestClient_GetTransferNFTResult(t *testing.T) {
//	getTransferNFTResultResponse, err := cli.GetTransferNFTResult(ctx, &GetTransferNFTResultRequest{
//		TxHash:   "dddc644c5467403a99d7ff9702bcb0d93bfb58cc408347b19e69e1f532c3e18a",
//		TenantId: "tid-yuhu1",
//	})
//	if err != nil {
//		t.Error(err)
//	}
//	t.Log(getTransferNFTResultResponse)
//}
//
//func TestClient_MintNFT(t *testing.T) {
//	mintNFTResponse, err := cli.MintNFT(ctx, &MintNFTRequest{
//		ContractAddress: "74a55fb59f51faba6fdc8ac94e1706680cb7b622",
//		ReceiverAddress: "0x7dbd5d3efb0c583257167b1efd58af562053b16c",
//		Signer: &Signer{
//			WalletId:     "wid-llw34jmEw8vE",
//			SignedUserId: "yuhu1",
//		},
//		TenantId:  "tid-yuhu1",
//		Amount:    "3",
//		RequestId: "202401242",
//	})
//	if err != nil {
//		t.Error(err)
//	}
//	t.Log(mintNFTResponse)
//}
//
//func TestClient_TransferNFT(t *testing.T) {
//	transferNFTResponse, err := cli.TransferNFT(ctx, &TransferNFTRequest{
//		ReceiverAddress: "0xec2d2fcf440de78163dc7dc95a46520c0114867b",
//		ContractAddress: "74a55fb59f51faba6fdc8ac94e1706680cb7b622",
//		TokenId:         "10",
//		Signer: &Signer{
//			WalletId:     "wid-llw34jmEw8vE",
//			SignedUserId: "yuhu1",
//		},
//		TenantId:  "tid-yuhu1",
//		RequestId: "202401232",
//	})
//	if err != nil {
//		t.Error(err)
//	}
//	t.Log(transferNFTResponse)
//}
