package main

import (
	pb "github.com/CcccFz/shippy/shippy-user-service/proto/user"
	"github.com/micro/go-micro"
	"log"
)

func main() {
	// 连接到数据库
	db, err := CreateConnection()
	defer db.Close()

	if err != nil {
		log.Fatalf("connect error: %v\n", err)
	}

	repo := &UserRepository{db}

	// 自动检查 User 结构是否变化
	db.AutoMigrate(&pb.User{})

	// 作者使用了新仓库 shippy-user-service
	// 但 auth.proto 和 user.proto 定义的内容是一致的
	// 修改 shippy.auth 为 go.micro.srv.user
	// 注意 API 调用参数也需对应修改
	srv := micro.NewService(
		micro.Name("go.micro.srv.user"),
		micro.Version("latest"),
	)

	srv.Init()

	// 获取 broker 实例
	// pubSub := s.Server().Options().Broker
	publisher := micro.NewPublisher(topic, srv.Client())

	t := TokenService{repo}
	pb.RegisterUserServiceHandler(srv.Server(), &handler{repo, &t, publisher})

	if err := srv.Run(); err != nil {
		log.Fatalf("user service error: %v\n", err)
	}

}
