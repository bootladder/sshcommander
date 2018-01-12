package sshcommandcreator

import (
  "fmt"
  "strings"
)

type SSHCommandCreator struct {
  User, Hostname string
  Port string
  Key string
}

func (s *SSHCommandCreator) CreateCommandString(commandstring string) (out string, err error) {

  cmd := []string{
    fmt.Sprintf("ssh -p %s %s@%s", s.Port, s.User, s.Hostname),
  }
  cmd = append(cmd, " ", commandstring)

  return  strings.Join(cmd,""),nil
}
