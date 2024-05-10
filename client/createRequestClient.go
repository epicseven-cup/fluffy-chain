package client

import (
	"context"

	"github.com/epicseven-cup/fluffy-chain/CreateRedirect"
	"google.golang.org/grpc"
)

type CreateRequestClient struct {
}

func (createRequestClient *CreateRequestClient) NewCreateRedirectServiceClient() CreateRedirect.CreateRedirectServiceClient {
	return nil
}

func (createRequestClient *CreateRequestClient) CreateRedirect(ctx context.Context, in *CreateRedirectRequest, opts ...grpc.CallOption) (*CreateRedirect.CreateRedirectRespond, error) {
	return nil, nil
}
