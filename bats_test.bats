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


@test "No default hostconfig returns error msg" {
    rm -rf ~/.sshcommander
    ./sshcommander myinvalidhost echo hello | grep -i "error loading hostconfig"
}
@test "Invalid default hostconfig returns error msg" {
    mkdir -p ~/.sshcommander
    echo blah > ~/.sshcommander/hostconfig.json
    ./sshcommander myinvalidhost echo hello | grep -i "invalid"
}
@test "Valid default hostconfig invalid hostname returns error msg" {
    mkdir -p ~/.sshcommander
    cp hostconfig/samplehostconfig.json ~/.sshcommander/hostconfig.json
    ./sshcommander myinvalidhost echo hello | grep -i "invalid hostname"
}

##########################################################
#
@test "Test feature only prints the command, does not execute" {
    mkdir -p ~/.sshcommander
    cp hostconfig/samplehostconfig.json ~/.sshcommander/hostconfig.json
    ./sshcommander -N myfirsthostname echo hello | grep -i "ssh"
}

@test "Multiple Arguments in commandline show up" {
    mkdir -p ~/.sshcommander
    cp hostconfig/samplehostconfig.json ~/.sshcommander/hostconfig.json
    ./sshcommander -N myfirsthostname echo hello | grep -i "echo hello"
}

@test "Nested Host correct commandline" {
    mkdir -p ~/.sshcommander
    cp hostconfig/samplehostconfig.json ~/.sshcommander/hostconfig.json
    ./sshcommander -N myfirstnestedhostname echo hello | grep -i "myfirsthostname"
}
