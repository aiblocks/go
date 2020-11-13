package federation

import (
	"net/http"
	"net/url"

	hc "github.com/aiblocks/go/clients/millenniumclient"
	"github.com/aiblocks/go/clients/aiblockstoml"
	proto "github.com/aiblocks/go/protocols/federation"
)

// FederationResponseMaxSize is the maximum size of response from a federation server
const FederationResponseMaxSize = 100 * 1024

// DefaultTestNetClient is a default federation client for testnet
var DefaultTestNetClient = &Client{
	HTTP:        http.DefaultClient,
	Millennium:     hc.DefaultTestNetClient,
	AiBlocksTOML: aiblockstoml.DefaultClient,
}

// DefaultPublicNetClient is a default federation client for pubnet
var DefaultPublicNetClient = &Client{
	HTTP:        http.DefaultClient,
	Millennium:     hc.DefaultPublicNetClient,
	AiBlocksTOML: aiblockstoml.DefaultClient,
}

// Client represents a client that is capable of resolving a federation request
// using the internet.
type Client struct {
	AiBlocksTOML AiBlocksTOML
	HTTP        HTTP
	Millennium     Millennium
	AllowHTTP   bool
}

type ClientInterface interface {
	LookupByAddress(addy string) (*proto.NameResponse, error)
	LookupByAccountID(aid string) (*proto.IDResponse, error)
	ForwardRequest(domain string, fields url.Values) (*proto.NameResponse, error)
}

// Millennium represents a millennium client that can be consulted for data when
// needed as part of the federation protocol
type Millennium interface {
	HomeDomainForAccount(aid string) (string, error)
}

// HTTP represents the http client that a federation client uses to make http
// requests.
type HTTP interface {
	Get(url string) (*http.Response, error)
}

// AiBlocksTOML represents a client that can resolve a given domain name to
// aiblocks.toml file.  The response is used to find the federation server that a
// query should be made against.
type AiBlocksTOML interface {
	GetAiBlocksToml(domain string) (*aiblockstoml.Response, error)
}

// confirm interface conformity
var _ AiBlocksTOML = aiblockstoml.DefaultClient
var _ HTTP = http.DefaultClient
var _ ClientInterface = &Client{}
