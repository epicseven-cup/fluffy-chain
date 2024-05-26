package test

import (
	"log"

	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	endpoint = ""
	logger = log.Default()
	
)

import "fluffychain/client"

func testCreateRedirectRequest() {
	source:= "/google"
	destation := "https://www.google.com"
	startTime := timestamppb.New()
	endTime := timestamppb.New()
	status := false
	respond := client.SendCreateRequest(source, destation, startTime, endTime, status)

}
