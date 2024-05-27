package client

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	pb "github.com/epicseven-cup/fluffy-chain/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// methods for client

const (
	addr    = "localhost:8080"
	logPath = "../client/log/client-log"
)

func SendCreateRequest(source string, destation string, startTime *timestamppb.Timestamp, endTime *timestamppb.Timestamp, status bool) (*pb.CreateRedirectRespond, error) {
	logFile, err := os.OpenFile(logPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println(err)
	}
	logger := log.Default()
	logger.SetOutput(logFile)
	cred := insecure.NewCredentials()
	grpcCred := grpc.WithTransportCredentials(cred)
	conn, err := grpc.NewClient(addr, grpcCred)
	logger.Println(conn.GetState().String())
	logger.Println(conn)
	if err != nil {
		logger.Fatalln(err)
	}
	defer conn.Close()

	client := pb.NewRedirectServiceClient(conn)
	logger.Println(client)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	request := pb.CreateRedirectRequest{Source: source, Destation: destation, StartTime: startTime, EndTime: endTime, Status: status}
	logger.Println(request)
	respond, err := client.CreateRedirect(ctx, &request)
	if err != nil {
		logger.Println("There is an error when you get and respond")
		logger.Fatalln(err)
	}
	logger.Println(respond.GetMessage())
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
