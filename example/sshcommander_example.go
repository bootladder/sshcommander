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
  commander := sshcommander.SSHCommander{"steve", "localhost", 22, ""}

  out, _ := commander.Command("echo I can echo stdout")

  fmt.Println(out)
}

