package resourceadapter

import (
	"context"
	"fmt"

	"github.com/aiblocks/go/amount"
	protocol "github.com/aiblocks/go/protocols/millennium"
	millenniumContext "github.com/aiblocks/go/services/millennium/internal/context"
	"github.com/aiblocks/go/services/millennium/internal/db2/history"
	"github.com/aiblocks/go/support/errors"
	"github.com/aiblocks/go/support/render/hal"
	"github.com/aiblocks/go/xdr"
)

// PopulateClaimableBalance fills out the resource's fields
func PopulateClaimableBalance(
	ctx context.Context,
	dest *protocol.ClaimableBalance,
	claimableBalance history.ClaimableBalance,
	ledger *history.Ledger,
) error {
	balanceID, err := xdr.MarshalHex(claimableBalance.BalanceID)
	if err != nil {
		return errors.Wrap(err, "marshalling BalanceID")
	}
	dest.BalanceID = balanceID
	dest.Asset = claimableBalance.Asset.StringCanonical()
	dest.Amount = amount.StringFromInt64(int64(claimableBalance.Amount))
	if claimableBalance.Sponsor.Valid {
		dest.Sponsor = claimableBalance.Sponsor.String
	}
	dest.LastModifiedLedger = claimableBalance.LastModifiedLedger
	dest.Claimants = make([]protocol.Claimant, len(claimableBalance.Claimants))
	for i, c := range claimableBalance.Claimants {
		dest.Claimants[i].Destination = c.Destination
		dest.Claimants[i].Predicate = c.Predicate
	}

	if ledger != nil {
		dest.LastModifiedTime = &ledger.ClosedAt
	}

	lb := hal.LinkBuilder{Base: millenniumContext.BaseURL(ctx)}
	self := fmt.Sprintf("/claimable_balances/%s", dest.BalanceID)
	dest.Links.Self = lb.Link(self)
	dest.PT = fmt.Sprintf("%d-%s", claimableBalance.LastModifiedLedger, dest.BalanceID)
	return nil
}
