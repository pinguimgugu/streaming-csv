#!/bin/bash

function mongo {
    docker-compose -f docker-compose-dev.yml up -d mongo.local   
    sleep 5
}

function create_users {
    bash create_users.sh
}

function dep {
    docker-compose -f docker-compose-dev.yml up dep
}

function build {
    docker build -t go-tools .
}

function serve {
    build && dep && mongo && create_users && run
}

function run {
    docker-compose -f docker-compose-dev.yml up streaming_csv
}

"$@"