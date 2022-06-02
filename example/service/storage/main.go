package main

import (
	"context"
	"log"
	"os"

	"github.com/yuhu-tech/qilin-sdk-go/service/storage"
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
	cli, err := storage.NewClient(context.Background(), &storage.Config{AK: Ak, SK: Sk, Endpoint: Endpoint})
	if err != nil {
		log.Fatal(err)
	}

	// upload files
	file1, err := os.Open("./1.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer file1.Close()
	// file2Name, file2 := "file2", bytes.NewReader([]byte("file2-stream"))

	input, err := storage.NewUploadFilesInput(
		TestTenant,
		[]storage.Files{
			{FileName: "0.jpg", Data: file1},
			// {FileName: file2Name, Data: file2},
		})
	if err != nil {
		log.Fatal(err)
	}
	_, err = cli.UploadFiles(context.Background(), input)
	if err != nil {
		log.Fatal(err)
	}
}
