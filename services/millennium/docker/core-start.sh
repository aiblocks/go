#!/usr/bin/env bash

set -e

source /etc/profile

echo "using config:"
cat aiblocks-core.cfg

# initialize new db
aiblocks-core new-db

if [ "$1" = "standalone" ]; then
  # start a network from scratch
  aiblocks-core force-scp

  # initialze history archive for stand alone network
  aiblocks-core new-hist vs

  # serve history archives to millennium on port 1570
  pushd /history/vs/
  python3 -m http.server 1570 &
  popd
fi

exec /init -- aiblocks-core run