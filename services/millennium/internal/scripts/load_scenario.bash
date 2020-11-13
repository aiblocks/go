#! /usr/bin/env bash

set -e

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
SCENARIO=$1
CORE_SQL=$DIR/../test/scenarios/$SCENARIO-core.sql
MILLENNIUM_SQL=$DIR/../test/scenarios/$SCENARIO-millennium.sql

echo "psql $AIBLOCKS_CORE_DATABASE_URL < $CORE_SQL" 
psql $AIBLOCKS_CORE_DATABASE_URL < $CORE_SQL 
echo "psql $DATABASE_URL < $MILLENNIUM_SQL"
psql $DATABASE_URL < $MILLENNIUM_SQL 
