#!/bin/bash

OPTIONS="-config=config/dbconfig.yml -env postgres"

# Collect all command line arguments into an array
ARGS=("$@")

# Extract the first argument for the command
COMMAND="${ARGS[0]}"

# Remove the first argument from the array
ARGS=("${ARGS[@]:1}")

# Keep only args that don't start with "-" prefix
FILTERED_ARGS=()
for arg in "${ARGS[@]}"; do
  if [[ "$arg" != -* ]]; then
    FILTERED_ARGS+=("$arg")
  fi
done

# Join remaining arguments as additional options
# Filter only args that start with "-" prefix
ADDITIONAL_OPTIONS=""
for arg in "${ARGS[@]}"; do
  if [[ "$arg" == -* ]]; then
    ADDITIONAL_OPTIONS+=" $arg"
  fi
done
ARGS=("${FILTERED_ARGS[@]}")

set -ex

sql-migrate $COMMAND $OPTIONS $ADDITIONAL_OPTIONS $ARGS