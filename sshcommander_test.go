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
// Check that we can inject an OSCommandExecuter
type FakeOSCommandExecuter struct {}
var globalCheck bool = false
var globalCommandString string

func (f FakeOSCommandExecuter) Execute(commandstring string) {
  globalCheck = true
  globalCommandString  = commandstring
}

func TestInjectOSCommandExecuter( t *testing.T) {
  faker := FakeOSCommandExecuter{}
  sshcommander.InjectOSCommandExecuter(faker)
  commander := sshcommander.SSHCommander{"myuser", "differenthost", 20010, ""}
  commander.Command("echo newcommand")
}

func TestCommand_CallsInjectedFakeExecuter( t *testing.T) {
  faker := FakeOSCommandExecuter{}
  sshcommander.InjectOSCommandExecuter(faker)
  commander := sshcommander.SSHCommander{"myuser", "differenthost", 20010, ""}
  commander.Command("cat /etc/issue")
  assert.Equal(t, true, globalCheck)
  assert.Equal(t, "ssh -p 20010 myuser@differenthost \"cat /etc/issue\"", globalCommandString)
}
