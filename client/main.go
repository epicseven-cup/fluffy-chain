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

// This should be what the client is doing

// convert this to be a modualr that can get pull by an another interal service later
const (
	addr      = ""
	source    = ""
	destation = ""
	status    = true
)

var startTime = timestamppb.New(time.Now())
var endTime = timestamppb.New(time.Now())

func main() {
	cred := insecure.NewCredentials()
	grpcCred := grpc.WithTransportCredentials(cred)
	conn, err := grpc.Dial(addr, grpcCred)
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	client := pb.NewCreateRedirectServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	redirectRequest := pb.CreateRedirectRequest{Source: source, Destation: destation, StartTime: startTime, EndTime: endTime, Status: status}
	request, err := client.CreateRedirect(ctx, &redirectRequest)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(request.GetMessage())

}
