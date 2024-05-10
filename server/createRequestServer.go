package fluffychain

import (
	"context"

	"github.com/epicseven-cup/fluffy-chain/CreateRedirect"
)

type CreateRedirectService struct {
}

func (createRedirectService *CreateRedirectService) NewCreateRedirectServiceClient() CreateRedirect.CreateRedirectServiceServer {
	return nil
}

func (createRedirectService *CreateRedirectService) CreateRedirect(context.Context, *CreateRedirect.CreateRedirectRequest) (CreateRedirect.CreateRedirectRespond, error) {
	return CreateRedirect.CreateRedirectRespond{}, nil
}
