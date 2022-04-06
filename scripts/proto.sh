#!/bin/bash

# Generate all proto code

protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    notes/notes.proto