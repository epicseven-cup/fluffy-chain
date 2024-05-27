package main

import (
	"fluffychain/server"
	"fmt"
	"log"
	"net"
	"os"

	pb "github.com/epicseven-cup/fluffy-chain/api"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

const (
	PORT     = "localhost:8080"
	LOG_PATH = "./log/service-log"
)

func main() {
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	logFile, err := os.OpenFile(LOG_PATH, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println(err)
	}
	logger := log.Default()
	logger.SetOutput(logFile)
	logger.Println("Starting up listener")
	listener, err := net.Listen("tcp", "localhost:8080")
	logger.Println("Listener created")
	if err != nil {
		logger.Fatalln(err)
	}
	logger.Println("Starting up grpc server")
	s := grpc.NewServer()
	logger.Println(listener.Addr())
	logger.Println("it hits here")
	logger.Printf("server listening at %v", listener.Addr())
	logger.Println("grpc server started sucessfully")
	pb.RegisterRedirectServiceServer(s, server.NewRedirectServer())
	// logger.Println("Service registered")
	s.Serve(listener)
}
