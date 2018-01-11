#!/bin/bash
echo Building the main executable
go build -o sshcommander
echo Running Unit Tests
go test ./...
echo Running BATS bash tests
bats bats_test.bats
