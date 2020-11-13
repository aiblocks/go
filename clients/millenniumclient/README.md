# millenniumclient


`millenniumclient` is a [AiBlocks Go SDK](https://www.aiblocks.io/developers/reference/) package that provides client access to a millennium server. It supports all endpoints exposed by the [millennium API](https://www.aiblocks.io/developers/millennium/reference/index.html).

This project is maintained by the AiBlocks Development Foundation.

## Getting Started
This library is aimed at developers building Go applications that interact with the [AiBlocks network](https://www.aiblocks.io/). It allows users to query the network and submit transactions to the network. The recommended transaction builder for Go programmers is [txnbuild](https://github.com/aiblocks/go/tree/master/txnbuild). Together, these two libraries provide a complete AiBlocks SDK.

* The [millenniumclient API reference](https://godoc.org/github.com/aiblocks/go/clients/millenniumclient).
* The [txnbuild API reference](https://godoc.org/github.com/aiblocks/go/txnbuild).

### Prerequisites
* Go 1.14 or greater
* [Modules](https://github.com/golang/go/wiki/Modules) to manage dependencies

### Installing
* `go get github.com/aiblocks/go/clients/millenniumclient`

### Usage

``` golang
    ...
    import hClient "github.com/aiblocks/go/clients/millenniumclient"
    ...

    // Use the default pubnet client
    client := hClient.DefaultPublicNetClient

    // Create an account request
    accountRequest := hClient.AccountRequest{AccountID: "GCLWGQPMKXQSPF776IU33AH4PZNOOWNAWGGKVTBQMIC5IMKUNP3E6NVU"}

    // Load the account detail from the network
    account, err := client.AccountDetail(accountRequest)
    if err != nil {
        fmt.Println(err)
        return
    }
    // Account contains information about the aiblocks account
    fmt.Print(account)
```
For more examples, refer to the [documentation](https://godoc.org/github.com/aiblocks/go/clients/millenniumclient).

## Running the tests
Run the unit tests from the package directory: `go test`

## Contributing
Please read [Code of Conduct](https://github.com/aiblocks/.github/blob/master/CODE_OF_CONDUCT.md) to understand this project's communication rules.

To submit improvements and fixes to this library, please see [CONTRIBUTING](../CONTRIBUTING.md).

## License
This project is licensed under the Apache License - see the [LICENSE](../../LICENSE-APACHE.txt) file for details.
