#!/bin/bash

dlv_pid=""

function cleanup ()
{
  echo "kill dlv"
  kill "${dlv_pid}"
}

trap "cleanup" EXIT INT

# build code
go mod vendor
go build -o api

# start debugger and open port for remote
dlv exec ./api --headless --listen=:2345 --continue --api-version=2 --accept-multiclient &

dlv_pid=$!

# Improvements: Just montior .go files?
inotifywait -e create -e modify -e delete -e move --exclude .git -r /app
