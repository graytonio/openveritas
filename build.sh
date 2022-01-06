#!/bin/bash

server_package="veritas-server"
client_package="veritas-client"

target=$1

buildServer() {
    echo "Building Server"
    go build -o $server_package ./server
}

buildCLI() {
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