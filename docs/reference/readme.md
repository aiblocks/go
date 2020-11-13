---
title: Overview
---

The Go SDK is a set of packages for interacting with most aspects of the AiBlocks ecosystem. The primary component is the Millennium SDK, which provides convenient access to Millennium services. There are also packages for other AiBlocks services such as [TOML support](https://github.com/aiblocks/aiblocks-protocol/blob/master/ecosystem/sep-0001.md) and [federation](https://github.com/aiblocks/aiblocks-protocol/blob/master/ecosystem/sep-0002.md).

## Millennium SDK

The Millennium SDK is composed of two complementary libraries: `txnbuild` + `millenniumclient`.
The `txnbuild` ([source](https://github.com/aiblocks/go/tree/master/txnbuild), [docs](https://godoc.org/github.com/aiblocks/go/txnbuild)) package enables the construction, signing and encoding of AiBlocks [transactions](https://www.aiblocks.io/developers/guides/concepts/transactions.html) and [operations](https://www.aiblocks.io/developers/guides/concepts/list-of-operations.html) in Go. The `millenniumclient` ([source](https://github.com/aiblocks/go/tree/master/clients/millenniumclient), [docs](https://godoc.org/github.com/aiblocks/go/clients/millenniumclient)) package provides a web client for interfacing with [Millennium](https://www.aiblocks.io/developers/guides/get-started/) server REST endpoints to retrieve ledger information, and to submit transactions built with `txnbuild`.

## List of major SDK packages

- `millenniumclient` ([source](https://github.com/aiblocks/go/tree/master/clients/millenniumclient), [docs](https://godoc.org/github.com/aiblocks/go/clients/millenniumclient)) - programmatic client access to Millennium
- `txnbuild` ([source](https://github.com/aiblocks/go/tree/master/txnbuild), [docs](https://godoc.org/github.com/aiblocks/go/txnbuild)) - construction, signing and encoding of AiBlocks transactions and operations
- `aiblockstoml` ([source](https://github.com/aiblocks/go/tree/master/clients/aiblockstoml), [docs](https://godoc.org/github.com/aiblocks/go/clients/aiblockstoml)) - parse [AiBlocks.toml](../../guides/concepts/aiblocks-toml.md) files from the internet
- `federation` ([source](https://godoc.org/github.com/aiblocks/go/clients/federation)) - resolve federation addresses  into aiblocks account IDs, suitable for use within a transaction

