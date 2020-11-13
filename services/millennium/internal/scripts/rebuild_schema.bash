#! /usr/bin/env bash
set -e

# This scripts rebuilds the latest.sql file included in the schema package.
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
GOTOP="$( cd "$DIR/../../../../../../../.." && pwd )"

go generate github.com/aiblocks/go/services/millennium/internal/db2/schema
go generate github.com/aiblocks/go/services/millennium/internal/test
go install github.com/aiblocks/go/services/millennium
