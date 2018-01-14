package sshcommandcreator

import (
  "fmt"
  "strings"
)

type SSHCommandCreator struct {
  User, Hostname string
  Port string
  Key string
  AdditionalArgs string
}

func (s *SSHCommandCreator) Create(user,hostname, port, key string) {
  s.Port = port
  s.User = user
  s.Hostname = hostname
  s.Key = key
  s.AdditionalArgs = ""
}

func (s *SSHCommandCreator) CreateCommandString(commandstring string) (out string, err error) {

  cmd := []string{
    "ssh",
  }
  if s.AdditionalArgs != "" {
    cmd = append(cmd, s.AdditionalArgs)
  }
  if s.Key != "" {
    cmd = append(cmd, "-i",s.Key)
  }
    //fmt.Sprintf("%s", s.AdditionalArgs),
  lastpart :=  fmt.Sprintf("-p %s %s@%s", s.Port, s.User, s.Hostname)
  cmd = append(cmd, lastpart, commandstring)

  return  strings.Join(cmd," "),nil
}
