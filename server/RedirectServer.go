package server

import (
	"context"
	"fmt"
	"log"
	"os"

	pb "github.com/epicseven-cup/fluffy-chain/api"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	uri     = "mongodb://database:27017"
	logPath = "./server/log/server-log"
)

type RedirectServer struct {
	pb.UnimplementedRedirectServiceServer
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

	// Checks if this path already exist in the database
	query := bson.D{{"destation", in.GetDestation()}}
	err = cursor.FindOne(context.TODO(), query).Decode(nil)
	if err != mongo.ErrNoDocuments {
		// there is already a doc
		redirectServer.logger.Fatalln(err)
		return nil, nil
	}

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

func (redirectServer *RedirectServer) Redirect(ctx context.Context, in *pb.RedirectRequest) (*pb.RedirectRespond, error) {
	redirectServer.logger.Println("Starting a Redirect service")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	redirectServer.logger.Println("Client is create successfuly")
	if err != nil {
		redirectServer.logger.Fatal(err)
	}
	cursor := client.Database("redirect").Collection("url")
	redirectServer.logger.Println("Cursor is standardized correctly")
	// result of the redirect request
	var result pb.CreateRedirectRequest
	// Select on the url or request a default locaiton
	redirectPath := in.GetPath()
	query := bson.D{{Key: "source", Value: redirectPath}}
	err = cursor.FindOne(context.TODO(), query).Decode(&result)
	if err != nil {
		redirectServer.logger.Fatal(err)
	}
	redirectServer.logger.Println("FindOne results are done")

	dest := result.GetDestation()
	respond := pb.RedirectRespond{Path: dest}
	return &respond, nil
}
