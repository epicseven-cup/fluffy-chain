package client

import (
	"context"
	"log"
	"time"

	pb "github.com/epicseven-cup/fluffy-chain/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// methods for client

const (
	addr = "localhost"
)

func SendCreateRequest(source string, destation string, startTime *timestamppb.Timestamp, endTime *timestamppb.Timestamp, status bool) (*pb.CreateRedirectRespond, error) {
	cred := insecure.NewCredentials()
	grpcCred := grpc.WithTransportCredentials(cred)
	conn, err := grpc.NewClient(addr, grpcCred)
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	client := pb.NewRedirectServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	request := pb.CreateRedirectRequest{Source: source, Destation: destation, StartTime: startTime, EndTime: endTime, Status: status}
	respond, err := client.CreateRedirect(ctx, &request)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(respond.GetMessage())
	return respond, err
}

func SendRedirectRequest(requestPath string) {
	cred := insecure.NewCredentials()
	grpcCred := grpc.WithTransportCredentials(cred)
	conn, err := grpc.NewClient(addr, grpcCred)
	if err != nil {
		log.Fatalln(err)
	}

	defer conn.Close()

	client := pb.NewRedirectServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	request := pb.RedirectRequest{Path: requestPath}

	respond, err := client.Redirect(ctx, &request)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(respond.Path)
}
