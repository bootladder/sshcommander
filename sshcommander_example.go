package main

import(
  "sshcommander"
  "fmt"
  "os/exec"
)
type RealOSCommandExecuter struct {}

func (r RealOSCommandExecuter) Execute(commandstring string) (out string, err error) {
  bytesout,err := exec.Command("sh" , "-c" , commandstring).Output()
  out = string(bytesout)
  return
}

func main() {
  realCommandExecuter := RealOSCommandExecuter{}
  sshcommander.InjectOSCommandExecuter(realCommandExecuter)
  commander := sshcommander.SSHCommander{"root", "209.148.82.100", 22, ""}

  out, _ := commander.Command("cat /etc/issue")

  fmt.Println(out)
}

