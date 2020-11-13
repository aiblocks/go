---
title: Millennium Development Guide
---
## Millennium Development Guide

This document describes how to build Millennium from source, so that you can test and edit the code locally to develop bug fixes and new features.

If you are just starting with Millennium and want to try it out, consider the [Quickstart Guide](quickstart.md) instead. For information about administrating a Millennium instance in production, check out the [Administration Guide](admin.md).

## Building Millennium
Building Millennium requires the following developer tools:

- A [Unix-like](https://en.wikipedia.org/wiki/Unix-like) operating system with the common core commands (cp, tar, mkdir, bash, etc.)
- Golang 1.14 or later
- [git](https://git-scm.com/) (to check out Millennium's source code)
- [mercurial](https://www.mercurial-scm.org/) (needed for `go-dep`)

1. Set your [GOPATH](https://github.com/golang/go/wiki/GOPATH) environment variable, if you haven't already. The default `GOPATH` is `$HOME/go`. When building any Go package or application the binaries will be installed by default to `$GOPATH/bin`.
2. Checkout the code into any directory you prefer:
   ```
   git checkout https://github.com/aiblocks/go
   ```
   Or if you prefer to develop inside `GOPATH` check it out to `$GOPATH/src/github.com/aiblocks/go`:
   ```
   git checkout https://github.com/aiblocks/go $GOPATH/src/github.com/aiblocks/go
   ```
   If developing inside `GOPATH` set the `GO111MODULE=on` environment variable to turn on Modules for managing dependencies. See the repository [README](../../../../README.md#dependencies) for more information.
3. Change to the directory where the repository is checked out. e.g. `cd go`, or if developing inside the `GOPATH`, `cd $GOPATH/src/github.com/aiblocks/go`.
4. Compile the Millennium binary: `go install ./services/millennium`. You should see the resulting `millennium` executable in `$GOPATH/bin`.
5. Add Go binaries to your PATH in your `bashrc` or equivalent, for easy access: `export PATH=${GOPATH//://bin:}/bin:$PATH`

Open a new terminal. Confirm everything worked by running `millennium --help` successfully. You should see an informative message listing the command line options supported by Millennium.

## Set up Millennium's database
Millennium uses a Postgres database backend to store test fixtures and record information ingested from an associated AiBlocks Core. To set this up:
1. Install [PostgreSQL](https://www.postgresql.org/).
2. Run `createdb millennium_dev` to initialise an empty database for Millennium's use.
3. Run `millennium db init --db-url postgres://localhost/millennium_dev` to install Millennium's database schema.

### Database problems?
1. Depending on your installation's defaults, you may need to configure a Postgres DB user with appropriate permissions for Millennium to access the database you created. Refer to the [Postgres documentation](https://www.postgresql.org/docs/current/sql-createuser.html) for details. Note: Remember to restart the Postgres server after making any changes to `pg_hba.conf` (the Postgres configuration file), or your changes won't take effect!
2. Make sure you pass the appropriate database name and user (and port, if using something non-standard) to Millennium using `--db-url`. One way is to use a Postgres URI with the following form: `postgres://USERNAME:PASSWORD@localhost:PORT/DB_NAME`.
3. If you get the error `connect failed: pq: SSL is not enabled on the server`, add `?sslmode=disable` to the end of the Postgres URI to allow connecting without SSL.
4. If your server is responding strangely, and you've exhausted all other options, reboot the machine. On some systems `service postgresql restart` or equivalent may not fully reset the state of the server.

## Run tests
At this point you should be able to run Millennium's unit tests:
```bash
cd $GOPATH/src/github.com/aiblocks/go/services/millennium
go test ./...
```

## Set up AiBlocks Core
Millennium provides an API to the AiBlocks network. It does this by ingesting data from an associated `aiblocks-core` instance. Thus, to run a full Millennium instance requires a `aiblocks-core` instance to be configured, up to date with the network state, and accessible to Millennium. Millennium accesses `aiblocks-core` through both an HTTP endpoint and by connecting directly to the `aiblocks-core` Postgres database.

The simplest way to set up AiBlocks Core is using the [AiBlocks Quickstart Docker Image](https://github.com/aiblocks/docker-aiblocks-core-millennium). This is a Docker container that provides both `aiblocks-core` and `millennium`, pre-configured for testing.

1. Install [Docker](https://www.docker.com/get-started).
2. Verify your Docker installation works: `docker run hello-world`
3. Create a local directory that the container can use to record state. This is helpful because it can take a few minutes to sync a new `aiblocks-core` with enough data for testing, and because it allows you to inspect and modify the configuration if needed. Here, we create a directory called `aiblocks` to use as the persistent volume: `cd $HOME; mkdir aiblocks`
4. Download and run the AiBlocks Quickstart container:

```bash
docker run --rm -it -p "8000:8000" -p "11626:11626" -p "11625:11625" -p"8002:5432" -v $HOME/aiblocks:/opt/aiblocks --name aiblocks aiblocks/quickstart --testnet
```

In this example we run the container in interactive mode. We map the container's Millennium HTTP port (`8000`), the `aiblocks-core` HTTP port (`11626`), and the `aiblocks-core` peer node port (`11625`) from the container to the corresponding ports on `localhost`. Importantly, we map the container's `postgresql` port (`5432`) to a custom port (`8002`) on `localhost`, so that it doesn't clash with our local Postgres install.
The `-v` option mounts the `aiblocks` directory for use by the container. See the [Quickstart Image documentation](https://github.com/aiblocks/docker-aiblocks-core-millennium) for a detailed explanation of these options.

5. The container is running both a `aiblocks-core` and a `millennium` instance. Log in to the container and stop Millennium:
```bash
docker exec -it aiblocks /bin/bash
supervisorctl
stop millennium
```

## Check AiBlocks Core status
AiBlocks Core takes some time to synchronise with the rest of the network. The default configuration will pull roughly a couple of day's worth of ledgers, and may take 15 - 30 minutes to catch up. Logs are stored in the container at `/var/log/supervisor`. You can check the progress by monitoring logs with `supervisorctl`:
```bash
docker exec -it aiblocks /bin/bash
supervisorctl tail -f aiblocks-core
```

You can also check status by looking at the HTTP endpoint, e.g. by visiting http://localhost:11626 in your browser.

## Connect Millennium to AiBlocks Core
You can connect Millennium to `aiblocks-core` at any time, but Millennium will not begin ingesting data until `aiblocks-core` has completed its catch-up process.

Now run your development version of Millennium (which is outside of the container), pointing it at the `aiblocks-core` running inside the container:

```bash
millennium --db-url="postgres://localhost/millennium_dev" --aiblocks-core-db-url="postgres://aiblocks:postgres@localhost:8002/core" --aiblocks-core-url="http://localhost:11626" --port 8001 --network-passphrase "Test SDF Network ; September 2015" --ingest
```

If all is well, you should see ingest logs written to standard out. You can test your Millennium instance with a query like: http://localhost:8001/transactions?limit=10&order=asc. Use the [AiBlocks Laboratory](https://www.aiblocks.io/laboratory/) to craft other queries to try out,
and read about the available endpoints and see examples in the [Millennium API reference](https://www.aiblocks.io/developers/millennium/reference/).

## The development cycle
Congratulations! You can now run the full development cycle to build and test your code.
1. Write code + tests
2. Run tests
3. Compile Millennium: `go install github.com/aiblocks/go/services/millennium`
4. Run Millennium (pointing at your running `aiblocks-core`)
5. Try Millennium queries

Check out the [AiBlocks Contributing Guide](https://github.com/aiblocks/docs/blob/master/CONTRIBUTING.md) to see how to contribute your work to the AiBlocks repositories. Once you've got something that works, open a pull request, linking to the issue that you are resolving with your contribution. We'll get back to you as quickly as we can.
