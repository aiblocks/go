---
title: Overview
---

Millennium is an API server for the AiBlocks ecosystem.  It acts as the interface between [aiblocks-core](https://github.com/aiblocks/aiblocks-core) and applications that want to access the AiBlocks network. It allows you to submit transactions to the network, check the status of accounts, subscribe to event streams, etc. See [an overview of the AiBlocks ecosystem](https://www.aiblocks.io/developers/guides/) for details of where Millennium fits in.

Millennium provides a RESTful API to allow client applications to interact with the AiBlocks network. You can communicate with Millennium using cURL or just your web browser. However, if you're building a client application, you'll likely want to use a AiBlocks SDK in the language of your client.
SDF provides a [JavaScript SDK](https://www.aiblocks.io/developers/js-aiblocks-sdk/reference/index.html) for clients to use to interact with Millennium.

SDF runs a instance of Millennium that is connected to the test net: [https://millennium-testnet.aiblocks.io/](https://millennium-testnet.aiblocks.io/) and one that is connected to the public AiBlocks network:
[https://millennium.aiblocks.io/](https://millennium.aiblocks.io/).

## Libraries

SDF maintained libraries:<br />
- [JavaScript](https://github.com/aiblocks/js-aiblocks-sdk)
- [Go](https://github.com/aiblocks/go/tree/master/clients/millenniumclient)
- [Java](https://github.com/aiblocks/java-aiblocks-sdk)

Community maintained libraries for interacting with Millennium in other languages:<br>
- [Python](https://github.com/AiBlocksCN/py-aiblocks-base)
- [C# .NET Core 2.x](https://github.com/elucidsoft/dotnetcore-aiblocks-sdk)
- [Ruby](https://github.com/astroband/ruby-aiblocks-sdk)
- [iOS and macOS](https://github.com/Soneso/aiblocks-ios-mac-sdk)
- [Scala SDK](https://github.com/synesso/scala-aiblocks-sdk)
- [C++ SDK](https://github.com/bnogalm/AiBlocksQtSDK)
