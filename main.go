package fluffychain

import (
	"fmt"
	"log"
	"net"
	"os"

	"fluffychain/server"

	pb "github.com/epicseven-cup/fluffy-chain/api"
	"google.golang.org/grpc"
)

const (
	PORT     = ""
	LOG_PATH = ""
)

func main() {
	logFile, err := os.OpenFile(LOG_PATH, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println(err)
	}
	logger := log.Default()
	logger.SetOutput(logFile)

	listener, err := net.Listen("tcp", PORT)

	if err != nil {
		logger.Fatalln(err)
	}

	s := grpc.NewServer()
	pb.RegisterRedirectServiceServer(s, server.NewRedirectServer())
	logger.Println("Service registered")
	if err := s.Serve(listener); err != nil {
		logger.Fatalln(err)
	}
}
