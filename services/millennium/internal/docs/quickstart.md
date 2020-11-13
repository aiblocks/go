---
title: Millennium Quickstart
replacement: https://developers.aiblocks.io/docs/run-api-server/quickstart/
---
## Millennium Quickstart
This document describes how to quickly set up a **test** AiBlocks Core + Millennium node, that you can play around with to get a feel for how a aiblocks node operates. **This configuration is not secure!** It is **not** intended as a guide for production administration.

For detailed information about running Millennium and AiBlocks Core safely in production see the [Millennium Administration Guide](admin.md) and the [AiBlocks Core Administration Guide](https://www.aiblocks.io/developers/aiblocks-core/software/admin.html).

If you're ready to roll up your sleeves and dig into the code, check out the [Developer Guide](developing.md).

### Install and run the Quickstart Docker Image
The fastest way to get up and running is using the [AiBlocks Quickstart Docker Image](https://github.com/aiblocks/docker-aiblocks-core-millennium). This is a Docker container that provides both `aiblocks-core` and `millennium`, pre-configured for testing.

1. Install [Docker](https://www.docker.com/get-started).
2. Verify your Docker installation works: `docker run hello-world`
3. Create a local directory that the container can use to record state. This is helpful because it can take a few minutes to sync a new `aiblocks-core` with enough data for testing, and because it allows you to inspect and modify the configuration if needed. Here, we create a directory called `aiblocks` to use as the persistent volume:
`cd $HOME; mkdir aiblocks`
4. Download and run the AiBlocks Quickstart container, replacing `USER` with your username:

```bash
docker run --rm -it -p "8000:8000" -p "11626:11626" -p "11625:11625" -p"8002:5432" -v $HOME/aiblocks:/opt/aiblocks --name aiblocks aiblocks/quickstart --testnet
```

You can check out AiBlocks Core status by browsing to http://localhost:11626.

You can check out your Millennium instance by browsing to http://localhost:8000.

You can tail logs within the container to see what's going on behind the scenes:
```bash
docker exec -it aiblocks /bin/bash
supervisorctl tail -f aiblocks-core
supervisorctl tail -f millennium stderr
```

On a modern laptop this test setup takes about 15 minutes to synchronise with the last couple of days of testnet ledgers. At that point Millennium will be available for querying. 

See the [Quickstart Docker Image](https://github.com/aiblocks/docker-aiblocks-core-millennium) documentation for more details, and alternative ways to run the container. 

You can test your Millennium instance with a query like: http://localhost:8000/transactions?cursor=&limit=10&order=asc. Use the [AiBlocks Laboratory](https://www.aiblocks.io/laboratory/) to craft other queries to try out,
and read about the available endpoints and see examples in the [Millennium API reference](https://www.aiblocks.io/developers/millennium/reference/).

