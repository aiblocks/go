package resourceadapter

import (
	"context"

	"github.com/aiblocks/go/amount"
	"github.com/aiblocks/go/protocols/millennium"
	"github.com/aiblocks/go/services/millennium/internal/paths"
)

// PopulatePath converts the paths.Path into a Path
func PopulatePath(ctx context.Context, dest *millennium.Path, p paths.Path) (err error) {
	dest.DestinationAmount = amount.String(p.DestinationAmount)
	dest.SourceAmount = amount.String(p.SourceAmount)

	err = p.Source.Extract(
		&dest.SourceAssetType,
		&dest.SourceAssetCode,
		&dest.SourceAssetIssuer)
	if err != nil {
		return
	}

	err = p.Destination.Extract(
		&dest.DestinationAssetType,
		&dest.DestinationAssetCode,
		&dest.DestinationAssetIssuer)
	if err != nil {
		return
	}

	dest.Path = make([]millennium.Asset, len(p.Path))
	for i, a := range p.Path {
		err = a.Extract(
			&dest.Path[i].Type,
			&dest.Path[i].Code,
			&dest.Path[i].Issuer)
		if err != nil {
			return
		}
	}
	return
}
