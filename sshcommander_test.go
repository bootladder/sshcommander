package sshcommander_test
import (
  "testing"
  "sshcommander"
  "github.com/stretchr/testify/assert"
)

func TestNewSSHCommander( t *testing.T) {
  commander := sshcommander.SSHCommander{"root", "localhost", 22, ""}
  err := commander.Command("echo hellocommand")
  assert.Equal(t, err, nil)
}
//////////////////////////////////////////////////////////////////////
func TestCreateCommandString( t *testing.T) {
  commander := sshcommander.SSHCommander{"root", "localhost", 22, ""}
  actualResult, _ := commander.CreateCommandString("echo hellocommand")
  assert.Equal(t, "ssh -p 22 root@localhost \"echo hellocommand\"", actualResult)
}
func TestCreateCommandStringSetsTargetCommand( t *testing.T) {
  commander := sshcommander.SSHCommander{"root", "localhost", 22, ""}
  actualResult, _ := commander.CreateCommandString("echo a different command")
  assert.Equal(t, "ssh -p 22 root@localhost \"echo a different command\"", actualResult)
}
func TestCreateCommandStringSetsPortHostUser( t *testing.T) {
  commander := sshcommander.SSHCommander{"myuser", "differenthost", 20010, ""}
  actualResult, _ := commander.CreateCommandString("echo hellocommand")
  assert.Equal(t, "ssh -p 20010 myuser@differenthost \"echo hellocommand\"", actualResult)
}
//////////////////////////////////////////////////////////////////////
