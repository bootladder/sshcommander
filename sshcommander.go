package sshcommander
import (
  //"errors"
)
type SSHCommander struct {
  User, Hostname string
  Port int
  Key string
}

func (s *SSHCommander) Command(commandstring string) (err error) {

  return 
}
