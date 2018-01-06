package sshcommander
import (
  "fmt"
  "strings"
)
type SSHCommander struct {
  User, Hostname string
  Port int
  Key string
}

type OSCommandExecuter interface {

}


func (s *SSHCommander) Command(commandstring string) (err error) {
  return
}

func (s *SSHCommander) CreateCommandString(commandstring string) (out string, err error) {

  cmd := []string{
    fmt.Sprintf("ssh -p %d %s@%s", s.Port, s.User, s.Hostname),
  }
  cmd = append(cmd, " ", "\"",commandstring, "\"")

  return  strings.Join(cmd,""),nil
}

func InjectOSCommandExecuter( z OSCommandExecuter) {

}
