package server

import (
	"context"
	"fmt"
	"log"
	"os"

	pb "github.com/epicseven-cup/fluffy-chain/api"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (redirectServer *RedirectServer) CreateRedirect(ctx context.Context, in *pb.CreateRedirectRequest) (*pb.CreateRedirectRespond, error) {
	redirectServer.logger.Println("Starting to CreateRedirect service")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	redirectServer.logger.Println("Client is standardized correctly")
	if err != nil {
		redirectServer.logger.Fatalln(err)
	}
	cursor := client.Database("redirect").Collection("url")
	redirectServer.logger.Println("Cursor is standardized correctly")
	redirectServer.logger.Println("Starting to Insert new redirect url into the Mongo Database")
	result, err := cursor.InsertOne(context.TODO(), in)
	redirectServer.logger.Println("The cursor has inserted the new reidrect url here is the result: ", result)
	status := true
	message := "Success"
	if err != nil {
		redirectServer.logger.Fatalln(err)
		message = "Fail to Insert"
		status = false
	}
	startTime := in.GetStartTime()
	endTime := in.GetEndTime()
	respond := pb.CreateRedirectRespond{Status: status, Message: message, StartTime: startTime, EndTime: endTime}
	return &respond, nil
}
