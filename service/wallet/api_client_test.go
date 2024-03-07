/**
 * Created by zhouwenzhe on 2024/1/15
 */

package wallet

// import (
// 	"context"
// 	"testing"

// )

// const (
// 	TestTenant = "tid-yuhu1"
// 	Ak         = "test-ak"
// 	Sk         = "test-sk"
// 	// Endpoint   = "119.3.106.151:10100"
// 	Endpoint = "127.0.0.1:10000"
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

// func TestClient_CreateWallet(t *testing.T) {
// 	createWalletResponse, err := cli.CreateWallet(ctx, &CreateWalletRequest{
// 		SignUserId:      "cl11xbqem6md90764cnm3vs2u",
// 		TenantId:        "tid-yuhu1",
// 		ChainInstanceId: "ChainMakerV2CTC",
// 	})
// 	t.Log(err==ErrDuplicateWallet)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	t.Log(createWalletResponse)
// }

// func TestClient_GetWallet(t *testing.T) {
// 	getWalletResponse, err := cli.GetWallet(ctx, &GetWalletRequest{
// 		OutUserId:       "cl11xbqem6md90764cnm3vs2u-",
// 		ChainInstanceId: "ChainMakerV2CTC",
// 		TenantId:        "tid-yuhu1",
// 	})
// 	t.Log(err==ErrNotFoundWallet)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	t.Log(getWalletResponse)
// }
