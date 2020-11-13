# captivecore

The Captive AiBlocks-Core Server allows you to run a dedicated AiBlocks-Core instance
for the purpose of ingestion. The server must be bundled with a AiBlocks Core binary.

If you run Millennium with Captive AiBlocks-Core ingestion enabled Millennium will spawn a AiBlocks-Core
subprocess. Millennium's ingestion system will then stream ledgers from the subprocess via
a filesystem pipe. The disadvantage of running both Millennium and the AiBlocks-Core subprocess
on the same machine is it requires detailed per-process monitoring to be able to attribute
potential issues (like memory leaks) to a specific service.

Now you can run Millennium and pair it with a remote Captive AiBlocks-Core instance. The
Captive AiBlocks-Core Server can run on a separate machine from Millennium. The server
will manage AiBlocks-Core as a subprocess and provide an HTTP API which Millennium
can use remotely to stream ledgers for the purpose of ingestion.

Note that, currently, a single Captive AiBlocks-Core Server cannot be shared by
multiple Millennium instances.

## API

### `GET /latest-sequence`

Fetches the latest ledger sequence available on the captive core instance.

Response:

```json
{
	"sequence": 12345
}
```


### `GET /ledger/<sequence>`

Fetches the ledger with the given sequence number from the captive core instance.

Response:


```json
{
    "present": true,
    "ledger": "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=="
}
```

### `POST /prepare-range`

Preloads the given range of ledgers in the captive core instance.

Bounded request:
```json
{
    "from": 123,
    "to":   150,
    "bounded": true
}
```

Unbounded request:
```json
{
    "from": 123,
    "bounded": false
}
```

Response:
```json
{
    "ledgerRange": {"from":  123, "bounded":  false},
    "startTime": "2020-08-31T13:29:09Z",
    "ready": true,
    "readyDuration": 1000
}
```

## Usage

```
$ captivecore --help
Run the Captive AiBlocks-Core Server

Usage:
  captivecore [flags]

Flags:
      --aiblocks-core-binary-path           Path to aiblocks core binary
      --aiblocks-core-config-path           Path to aiblocks core config file
      --history-archive-urls               Comma-separated list of aiblocks history archives to connect with
      --log-level                          Minimum log severity (debug, info, warn, error) to log (default info)
      --network-passphrase string          Network passphrase of the AiBlocks network transactions should be signed for (NETWORK_PASSPHRASE) (default "Test SDF Network ; September 2015")
      --port int                           Port to listen and serve on (PORT) (default 8000)
```