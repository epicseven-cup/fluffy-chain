package fluffychain

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/epicseven-cup/fluffy-chain/CreateRedirect"
	pb "github.com/epicseven-cup/fluffy-chain/CreateRedirect"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/protobuf/proto"
)

const (
	uri     = ""
	logPath = ""
)

type RedirectServer struct {
	pb.UnimplementedCreateRedirectServiceServer
	logger *log.Logger
}

func NewRedirectServer() *RedirectServer {
	logFile, err := os.OpenFile(logPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println(err)
	}
	logger := log.Default()
	logger.SetOutput(logFile)
	return &RedirectServer{logger: logger}
}

func (redirectServer *RedirectServer) CreateRedirect(ctx context.Context, in *pb.CreateRedirectRequest) (pb.CreateRedirectRespond, error) {
	redirectServer.logger.Println("Starting to CreateRedirect service")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	redirectServer.logger.Println("Client is standardized correctly")
	if err != nil {
		redirectServer.logger.Fatalln(err)
	}
	cursor := client.Database("redirect").Collection("url")
	redirectServer.logger.Println("Cursor is standardized correctly")
	url := proto.Unmarshal(in)
	cursor.InsertOne()
	return CreateRedirect.CreateRedirectRespond{}, nil
}
