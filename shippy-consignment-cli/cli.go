package main

import (
	"context"
	"encoding/json"
	"errors"
	pb "github.com/CcccFz/shippy/shippy-consignment-service/proto/consignment"
	microclient "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
	"github.com/micro/go-micro/metadata"
	"io/ioutil"
	"log"
)

const (
	DefaultInfoFile = "shippy-consignment-cli/consignment.json"
)

// 读取 consignment.json 中记录的货物信息
func parseFile(fileName string) (*pb.Consignment, error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	var consignment *pb.Consignment
	err = json.Unmarshal(data, &consignment)
	if err != nil {
		return nil, errors.New("consignment.json file content error")
	}
	return consignment, nil
}

func main() {

	cmd.Init()

	// 创建微服务的客户端，简化了手动 Dial 连接服务端的步骤
	client := pb.NewShippingServiceClient("go.micro.srv.consignment", microclient.DefaultClient)

	// 在命令行中指定新的货物信息 json 件

	infoFile := DefaultInfoFile
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjp7ImlkIjoiZGQ3NTI1ZWEtOWM1Yy00ODcwLWI1ZDgtMmY3MzNiYWI4OGVhIiwibmFtZSI6IkV3YW4gVmFsZW50aW5lIiwiY29tcGFueSI6IkJCQyIsImVtYWlsIjoiZXdhbi52YWxlbnRpbmU4OUBnbWFpbC5jb20iLCJwYXNzd29yZCI6IiQyYSQxMCR5SHBjSk9rR3NGcTRjYmdUODV2WHQuLi5YZ1A1ZnZBR0VXR3l6UzB5TXdBUU9KdmFiVUVSVyJ9LCJleHAiOjE1NDE0ODY3NjcsImlzcyI6ImdvLm1pY3JvLnNydi51c2VyIn0.JCCsG5TQQ1FOQTnGHAIevlrQnsWwTeUjVvKn2LZdTns"

	// 解析货物信息
	consignment, err := parseFile(infoFile)
	if err != nil {
		log.Fatalf("parse info file error: %v", err)
	}

	// 创建带有用户 token 的 context
	// consignment-service 服务端将从中取出 token，解密取出用户身份
	tokenContext := metadata.NewContext(context.Background(), map[string]string{
		"token": token,
	})

	// 调用 RPC
	// 将货物存储到指定用户的仓库里
	resp, err := client.CreateConsignment(tokenContext, consignment)
	if err != nil {
		log.Fatalf("create consignment error: %v", err)
	}
	log.Printf("created: %t", resp.Created)

	// 列出目前所有托运的货物
	resp, err = client.GetConsignments(tokenContext, &pb.GetRequest{})
	if err != nil {
		log.Fatalf("failed to list consignments: %v", err)
	}
	for i, c := range resp.Consignments {
		log.Printf("consignment_%d: %v\n", i, c)
	}
}
