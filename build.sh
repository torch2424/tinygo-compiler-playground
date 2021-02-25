#!/bin/bash

# Install the latest tinygo
(cd ../tinygo && go build && go install)

# Build the module ll and wasm
tinygo build -o ./main.ll -target wasi ./main.go
tinygo build -o ./main.wasm -target wasi ./main.go


