#!/usr/bin/env bats

@test "Check that the executable exists" {
    command -v ./sshcommander
}
@test "Executing with no arguments returns exit code 1" {
    run ./sshcommander
    echo "output = ${output}"
    [ $status = 1 ]
}
@test "Executing with invalid host parameter returns exit code 1" {
    run ./sshcommander myinvalidhost
    echo "output = ${output}"
    [ $status = 1 ]
}
@test "Executing with no default hostconfig file returns exit code 1" {
    rm -rf ~/.sshcommander/hostconfig.json 
    run ./sshcommander myinvalidhost
    echo "output = ${output}"
    [ $status = 1 ]
}
@test "Executing with valid default hostconfig returns exit code 1" {
    mkdir -p ~/.sshcommander
    echo blah > ~/.sshcommander/hostconfig.json 
    run ./sshcommander myinvalidhost
    echo "output = ${output}"
    [ $status = 1 ]
}

