# txnbuild

`txnbuild` is a [AiBlocks SDK](https://www.aiblocks.io/developers/reference/), implemented in [Go](https://golang.org/). It provides a reference implementation of the complete [set of operations](https://www.aiblocks.io/developers/guides/concepts/list-of-operations.html) that compose [transactions](https://www.aiblocks.io/developers/guides/concepts/transactions.html) for the AiBlocks distributed ledger.

This project is maintained by the AiBlocks Development Foundation.

```golang
    import (
        "log"
        
        "github.com/aiblocks/go/clients/millenniumclient"
        "github.com/aiblocks/go/keypair"
        "github.com/aiblocks/go/network"
        "github.com/aiblocks/go/txnbuild"
    )
    
    // Make a keypair for a known account from a secret seed
    kp, _ := keypair.Parse("SBPQUZ6G4FZNWFHKUWC5BEYWF6R52E3SEP7R3GWYSM2XTKGF5LNTWW4R")
    
    // Get the current state of the account from the network
    client := millenniumclient.DefaultTestNetClient
    ar := millenniumclient.AccountRequest{AccountID: kp.Address()}
    sourceAccount, err := client.AccountDetail(ar)
    if err != nil {
        log.Fatalln(err)
    }
    
    // Build an operation to create and fund a new account
    op := txnbuild.CreateAccount{
        Destination: "GCCOBXW2XQNUSL467IEILE6MMCNRR66SSVL4YQADUNYYNUVREF3FIV2Z",
        Amount:      "10",
    }
    
    // Construct the transaction that holds the operations to execute on the network
    tx, err := txnbuild.NewTransaction(
        txnbuild.TransactionParams{
            SourceAccount:        &sourceAccount,
            IncrementSequenceNum: true,
            Operations:           []txnbuild.Operation{&op},
            BaseFee:              txnbuild.MinBaseFee,
            Timebounds:           txnbuild.NewTimeout(300),
        },
    )
    if err != nil {
        log.Fatalln(err)
    )
    
    // Sign the transaction
    tx, err = tx.Sign(network.TestNetworkPassphrase, kp.(*keypair.Full))
    if err != nil {
        log.Fatalln(err)
    )
    
    // Get the base 64 encoded transaction envelope
    txe, err := tx.Base64()
    if err != nil {
        log.Fatalln(err)
    }
    
    // Send the transaction to the network
    resp, err := client.SubmitTransactionXDR(txe)
    if err != nil {
        log.Fatalln(err)
    }
```

## Getting Started
This library is aimed at developers building Go applications on top of the [AiBlocks network](https://www.aiblocks.io/). Transactions constructed by this library may be submitted to any Millennium instance for processing onto the ledger, using any AiBlocks SDK client. The recommended client for Go programmers is [millenniumclient](https://github.com/aiblocks/go/tree/master/clients/millenniumclient). Together, these two libraries provide a complete AiBlocks SDK.

* The [txnbuild API reference](https://godoc.org/github.com/aiblocks/go/txnbuild).
* The [millenniumclient API reference](https://godoc.org/github.com/aiblocks/go/clients/millenniumclient).

An easy-to-follow demonstration that exercises this SDK on the TestNet with actual accounts is also included! See the [Demo](#demo) section below.

### Prerequisites
* Go 1.14 or greater
* [Modules](https://github.com/golang/go/wiki/Modules) to manage dependencies

### Installing
* `go get github.com/aiblocks/go/clients/txnbuild`

## Running the tests
Run the unit tests from the package directory: `go test`

## Demo
To see the SDK in action, build and run the demo:
* Enter the demo directory: `cd $GOPATH/src/github.com/aiblocks/go/txnbuild/cmd/demo`
* Build the demo: `go build`
* Run the demo: `./demo init`


## Contributing
Please read [Code of Conduct](https://github.com/aiblocks/.github/blob/master/CODE_OF_CONDUCT.md) to understand this project's communication rules.

To submit improvements and fixes to this library, please see [CONTRIBUTING](../CONTRIBUTING.md).

## License
This project is licensed under the Apache License - see the [LICENSE](../../LICENSE-APACHE.txt) file for details.
