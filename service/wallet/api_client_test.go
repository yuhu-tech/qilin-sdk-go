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
//func TestClient_CreateWallet(t *testing.T) {
//	createWalletResponse, err := cli.CreateWallet(ctx, &CreateWalletRequest{
//		SignUserId:      "202401311",
//		TenantId:        "tid-yuhu1",
//		ChainInstanceId: "ChainMakerV2CTC",
//	})
//	if err != nil {
//		t.Error(err)
//	}
//	t.Log(createWalletResponse)
//}
