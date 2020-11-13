// Package meta provides helpers for processing the metadata that is produced by
// aiblocks-core while processing transactions.
package meta

import "github.com/aiblocks/go/xdr"

// Bundle represents all of the metadata emitted from the application of a single
// aiblocks transaction; Both fee meta and result meta is included.
type Bundle struct {
	FeeMeta         xdr.LedgerEntryChanges
	TransactionMeta xdr.TransactionMeta
}
