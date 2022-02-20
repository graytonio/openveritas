#!/bin/bash

server_package="veritas-server"
client_package="veritas-client"
parent_path=$(cd "$(dirname "${BASH_SOURCE[0]}")"; pwd -P)

target=$1

buildServer() {
    cd "$parent_path"
    mkdir -p ./dev-build
    echo "Building Server"
    go build -o $server_package ./dev-build/server
}

buildCLI() {
    cd "$parent_path"
    mkdir -p ./dev-build
    echo "Building CLI"
    go build -o $client_package ./dev-build/cli
}

if [[ -z $target ]]; then
    buildServer
    buildCLI
elif [[ $target == "clean" ]]; then
    rm $server_package
    rm $client_package
elif [[ $target == "server" ]]; then
    buildServer
elif [[ $target == "client" ]]; then
    buildCLI
fi