package test

import (
	"fluffychain/client"
	"testing"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	endpoint = ""
)

func TestCreateRedirectRequest(t *testing.T) {
	source := "/google"
	destation := "https://www.google.com"
	startTime := timestamppb.New(time.Time{})
	endTime := timestamppb.New(time.Time{})
	status := false
	respond, err := client.SendCreateRequest(source, destation, startTime, endTime, status)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(respond.GetMessage())
}
