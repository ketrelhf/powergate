package wallet

import (
	"context"
	"math/big"

	"github.com/ipfs/go-cid"
)

// Module provides wallet management access to a Filecoin client.
type Module interface {
	NewAddress(ctx context.Context, typ string) (string, error)
	List(ctx context.Context) ([]string, error)
	Balance(ctx context.Context, addr string) (*big.Int, error)
	SendFil(ctx context.Context, from string, to string, amount *big.Int) (cid.Cid, error)
	FundFromFaucet(ctx context.Context, addr string) error
}
