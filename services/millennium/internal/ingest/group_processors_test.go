//lint:file-ignore U1001 Ignore all unused code, staticcheck doesn't understand testify/suite
package ingest

import (
	"errors"
	"testing"

	"github.com/aiblocks/go/ingest/io"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

var _ millenniumChangeProcessor = (*mockMillenniumChangeProcessor)(nil)

type mockMillenniumChangeProcessor struct {
	mock.Mock
}

func (m *mockMillenniumChangeProcessor) ProcessChange(change io.Change) error {
	args := m.Called(change)
	return args.Error(0)
}

func (m *mockMillenniumChangeProcessor) Commit() error {
	args := m.Called()
	return args.Error(0)
}

var _ millenniumTransactionProcessor = (*mockMillenniumTransactionProcessor)(nil)

type mockMillenniumTransactionProcessor struct {
	mock.Mock
}

func (m *mockMillenniumTransactionProcessor) ProcessTransaction(transaction io.LedgerTransaction) error {
	args := m.Called(transaction)
	return args.Error(0)
}

func (m *mockMillenniumTransactionProcessor) Commit() error {
	args := m.Called()
	return args.Error(0)
}

type GroupChangeProcessorsTestSuiteLedger struct {
	suite.Suite
	processors *groupChangeProcessors
	processorA *mockMillenniumChangeProcessor
	processorB *mockMillenniumChangeProcessor
}

func TestGroupChangeProcessorsTestSuiteLedger(t *testing.T) {
	suite.Run(t, new(GroupChangeProcessorsTestSuiteLedger))
}

func (s *GroupChangeProcessorsTestSuiteLedger) SetupTest() {
	s.processorA = &mockMillenniumChangeProcessor{}
	s.processorB = &mockMillenniumChangeProcessor{}
	s.processors = &groupChangeProcessors{
		s.processorA,
		s.processorB,
	}
}

func (s *GroupChangeProcessorsTestSuiteLedger) TearDownTest() {
	s.processorA.AssertExpectations(s.T())
	s.processorB.AssertExpectations(s.T())
}

func (s *GroupChangeProcessorsTestSuiteLedger) TestProcessChangeFails() {
	change := io.Change{}
	s.processorA.
		On("ProcessChange", change).
		Return(errors.New("transient error")).Once()

	err := s.processors.ProcessChange(change)
	s.Assert().Error(err)
	s.Assert().EqualError(err, "error in *ingest.mockMillenniumChangeProcessor.ProcessChange: transient error")
}

func (s *GroupChangeProcessorsTestSuiteLedger) TestProcessChangeSucceeds() {
	change := io.Change{}
	s.processorA.
		On("ProcessChange", change).
		Return(nil).Once()
	s.processorB.
		On("ProcessChange", change).
		Return(nil).Once()

	err := s.processors.ProcessChange(change)
	s.Assert().NoError(err)
}

func (s *GroupChangeProcessorsTestSuiteLedger) TestCommitFails() {
	s.processorA.
		On("Commit").
		Return(errors.New("transient error")).Once()

	err := s.processors.Commit()
	s.Assert().Error(err)
	s.Assert().EqualError(err, "error in *ingest.mockMillenniumChangeProcessor.Commit: transient error")
}

func (s *GroupChangeProcessorsTestSuiteLedger) TestCommitSucceeds() {
	s.processorA.
		On("Commit").
		Return(nil).Once()
	s.processorB.
		On("Commit").
		Return(nil).Once()

	err := s.processors.Commit()
	s.Assert().NoError(err)
}

type GroupTransactionProcessorsTestSuiteLedger struct {
	suite.Suite
	processors *groupTransactionProcessors
	processorA *mockMillenniumTransactionProcessor
	processorB *mockMillenniumTransactionProcessor
}

func TestGroupTransactionProcessorsTestSuiteLedger(t *testing.T) {
	suite.Run(t, new(GroupTransactionProcessorsTestSuiteLedger))
}

func (s *GroupTransactionProcessorsTestSuiteLedger) SetupTest() {
	s.processorA = &mockMillenniumTransactionProcessor{}
	s.processorB = &mockMillenniumTransactionProcessor{}
	s.processors = &groupTransactionProcessors{
		s.processorA,
		s.processorB,
	}
}

func (s *GroupTransactionProcessorsTestSuiteLedger) TearDownTest() {
	s.processorA.AssertExpectations(s.T())
	s.processorB.AssertExpectations(s.T())
}

func (s *GroupTransactionProcessorsTestSuiteLedger) TestProcessTransactionFails() {
	transaction := io.LedgerTransaction{}
	s.processorA.
		On("ProcessTransaction", transaction).
		Return(errors.New("transient error")).Once()

	err := s.processors.ProcessTransaction(transaction)
	s.Assert().Error(err)
	s.Assert().EqualError(err, "error in *ingest.mockMillenniumTransactionProcessor.ProcessTransaction: transient error")
}

func (s *GroupTransactionProcessorsTestSuiteLedger) TestProcessTransactionSucceeds() {
	transaction := io.LedgerTransaction{}
	s.processorA.
		On("ProcessTransaction", transaction).
		Return(nil).Once()
	s.processorB.
		On("ProcessTransaction", transaction).
		Return(nil).Once()

	err := s.processors.ProcessTransaction(transaction)
	s.Assert().NoError(err)
}

func (s *GroupTransactionProcessorsTestSuiteLedger) TestCommitFails() {
	s.processorA.
		On("Commit").
		Return(errors.New("transient error")).Once()

	err := s.processors.Commit()
	s.Assert().Error(err)
	s.Assert().EqualError(err, "error in *ingest.mockMillenniumTransactionProcessor.Commit: transient error")
}

func (s *GroupTransactionProcessorsTestSuiteLedger) TestCommitSucceeds() {
	s.processorA.
		On("Commit").
		Return(nil).Once()
	s.processorB.
		On("Commit").
		Return(nil).Once()

	err := s.processors.Commit()
	s.Assert().NoError(err)
}
