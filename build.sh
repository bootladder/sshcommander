#!/bin/bash
echo Building the main executable
go build -o sshcommander
echo Running Unit Tests
go test ./...
echo Running BATS bash tests

touch /tmp/executedhere
echo preserve the existing hostconfig.json
cp ~/.sshcommander/hostconfig.json /tmp/oldhostconfig.json
bats bats_test.bats

echo restore the existing hostconfig.json
cp /tmp/oldhostconfig.json ~/.sshcommander/hostconfig.json
