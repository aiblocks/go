package aiblockstoml

import "github.com/stretchr/testify/mock"

// MockClient is a mockable aiblockstoml client.
type MockClient struct {
	mock.Mock
}

// GetAiBlocksToml is a mocking a method
func (m *MockClient) GetAiBlocksToml(domain string) (*Response, error) {
	a := m.Called(domain)
	return a.Get(0).(*Response), a.Error(1)
}

// GetAiBlocksTomlByAddress is a mocking a method
func (m *MockClient) GetAiBlocksTomlByAddress(address string) (*Response, error) {
	a := m.Called(address)
	return a.Get(0).(*Response), a.Error(1)
}
