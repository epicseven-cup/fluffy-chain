package fluffychain

import (
	"context"
	"log"
	"net"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	DATABASE_URI = ""
	LISTEN_PORT  = ""
	LOG_PATH     = ""
)

var cursor = createCursor()
var logger = log.Default()

func main() {
	ln, err := net.Listen("tcp", LISTEN_PORT)
	if err != nil {
		logger.Fatalln(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			logger.Fatalln(err)
		}
		go requestHandle(conn)
	}

}

func requestHandle(conn net.Conn) {
	panic("unimplemented")

}

func createCursor() *mongo.Client {
	logger.Println("Starting to create mongo client")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(DATABASE_URI))

	if err != nil {
		logger.Fatalln(err)
	}

	logger.Println("Finish creating mongo client")

	return client
}
