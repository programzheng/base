package invokegrpc

import (
	"context"
	"log"
	"time"

	"github.com/programzheng/base/internal/grpc/entity"
	pb "github.com/programzheng/base/internal/grpc/proto"
	"github.com/programzheng/base/pkg/service/game"
	"github.com/programzheng/base/pkg/service/user"
)

func RandomTicket(count int) {
	conn, err := game.GetGamesGRPCConnection()
	if err != nil {
		log.Printf("could not get games grpc connection: %v", err)
	}
	defer conn.Close()
	c, err := game.GetGamesGRPCClient(conn)
	if err != nil {
		log.Printf("could not get games grpc client: %v", err)
	}

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.RandomTicket(ctx, &pb.RandomTicketRequest{Count: int64(count)})
	if err != nil {
		log.Printf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}

func AssignRandomIssuedTicketToThirdPartyUser(agentCode string, userUUID string) (*entity.UserTicket, error) {
	user, err := user.GetUserByUUID(userUUID)
	if err != nil {
		log.Printf("could not get user by uuid: %v", err)
		return nil, err
	}

	conn, err := game.GetGamesGRPCConnection()
	if err != nil {
		log.Printf("could not get games grpc connection: %v", err)
		return nil, err
	}
	defer conn.Close()
	c, err := game.GetGamesGRPCClient(conn)
	if err != nil {
		log.Printf("could not get games grpc client: %v", err)
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.AssignRandomIssuedTicketToThirdPartyUser(
		ctx,
		&pb.AssignRandomIssuedTicketToThirdPartyUserRequest{
			Code:         agentCode,
			ThirdPartyID: user.UUID,
		},
	)
	if err != nil {
		log.Printf("could not greet: %v", err)
		return nil, err
	}
	grpcUserTicket := r.GetUserTicket()

	userTicket := entity.UserTicket{
		Code: grpcUserTicket.Code,
		Name: grpcUserTicket.Name,
	}

	return &userTicket, nil
}

func GetIssuedUserTicketsByAgentCode(agentCode string) {
	conn, err := game.GetGamesGRPCConnection()
	if err != nil {
		log.Printf("could not get games grpc connection: %v", err)
	}
	defer conn.Close()
	c, err := game.GetGamesGRPCClient(conn)
	if err != nil {
		log.Printf("could not get games grpc client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetIssuedUserTicketsByAgentCode(
		ctx,
		&pb.GetIssuedUserTicketsByAgentCodeRequest{
			Code: agentCode,
		},
	)
	if err != nil {
		log.Printf("could not greet: %v", err)
	}

	log.Printf("Greeting: %v", r.GetUserTickets())
}
