package game

import (
	"github.com/programzheng/base/config"
	pb "github.com/programzheng/base/internal/grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func getGamesGRPCAddress() string {
	return config.Cfg.GetString("GAMES_GRPC_ADDRESS")
}

func GetGamesGRPCConnection() (*grpc.ClientConn, error) {
	addr := getGamesGRPCAddress()
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func GetGamesGRPCClient(conn *grpc.ClientConn) (pb.GreeterClient, error) {
	return pb.NewGreeterClient(conn), nil
}
