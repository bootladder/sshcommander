package sshcommander_test
import (
  "fmt"
  "testing"
  "sshcommander"
  "github.com/stretchr/testify/assert"
)

func TestNewSSHCommander( t *testing.T) {
  commander := sshcommander.SSHCommander{"root", "localhost", 22, ""}
  err := commander.Command("echo hellocommand")
  assert.Equal(t, err, nil)
    fmt.Print("hello")
}
