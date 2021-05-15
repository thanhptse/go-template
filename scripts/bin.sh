#!/usr/bin/env bash

SCRIPTPATH="$(
    cd "$(dirname "$0")"
    pwd -P
)"

CURRENT_DIR=$SCRIPTPATH
ROOT_DIR="$(dirname $CURRENT_DIR)"

function api_start(){
    echo "Starting infrastructure ..."

    set_env_variables

    echo "Start api with config file: $CONFIG_FILE"
    ENTRY_FILE="$ROOT_DIR/cmd/service/main.go"
    go run $ENTRY_FILE --config-file=$CONFIG_FILE
}

function set_env_variables(){
    set -a
    export $(grep -v '^#' "$ROOT_DIR/build/.base.env" | xargs -0) > /dev/null 2>&1
    . $ROOT_DIR/build/.base.env
    set +a
    export CONFIG_FILE=$ROOT_DIR/build/app.yaml
}

case $1 in 
start)
    api_start
    ;;
*)
    echo "./scripts/bin.sh [start|migrate|test]"
    ;;
esac