/**
 * Created by zhouwenzhe on 2024/1/15
 */

package assets

//
//import (
//	"context"
//	"testing"
//)
//
//const (
//	TestTenant = "tid-yuhu1"
//	Ak         = "test-ak"
//	Sk         = "test-sk"
//	// dev endpoint
//	Endpoint = "localhost:10000"
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
//func TestGetArtworkResultRequest_Payload(t *testing.T) {
//	getArtworkResultResponse, err := cli.GetArtworkResult(ctx, &GetArtworkResultRequest{
//		Txhash:   "872fdbcfd27f44758385705ee7ffb5784afffdd468e44bbda2714a07c1759b0d",
//		TenantId: "tid-yuhu1",
//	})
//	if err != nil {
//		t.Error(err)
//	}
//	t.Log(getArtworkResultResponse)
//}
