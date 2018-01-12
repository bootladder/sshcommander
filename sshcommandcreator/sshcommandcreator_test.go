package sshcommandcreator_test

import (
  "github.com/bootladder/sshcommander/sshcommandcreator"
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestNewSSHCommander( t *testing.T) {
  creator := sshcommandcreator.SSHCommandCreator{"root", "localhost", "22", ""}
  assert.Equal(t, creator, creator)
}
//////////////////////////////////////////////////////////////////////
func TestCreateCommandString( t *testing.T) {
  creator := sshcommandcreator.SSHCommandCreator{"root", "localhost", "22", ""}
  actualResult, _ := creator.CreateCommandString("echo hellocommand")
  assert.Equal(t, "ssh -p 22 root@localhost echo hellocommand", actualResult)
}
func TestCreateCommandStringSetsTargetCommand( t *testing.T) {
  creator := sshcommandcreator.SSHCommandCreator{"root", "localhost", "22", ""}
  actualResult, _ := creator.CreateCommandString("echo a different command")
  assert.Equal(t, "ssh -p 22 root@localhost echo a different command", actualResult)
}
func TestCreateCommandStringSetsPortHostUser( t *testing.T) {
  creator := sshcommandcreator.SSHCommandCreator{"myuser", "differenthost", "20010", ""}
  actualResult, _ := creator.CreateCommandString("echo hellocommand")
  assert.Equal(t, "ssh -p 20010 myuser@differenthost echo hellocommand", actualResult)
}
//////////////////////////////////////////////////////////////////////
//Nesting Tests
func TestCreateCommandStringNestedWorks( t *testing.T) {
  creator1 := sshcommandcreator.SSHCommandCreator{"myuser", "differenthost", "20010", ""}
  creator2 := sshcommandcreator.SSHCommandCreator{"myother", "otherhost", "10099", ""}
  actualResult1, _ := creator1.CreateCommandString("echo hellocommand")
  actualResult2, _ := creator2.CreateCommandString(actualResult1)
  assert.Equal(t, "ssh -p 10099 myother@otherhost ssh -p 20010 myuser@differenthost echo hellocommand", actualResult2)
}
