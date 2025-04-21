#!/bin/bash

# Tweak PATH for Travis
export PATH=$PATH:$HOME/gopath/bin

OPTIONS="-config=config/dbconfig.yml -env postgres"

set -ex

pwd
sql-migrate status $OPTIONS
sql-migrate up $OPTIONS
sql-migrate down $OPTIONS
sql-migrate redo $OPTIONS
sql-migrate status $OPTIONS