package main

import(
  "github.com/bootladder/sshcommander/sshcommandcreator"
  "github.com/bootladder/sshcommander/hostconfig"
  "fmt"
  "os"
  "flag"
)

func main() {

  flag.Parse()
  if flag.NArg() == 0 { //need atleast 1 argument after flags
    flag.Usage()
    os.Exit(1)
  }
  var pathtoconfigfile string = "~/.sshcommander/hostconfig.json"
  err := hostconfig.Load(pathtoconfigfile)
  if err != nil {
    fmt.Println("Error loading hostconfig")
    os.Exit(1)
  }

  err = hostconfig.LookupHostname(flag.Arg(0))
  if err != nil {
    fmt.Println("Invalid Hostname")
    os.Exit(1)
  }

  creator := sshcommandcreator.SSHCommandCreator{"root", "111.111.11.111", 22, ""}
  out, _ := creator.CreateCommandString("cat /etc/issue")
  fmt.Println(out)
  os.Exit(0)
}

