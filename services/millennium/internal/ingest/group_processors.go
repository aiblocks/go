package ingest

import (
	"github.com/aiblocks/go/ingest/io"
	"github.com/aiblocks/go/support/errors"
)

type millenniumChangeProcessor interface {
	io.ChangeProcessor
	// TODO maybe rename to Flush()
	Commit() error
}

type groupChangeProcessors []millenniumChangeProcessor

func (g groupChangeProcessors) ProcessChange(change io.Change) error {
	for _, p := range g {
		if err := p.ProcessChange(change); err != nil {
			return errors.Wrapf(err, "error in %T.ProcessChange", p)
		}
	}
	return nil
}

func (g groupChangeProcessors) Commit() error {
	for _, p := range g {
		if err := p.Commit(); err != nil {
			return errors.Wrapf(err, "error in %T.Commit", p)
		}
	}
	return nil
}

type groupTransactionProcessors []millenniumTransactionProcessor

func (g groupTransactionProcessors) ProcessTransaction(tx io.LedgerTransaction) error {
	for _, p := range g {
		if err := p.ProcessTransaction(tx); err != nil {
			return errors.Wrapf(err, "error in %T.ProcessTransaction", p)
		}
	}
	return nil
}

func (g groupTransactionProcessors) Commit() error {
	for _, p := range g {
		if err := p.Commit(); err != nil {
			return errors.Wrapf(err, "error in %T.Commit", p)
		}
	}
	return nil
}
