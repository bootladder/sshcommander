package main
import(
  "sshcommander"
  "fmt"
  "os/exec"
)
type RealOSCommandExecuter struct {}

func (r RealOSCommandExecuter) Execute(commandstring string) {
  fmt.Println(commandstring)
  exec.Command("sh" , "-c" , commandstring).CombinedOutput()
}

func main() {
  realCommandExecuter := RealOSCommandExecuter{}
  sshcommander.InjectOSCommandExecuter(realCommandExecuter)
  commander := sshcommander.SSHCommander{"root", "localhost", 22, ""}
  commander.Command("touch /tmp/touchedbygo")
}

