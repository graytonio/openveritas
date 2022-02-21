#!/bin/bash

parent_path=$(cd "$(dirname "${BASH_SOURCE[0]}")"; pwd -P)

server_package="$parent_path/dev-build/veritas-server"
client_package="$parent_path/dev-build/veritas-client"

target=$1

buildServer() {
    cd "$parent_path"
    mkdir -p ./dev-build
    echo "Building Server"
    go build -o $server_package ./server
}

buildCLI() {
    cd "$parent_path"
    mkdir -p ./dev-build
    echo "Building CLI"
    go build -o $client_package ./cli
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