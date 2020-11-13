package federation

import (
	"errors"
	"net/http"
	"net/url"
	"strings"
	"testing"

	hc "github.com/aiblocks/go/clients/millenniumclient"
	"github.com/aiblocks/go/clients/aiblockstoml"
	"github.com/aiblocks/go/support/http/httptest"
	"github.com/stretchr/testify/assert"
)

func TestLookupByAddress(t *testing.T) {
	hmock := httptest.NewClient()
	tomlmock := &aiblockstoml.MockClient{}
	c := &Client{AiBlocksTOML: tomlmock, HTTP: hmock}

	// happy path - string integer
	tomlmock.On("GetAiBlocksToml", "aiblocks.io").Return(&aiblockstoml.Response{
		FederationServer: "https://aiblocks.io/federation",
	}, nil)
	hmock.On("GET", "https://aiblocks.io/federation").
		ReturnJSON(http.StatusOK, map[string]string{
			"aiblocks_address": "scott*aiblocks.io",
			"account_id":      "GASTNVNLHVR3NFO3QACMHCJT3JUSIV4NBXDHDO4VTPDTNN65W3B2766C",
			"memo_type":       "id",
			"memo":            "123",
		})
	resp, err := c.LookupByAddress("scott*aiblocks.io")

	if assert.NoError(t, err) {
		assert.Equal(t, "GASTNVNLHVR3NFO3QACMHCJT3JUSIV4NBXDHDO4VTPDTNN65W3B2766C", resp.AccountID)
		assert.Equal(t, "id", resp.MemoType)
		assert.Equal(t, "123", resp.Memo.String())
	}

	// happy path - integer
	tomlmock.On("GetAiBlocksToml", "aiblocks.io").Return(&aiblockstoml.Response{
		FederationServer: "https://aiblocks.io/federation",
	}, nil)
	hmock.On("GET", "https://aiblocks.io/federation").
		ReturnJSON(http.StatusOK, map[string]interface{}{
			"aiblocks_address": "scott*aiblocks.io",
			"account_id":      "GASTNVNLHVR3NFO3QACMHCJT3JUSIV4NBXDHDO4VTPDTNN65W3B2766C",
			"memo_type":       "id",
			"memo":            123,
		})
	resp, err = c.LookupByAddress("scott*aiblocks.io")

	if assert.NoError(t, err) {
		assert.Equal(t, "GASTNVNLHVR3NFO3QACMHCJT3JUSIV4NBXDHDO4VTPDTNN65W3B2766C", resp.AccountID)
		assert.Equal(t, "id", resp.MemoType)
		assert.Equal(t, "123", resp.Memo.String())
	}

	// happy path - string
	tomlmock.On("GetAiBlocksToml", "aiblocks.io").Return(&aiblockstoml.Response{
		FederationServer: "https://aiblocks.io/federation",
	}, nil)
	hmock.On("GET", "https://aiblocks.io/federation").
		ReturnJSON(http.StatusOK, map[string]interface{}{
			"aiblocks_address": "scott*aiblocks.io",
			"account_id":      "GASTNVNLHVR3NFO3QACMHCJT3JUSIV4NBXDHDO4VTPDTNN65W3B2766C",
			"memo_type":       "text",
			"memo":            "testing",
		})
	resp, err = c.LookupByAddress("scott*aiblocks.io")

	if assert.NoError(t, err) {
		assert.Equal(t, "GASTNVNLHVR3NFO3QACMHCJT3JUSIV4NBXDHDO4VTPDTNN65W3B2766C", resp.AccountID)
		assert.Equal(t, "text", resp.MemoType)
		assert.Equal(t, "testing", resp.Memo.String())
	}

	// response exceeds limit
	tomlmock.On("GetAiBlocksToml", "toobig.org").Return(&aiblockstoml.Response{
		FederationServer: "https://toobig.org/federation",
	}, nil)
	hmock.On("GET", "https://toobig.org/federation").
		ReturnJSON(http.StatusOK, map[string]string{
			"aiblocks_address": strings.Repeat("0", FederationResponseMaxSize) + "*aiblocks.io",
			"account_id":      "GASTNVNLHVR3NFO3QACMHCJT3JUSIV4NBXDHDO4VTPDTNN65W3B2766C",
			"memo_type":       "id",
			"memo":            "123",
		})
	_, err = c.LookupByAddress("response*toobig.org")
	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "federation response exceeds")
	}

	// failed toml resolution
	tomlmock.On("GetAiBlocksToml", "missing.org").Return(
		(*aiblockstoml.Response)(nil),
		errors.New("toml failed"),
	)
	_, err = c.LookupByAddress("scott*missing.org")
	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "toml failed")
	}

	// 404 federation response
	tomlmock.On("GetAiBlocksToml", "404.org").Return(&aiblockstoml.Response{
		FederationServer: "https://404.org/federation",
	}, nil)
	hmock.On("GET", "https://404.org/federation").ReturnNotFound()
	_, err = c.LookupByAddress("scott*404.org")
	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "failed with (404)")
	}

	// connection error on federation response
	tomlmock.On("GetAiBlocksToml", "error.org").Return(&aiblockstoml.Response{
		FederationServer: "https://error.org/federation",
	}, nil)
	hmock.On("GET", "https://error.org/federation").ReturnError("kaboom!")
	_, err = c.LookupByAddress("scott*error.org")
	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "kaboom!")
	}
}

func TestLookupByID(t *testing.T) {
	millenniumMock := &hc.MockClient{}
	client := &Client{Millennium: millenniumMock}

	millenniumMock.On("HomeDomainForAccount", "GASTNVNLHVR3NFO3QACMHCJT3JUSIV4NBXDHDO4VTPDTNN65W3B2766C").
		Return("", errors.New("homedomain not set"))

	// an account without a homedomain set fails
	_, err := client.LookupByAccountID("GASTNVNLHVR3NFO3QACMHCJT3JUSIV4NBXDHDO4VTPDTNN65W3B2766C")
	assert.Error(t, err)
	assert.Equal(t, "get homedomain failed: homedomain not set", err.Error())
}

func TestForwardRequest(t *testing.T) {
	hmock := httptest.NewClient()
	tomlmock := &aiblockstoml.MockClient{}
	c := &Client{AiBlocksTOML: tomlmock, HTTP: hmock}

	// happy path - string integer
	tomlmock.On("GetAiBlocksToml", "aiblocks.io").Return(&aiblockstoml.Response{
		FederationServer: "https://aiblocks.io/federation",
	}, nil)
	hmock.On("GET", "https://aiblocks.io/federation").
		ReturnJSON(http.StatusOK, map[string]string{
			"account_id": "GASTNVNLHVR3NFO3QACMHCJT3JUSIV4NBXDHDO4VTPDTNN65W3B2766C",
			"memo_type":  "id",
			"memo":       "123",
		})
	fields := url.Values{}
	fields.Add("federation_type", "bank_account")
	fields.Add("swift", "BOPBPHMM")
	fields.Add("acct", "2382376")
	resp, err := c.ForwardRequest("aiblocks.io", fields)

	if assert.NoError(t, err) {
		assert.Equal(t, "GASTNVNLHVR3NFO3QACMHCJT3JUSIV4NBXDHDO4VTPDTNN65W3B2766C", resp.AccountID)
		assert.Equal(t, "id", resp.MemoType)
		assert.Equal(t, "123", resp.Memo.String())
	}
}

func Test_url(t *testing.T) {
	c := &Client{}

	// forward requests
	qstr := url.Values{}
	qstr.Add("type", "forward")
	qstr.Add("federation_type", "bank_account")
	qstr.Add("swift", "BOPBPHMM")
	qstr.Add("acct", "2382376")
	furl := c.url("https://aiblocks.io/federation", qstr)
	assert.Equal(t, "https://aiblocks.io/federation?acct=2382376&federation_type=bank_account&swift=BOPBPHMM&type=forward", furl)

	// regression: ensure that query is properly URI encoded
	qstr = url.Values{}
	qstr.Add("type", "q")
	qstr.Add("q", "scott+receiver1@aiblocks.io*aiblocks.io")
	furl = c.url("", qstr)
	assert.Equal(t, "?q=scott%2Breceiver1%40aiblocks.io%2Aaiblocks.io&type=q", furl)
}
