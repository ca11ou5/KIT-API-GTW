package auth

import (
	"API_Gateway/configs"
	"API_Gateway/internal/auth/pb"
	"google.golang.org/grpc"
	"log"
)

type ServiceClient struct {
	AuthClient pb.AuthServiceClient
}

func InitAuthClient(cfg *configs.Config) pb.AuthServiceClient {
	cc, err := grpc.Dial(cfg.AuthSvcUrl, grpc.WithInsecure())
	if err != nil {
		log.Fatal("failed to establish connection with auth svc: " + err.Error())
	}

	return pb.NewAuthServiceClient(cc)
}
