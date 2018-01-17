#!/bin/bash
echo Building the main executable
go get -t -u ./...
go build -o sshcommander || exit 1
echo Running Unit Tests
go test ./... || exit 1
echo Running BATS bash tests

echo preserve the existing hostconfig.json
cp ~/.sshcommander/hostconfig.json /tmp/oldhostconfig.json

bats bats_test.bats

echo restore the existing hostconfig.json
cp /tmp/oldhostconfig.json ~/.sshcommander/hostconfig.json

go install
