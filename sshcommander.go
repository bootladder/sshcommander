package main

import(
  "github.com/bootladder/sshcommander/sshcommandcreator"
  "github.com/bootladder/sshcommander/hostconfig"
  "fmt"
  "os"
  "os/user"
  "flag"
)

func GetPathToConfigFile() (path string) {

  usr, _ := user.Current()
  dir := usr.HomeDir
  return dir + "/.sshcommander/hostconfig.json"
}

func main() {

  flag.Parse()
  if flag.NArg() < 2 {
    fmt.Println("need atleast 2 arguments after flags, hostname and command")
    flag.Usage()
    os.Exit(1)
  }
  pathtoconfigfile := GetPathToConfigFile()
  err := hostconfig.Load(pathtoconfigfile)
  if err != nil {
    fmt.Printf("hostconfig.Load : %s\n",err)
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
