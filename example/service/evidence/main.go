package main

import (
	"context"
	"fmt"
	"log"

	"github.com/yuhu-tech/qilin-sdk-go/service/evidence"
)

const (
	TestTenant = "tid-yuhu1"
	Ak         = "test-ak"
	Sk         = "test-sk"
	// dev endpoint
	Endpoint = "localhost:10000"
)

func main() {
	// init client
	cli, err := evidence.NewClient(context.Background(), &evidence.Config{AK: Ak, SK: Sk, TenantId: TestTenant, Endpoint: Endpoint})
	if err != nil {
		log.Fatal(err)
	}

	resp, err := cli.CreateStructuredEvidence(context.Background(), &evidence.CreateStructuredEvidenceRequest{
		ModelName: "jingkuang-nft",
		NameValues: []*evidence.KeyValuePair{
			{
				Key:   "contractaddress",
				Value: "0x13FF557D9578f2Fd78fF47383184A70e685b600f",
			},
			{
				Key:   "zichanmingcheng",
				Value: "雨水卡",
			},
			{
				Key:   "chuangzuozhe",
				Value: "版易官方平台-004",
			},
			{
				Key:   "faxingliang",
				Value: "100",
			},
			{
				Key:   "faxingshijian",
				Value: "2023-07-17 18:55:15",
			},
			{
				Key:   "liulanqidizhi",
				Value: "https://explorer.jingkuang.info/address/0x13FF557D9578f2Fd78fF47383184A70e685b600f",
			},
		},
		TenantId: TestTenant,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("[CreateStructuredEvidence] success: evidenceId %+v\n", resp.EvidenceId)
}
