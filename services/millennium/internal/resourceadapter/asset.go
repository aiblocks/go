package resourceadapter

import (
	"context"

	protocol "github.com/aiblocks/go/protocols/millennium"
	"github.com/aiblocks/go/xdr"
)

func PopulateAsset(ctx context.Context, dest *protocol.Asset, asset xdr.Asset) error {
	return asset.Extract(&dest.Type, &dest.Code, &dest.Issuer)
}
