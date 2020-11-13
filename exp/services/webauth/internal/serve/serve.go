package serve

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/aiblocks/go/clients/millenniumclient"
	"github.com/aiblocks/go/keypair"
	"github.com/aiblocks/go/support/errors"
	supporthttp "github.com/aiblocks/go/support/http"
	supportlog "github.com/aiblocks/go/support/log"
	"github.com/aiblocks/go/support/render/health"
	"gopkg.in/square/go-jose.v2"
)

type Options struct {
	Logger                      *supportlog.Entry
	MillenniumURL                  string
	Port                        int
	NetworkPassphrase           string
	SigningKeys                 string
	AuthHomeDomains             string
	ChallengeExpiresIn          time.Duration
	JWK                         string
	JWTIssuer                   string
	JWTExpiresIn                time.Duration
	AllowAccountsThatDoNotExist bool
}

func Serve(opts Options) {
	handler, err := handler(opts)
	if err != nil {
		opts.Logger.Fatalf("Error: %v", err)
		return
	}

	addr := fmt.Sprintf(":%d", opts.Port)
	supporthttp.Run(supporthttp.Config{
		ListenAddr: addr,
		Handler:    handler,
		OnStarting: func() {
			opts.Logger.Info("Starting SEP-10 Web Authentication Server")
			opts.Logger.Infof("Listening on %s", addr)
		},
	})
}

func handler(opts Options) (http.Handler, error) {
	signingKeys := []*keypair.Full{}
	signingKeyStrs := strings.Split(opts.SigningKeys, ",")
	signingAddresses := make([]*keypair.FromAddress, 0, len(signingKeyStrs))

	for i, signingKeyStr := range signingKeyStrs {
		signingKey, err := keypair.ParseFull(signingKeyStr)
		if err != nil {
			return nil, errors.Wrap(err, "parsing signing key seed")
		}
		signingKeys = append(signingKeys, signingKey)
		signingAddresses = append(signingAddresses, signingKey.FromAddress())
		opts.Logger.Info("Signing key ", i, ": ", signingKey.Address())
	}

	homeDomains := strings.Split(opts.AuthHomeDomains, ",")
	trimmedHomeDomains := make([]string, 0, len(homeDomains))
	for _, homeDomain := range homeDomains {
		// In some cases the full stop (period) character is used at the end of a FQDN.
		trimmedHomeDomains = append(trimmedHomeDomains, strings.TrimSuffix(homeDomain, "."))
	}

	jwk := jose.JSONWebKey{}
	err := json.Unmarshal([]byte(opts.JWK), &jwk)
	if err != nil {
		return nil, errors.Wrap(err, "parsing JSON Web Key (JWK)")
	}
	if jwk.Algorithm == "" {
		return nil, errors.New("algorithm (alg) field must be set")
	}

	millenniumTimeout := millenniumclient.MillenniumTimeout
	httpClient := &http.Client{
		Timeout: millenniumTimeout,
	}
	millenniumClient := &millenniumclient.Client{
		MillenniumURL: opts.MillenniumURL,
		HTTP:       httpClient,
	}
	millenniumClient.SetMillenniumTimeout(millenniumTimeout)

	mux := supporthttp.NewAPIMux(opts.Logger)

	mux.NotFound(errorHandler{Error: notFound}.ServeHTTP)
	mux.MethodNotAllowed(errorHandler{Error: methodNotAllowed}.ServeHTTP)

	mux.Get("/health", health.PassHandler{}.ServeHTTP)
	mux.Get("/", challengeHandler{
		Logger:             opts.Logger,
		NetworkPassphrase:  opts.NetworkPassphrase,
		SigningKey:         signingKeys[0],
		ChallengeExpiresIn: opts.ChallengeExpiresIn,
		HomeDomains:        trimmedHomeDomains,
	}.ServeHTTP)
	mux.Post("/", tokenHandler{
		Logger:                      opts.Logger,
		MillenniumClient:               millenniumClient,
		NetworkPassphrase:           opts.NetworkPassphrase,
		SigningAddresses:            signingAddresses,
		JWK:                         jwk,
		JWTIssuer:                   opts.JWTIssuer,
		JWTExpiresIn:                opts.JWTExpiresIn,
		AllowAccountsThatDoNotExist: opts.AllowAccountsThatDoNotExist,
		HomeDomains:                 trimmedHomeDomains,
	}.ServeHTTP)

	return mux, nil
}
