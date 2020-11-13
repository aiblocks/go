package aiblockstoml

import (
	"strings"
	"testing"

	"net/http"

	"github.com/aiblocks/go/support/http/httptest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClientURL(t *testing.T) {
	//HACK:  we're testing an internal method rather than setting up a http client
	//mock.

	c := &Client{UseHTTP: false}
	assert.Equal(t, "https://aiblocks.io/.well-known/aiblocks.toml", c.url("aiblocks.io"))

	c = &Client{UseHTTP: true}
	assert.Equal(t, "http://aiblocks.io/.well-known/aiblocks.toml", c.url("aiblocks.io"))
}

func TestClient(t *testing.T) {
	h := httptest.NewClient()
	c := &Client{HTTP: h}

	// happy path
	h.
		On("GET", "https://aiblocks.io/.well-known/aiblocks.toml").
		ReturnString(http.StatusOK,
			`FEDERATION_SERVER="https://localhost/federation"`,
		)
	stoml, err := c.GetAiBlocksToml("aiblocks.io")
	require.NoError(t, err)
	assert.Equal(t, "https://localhost/federation", stoml.FederationServer)

	// aiblocks.toml exceeds limit
	h.
		On("GET", "https://toobig.org/.well-known/aiblocks.toml").
		ReturnString(http.StatusOK,
			`FEDERATION_SERVER="https://localhost/federation`+strings.Repeat("0", AiBlocksTomlMaxSize)+`"`,
		)
	stoml, err = c.GetAiBlocksToml("toobig.org")
	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "aiblocks.toml response exceeds")
	}

	// not found
	h.
		On("GET", "https://missing.org/.well-known/aiblocks.toml").
		ReturnNotFound()
	stoml, err = c.GetAiBlocksToml("missing.org")
	assert.EqualError(t, err, "http request failed with non-200 status code")

	// invalid toml
	h.
		On("GET", "https://json.org/.well-known/aiblocks.toml").
		ReturnJSON(http.StatusOK, map[string]string{"hello": "world"})
	stoml, err = c.GetAiBlocksToml("json.org")

	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "toml decode failed")
	}
}
