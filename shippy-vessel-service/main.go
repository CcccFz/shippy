package main

import (
	pb "github.com/CcccFz/shippy/shippy-vessel-service/proto/vessel"
	"github.com/micro/go-micro"
	"log"
	"os"
)

const (
	DefaultHost = "localhost:27017"
)

func main() {

	host := os.Getenv("DB_HOST")
	if host == "" {
		host = DefaultHost
	}
	session, err := CreateSession(host)
	defer session.Close()
	if err != nil {
		log.Fatalf("create session error: %v\n", err)
	}

	// 停留在港口的货船，先写死
	repo := &VesselRepository{session.Copy()}
	CreateDummyData(repo)
	server := micro.NewService(
		micro.Name("go.micro.srv.vessel"),
		micro.Version("latest"),
	)
	server.Init()

	// 将实现服务端的 API 注册到服务端
	pb.RegisterVesselServiceHandler(server.Server(), &handler{session})

	if err := server.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func CreateDummyData(repo Repository) {
	defer repo.Close()
	vessels := []*pb.Vessel{
		{Id: "vessel001", Name: "Boaty McBoatface 1", MaxWeight: 200000, Capacity: 500},
		{Id: "vessel002", Name: "Boaty McBoatface 2", MaxWeight: 210000, Capacity: 600},
		{Id: "vessel003", Name: "Boaty McBoatface 3", MaxWeight: 220000, Capacity: 700},
	}
	for _, v := range vessels {
		repo.Create(v)
	}
}
