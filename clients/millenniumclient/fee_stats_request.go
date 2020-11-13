package millenniumclient

import "github.com/aiblocks/go/support/errors"

// BuildURL returns the url for getting fee stats about a running millennium instance
func (fr feeStatsRequest) BuildURL() (endpoint string, err error) {
	endpoint = fr.endpoint
	if endpoint == "" {
		err = errors.New("invalid request: too few parameters")
	}

	return
}
