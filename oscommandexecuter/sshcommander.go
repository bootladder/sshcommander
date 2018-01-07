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

var myOSCommandExecuter OSCommandExecuter

type OSCommandExecuter interface {
  Execute(commandstring string) (string,error)
}


func (s *SSHCommander) Command(commandstring string) (out string,err error) {
  if myOSCommandExecuter == nil {
    fmt.Print("nil myOSCommandExecuter ! Must inject")
  }  else{
    fullSSHcommandline, _ := s.CreateCommandString(commandstring)
    out,err = myOSCommandExecuter.Execute(fullSSHcommandline)
  }
  return
}

func (s *SSHCommander) CreateCommandString(commandstring string) (out string, err error) {

  cmd := []string{
    fmt.Sprintf("ssh -p %d %s@%s", s.Port, s.User, s.Hostname),
  }
  cmd = append(cmd, " ", commandstring)

  return  strings.Join(cmd,""),nil
}

func InjectOSCommandExecuter( z OSCommandExecuter) {
  myOSCommandExecuter = z
}
