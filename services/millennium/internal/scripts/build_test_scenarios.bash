#! /usr/bin/env bash
set -e

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
GOTOP="$( cd "$DIR/../../../../../../../.." && pwd )"
PACKAGES=$(find $GOTOP/src/github.com/aiblocks/go/services/millennium/internal/test/scenarios -iname '*.rb' -not -name '_common_accounts.rb')
#PACKAGES=$(find $GOTOP/src/github.com/aiblocks/go/services/millennium/internal/test/scenarios -iname 'failed_transactions.rb')

go install github.com/aiblocks/go/services/millennium

dropdb hayashi_scenarios --if-exists
createdb hayashi_scenarios

export AIBLOCKS_CORE_DATABASE_URL="postgres://localhost/hayashi_scenarios?sslmode=disable"
export DATABASE_URL="postgres://localhost/millennium_scenarios?sslmode=disable"
export NETWORK_PASSPHRASE="Test SDF Network ; September 2015"
export AIBLOCKS_CORE_URL="http://localhost:8080"
export SKIP_CURSOR_UPDATE="true"
export INGEST_FAILED_TRANSACTIONS=true

# run all scenarios
for i in $PACKAGES; do
  echo $i
  CORE_SQL="${i%.rb}-core.sql"
  MILLENNIUM_SQL="${i%.rb}-millennium.sql"
  scc -r $i --allow-failed-transactions --dump-root-db > $CORE_SQL

  # load the core scenario
  psql $AIBLOCKS_CORE_DATABASE_URL < $CORE_SQL

  # recreate millennium dbs
  dropdb millennium_scenarios --if-exists
  createdb millennium_scenarios

  # import the core data into millennium
  $GOTOP/bin/millennium db init
  $GOTOP/bin/millennium db init-asset-stats
  $GOTOP/bin/millennium db rebase

  # write millennium data to sql file
  pg_dump $DATABASE_URL \
    --clean --if-exists --no-owner --no-acl --inserts \
    | sed '/SET idle_in_transaction_session_timeout/d' \
    | sed '/SET row_security/d' \
    > $MILLENNIUM_SQL
done


# commit new sql files to bindata
go generate github.com/aiblocks/go/services/millennium/internal/test/scenarios
# go test github.com/aiblocks/go/services/millennium/internal/ingest
