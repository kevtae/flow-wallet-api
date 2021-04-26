package store

import (
	"context"

	"github.com/onflow/flow-go-sdk"
	"github.com/onflow/flow-go-sdk/client"
	"github.com/onflow/flow-go-sdk/crypto"
)

type KeyStore interface {
	Generate(ctx context.Context, keyIndex int, weight int) (NewKeyWrapper, error)
	Save(AccountKey) error
	ServiceAuthorizer(ctx context.Context, fc *client.Client) (Authorizer, error)
	AccountAuthorizer(ctx context.Context, fc *client.Client, addr flow.Address) (Authorizer, error)
}

type Authorizer struct {
	Address flow.Address
	Key     *flow.AccountKey
	Signer  crypto.Signer
}

type NewKeyWrapper struct {
	FlowKey    *flow.AccountKey
	AccountKey AccountKey
}