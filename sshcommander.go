package main

import(
  "github.com/bootladder/sshcommander/sshcommandcreator"
  "github.com/bootladder/sshcommander/hostconfig"
  "fmt"
  "os"
  "os/user"
  "os/exec"
  "flag"
  "strings"
)

func GetPathToConfigFile() (path string) {

  usr, _ := user.Current()
  dir := usr.HomeDir
  return dir + "/.sshcommander/hostconfig.json"
}

func CreateCommandLine( thishost string, joinedargs string ) (out string){

  port := hostconfig.HostGetPort(thishost )
  user := hostconfig.HostGetUser(thishost )
  key := hostconfig.HostGetKey(thishost )
  hostname := hostconfig.HostGetHostname(thishost )

  //creator := sshcommandcreator.SSHCommandCreator{"root", "111.111.11.111", 22, ""}
  creator := sshcommandcreator.SSHCommandCreator{}
  creator.Port = port
  creator.User = user
  creator.Hostname = hostname
  creator.Key = key

  out, _ = creator.CreateCommandString( joinedargs )
  fmt.Println(out)
  return
}

func main() {
  donotexecute := flag.Bool("N", false, "Don't execute, just print the command")
  flag.Parse()
  if flag.NArg() < 2 {
    fmt.Println("need atleast 2 arguments after flags, hostname and command")
    flag.Usage()
    os.Exit(1)
  }
  pathtoconfigfile := GetPathToConfigFile()
  err := hostconfig.Load(pathtoconfigfile)
  if err != nil {
    fmt.Printf("Error loading hostconfig: %s\n",err)
    os.Exit(1)
  }

  err = hostconfig.LookupHostname(flag.Arg(0))
  if err != nil {
    fmt.Println("Invalid Hostname")
    os.Exit(1)
  }

  argslice := flag.Args()
  thishost := argslice[0]
  joinedargs := strings.Join(argslice[1:]," ")

  out := CreateCommandLine( thishost, joinedargs )

  if *donotexecute {
    os.Exit(0)
  }
  bytesout,err := exec.Command("sh","-c",out).Output()
  fmt.Println(string(bytesout))

  os.Exit(0)
}
