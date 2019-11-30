#!/bin/bash -ex

export PATH="/root/go/bin:$PATH"

cd /protobufs
rm -rf generated
mkdir -p generated/go generated/nanopb

cd proto

# Generate golang code
find . -name "*.proto" | xargs -L1 protoc --go_out=../generated/go
# Generate C code (nanopb)
find . -name "*.proto" | xargs -L1 protoc --plugin=protoc-gen-nanopb=/nanopb/generator/protoc-gen-nanopb --nanopb_out=../generated/nanopb
