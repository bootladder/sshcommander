package main

import(
  "github.com/bootladder/sshcommander/sshcommandcreator"
  "fmt"
)

func main() {
  creator := sshcommandcreator.SSHCommandCreator{"root", "111.111.11.111", 22, ""}
  out, _ := creator.CreateCommandString("cat /etc/issue")
  fmt.Println(out)
}

