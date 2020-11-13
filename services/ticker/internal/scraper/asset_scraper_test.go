package scraper

import (
	"net/url"
	"testing"

	hProtocol "github.com/aiblocks/go/protocols/millennium"
	"github.com/aiblocks/go/support/errors"
	"github.com/aiblocks/go/support/render/hal"
	"github.com/stretchr/testify/assert"
)

func TestShouldDiscardAsset(t *testing.T) {
	testAsset := hProtocol.AssetStat{
		Amount: "",
	}

	assert.Equal(t, shouldDiscardAsset(testAsset, true), true)

	testAsset = hProtocol.AssetStat{
		Amount: "0.0",
	}
	assert.Equal(t, shouldDiscardAsset(testAsset, true), true)

	testAsset = hProtocol.AssetStat{
		Amount: "0",
	}
	assert.Equal(t, shouldDiscardAsset(testAsset, true), true)

	testAsset = hProtocol.AssetStat{
		Amount:      "123901.0129310",
		NumAccounts: 8,
	}
	assert.Equal(t, shouldDiscardAsset(testAsset, true), true)

	testAsset = hProtocol.AssetStat{
		Amount:      "123901.0129310",
		NumAccounts: 12,
	}
	testAsset.Code = "REMOVE"
	assert.Equal(t, shouldDiscardAsset(testAsset, true), true)

	testAsset = hProtocol.AssetStat{
		Amount:      "123901.0129310",
		NumAccounts: 100,
	}
	testAsset.Code = "SOMETHINGVALID"
	testAsset.Links.Toml.Href = ""
	assert.Equal(t, shouldDiscardAsset(testAsset, true), false)

	testAsset = hProtocol.AssetStat{
		Amount:      "123901.0129310",
		NumAccounts: 40,
	}
	testAsset.Code = "SOMETHINGVALID"
	testAsset.Links.Toml.Href = "http://www.aiblocks.io/.well-known/aiblocks.toml"
	assert.Equal(t, shouldDiscardAsset(testAsset, true), true)

	testAsset = hProtocol.AssetStat{
		Amount:      "123901.0129310",
		NumAccounts: 40,
	}
	testAsset.Code = "SOMETHINGVALID"
	testAsset.Links.Toml.Href = ""
	assert.Equal(t, shouldDiscardAsset(testAsset, true), true)

	testAsset = hProtocol.AssetStat{
		Amount:      "123901.0129310",
		NumAccounts: 40,
	}
	testAsset.Code = "SOMETHINGVALID"
	testAsset.Links.Toml.Href = "https://www.aiblocks.io/.well-known/aiblocks.toml"
	assert.Equal(t, shouldDiscardAsset(testAsset, true), false)
}

func TestDomainsMatch(t *testing.T) {
	tomlURL, _ := url.Parse("https://aiblocks.io/aiblocks.toml")
	orgURL, _ := url.Parse("https://aiblocks.io/")
	assert.True(t, domainsMatch(tomlURL, orgURL))

	tomlURL, _ = url.Parse("https://assets.aiblocks.io/aiblocks.toml")
	orgURL, _ = url.Parse("https://aiblocks.io/")
	assert.False(t, domainsMatch(tomlURL, orgURL))

	tomlURL, _ = url.Parse("https://aiblocks.io/aiblocks.toml")
	orgURL, _ = url.Parse("https://home.aiblocks.io/")
	assert.True(t, domainsMatch(tomlURL, orgURL))

	tomlURL, _ = url.Parse("https://aiblocks.io/aiblocks.toml")
	orgURL, _ = url.Parse("https://home.aiblocks.com/")
	assert.False(t, domainsMatch(tomlURL, orgURL))

	tomlURL, _ = url.Parse("https://aiblocks.io/aiblocks.toml")
	orgURL, _ = url.Parse("https://aiblocks.com/")
	assert.False(t, domainsMatch(tomlURL, orgURL))
}

func TestIsDomainVerified(t *testing.T) {
	tomlURL := "https://aiblocks.io/aiblocks.toml"
	orgURL := "https://aiblocks.io/"
	hasCurrency := true
	assert.True(t, isDomainVerified(orgURL, tomlURL, hasCurrency))

	tomlURL = "https://aiblocks.io/aiblocks.toml"
	orgURL = ""
	hasCurrency = true
	assert.True(t, isDomainVerified(orgURL, tomlURL, hasCurrency))

	tomlURL = ""
	orgURL = ""
	hasCurrency = true
	assert.False(t, isDomainVerified(orgURL, tomlURL, hasCurrency))

	tomlURL = "https://aiblocks.io/aiblocks.toml"
	orgURL = "https://aiblocks.io/"
	hasCurrency = false
	assert.False(t, isDomainVerified(orgURL, tomlURL, hasCurrency))

	tomlURL = "http://aiblocks.io/aiblocks.toml"
	orgURL = "https://aiblocks.io/"
	hasCurrency = true
	assert.False(t, isDomainVerified(orgURL, tomlURL, hasCurrency))

	tomlURL = "https://aiblocks.io/aiblocks.toml"
	orgURL = "http://aiblocks.io/"
	hasCurrency = true
	assert.False(t, isDomainVerified(orgURL, tomlURL, hasCurrency))

	tomlURL = "https://aiblocks.io/aiblocks.toml"
	orgURL = "https://aiblocks.com/"
	hasCurrency = true
	assert.False(t, isDomainVerified(orgURL, tomlURL, hasCurrency))
}

func TestIgnoreInvalidTOMLUrls(t *testing.T) {
	invalidURL := "https:// there is something wrong here.com/aiblocks.toml"
	assetStat := hProtocol.AssetStat{}
	assetStat.Links.Toml = hal.Link{Href: invalidURL}

	_, err := fetchTOMLData(assetStat)

	urlErr, ok := errors.Cause(err).(*url.Error)
	if !ok {
		t.Fatalf("err expected to be a url.Error but was %#v", err)
	}
	assert.Equal(t, "parse", urlErr.Op)
	assert.Equal(t, "https:// there is something wrong here.com/aiblocks.toml", urlErr.URL)
	assert.EqualError(t, urlErr.Err, `invalid character " " in host name`)
}
