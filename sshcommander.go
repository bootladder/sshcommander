package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"strings"

	"github.com/bootladder/sshcommander/hostconfig"
	"github.com/bootladder/sshcommander/sshcommandcreator"
)

var globalInitFlag bool = true
var donotexecute *bool
var pty_dash_t_flag *bool
var list_hosts_flag *bool
var addionalargs string = ""

func GetPathToConfigFile() (path string) {

	usr, _ := user.Current()
	dir := usr.HomeDir
	return dir + "/.sshcommander/hostconfig.json"
}

func CreateCommandLine(thishost string, joinedargs string) (out string) {

	if globalInitFlag {
		globalInitFlag = false
		if *pty_dash_t_flag {
			addionalargs = "-t"
		}
	}
	port := hostconfig.HostGetPort(thishost)
	user := hostconfig.HostGetUser(thishost)
	key := hostconfig.HostGetKey(thishost)
	hostname := hostconfig.HostGetHostname(thishost)

	creator := sshcommandcreator.SSHCommandCreator{}
	creator.Port = port
	creator.User = user
	creator.Hostname = hostname
	creator.Key = key
	creator.AdditionalArgs = addionalargs

	out, _ = creator.CreateCommandString(joinedargs)
	//fmt.Println(out)

	if hostconfig.HostGetBehind(thishost) != "" {
		out = CreateCommandLine(hostconfig.HostGetBehind(thishost), out)
	}
	return
}

func main() {
	donotexecute = flag.Bool("N", false, "Don't execute, just print the command")
	pty_dash_t_flag = flag.Bool("t", false, "pass -t to ssh, for pseudo-tty")
	list_hosts_flag = flag.Bool("l", false, "Lists Hosts from Config File")

	flag.Parse()
	//if flag.NArg() < 2 {
	//  fmt.Println("need atleast 2 arguments after flags, hostname and command")
	//  flag.Usage()
	//  os.Exit(1)
	//}
	pathtoconfigfile := GetPathToConfigFile()
	err := hostconfig.Load(pathtoconfigfile)
	if err != nil {
		fmt.Printf("Error loading hostconfig: %s\n", err)
		os.Exit(1)
	}

	if *list_hosts_flag {
		fmt.Println("Configured Hosts:")
		fmt.Println(hostconfig.String())
		os.Exit(0)
	}

	err = hostconfig.LookupHostname(flag.Arg(0))
	if err != nil {
		fmt.Println("Invalid Hostname")
		os.Exit(1)
	}

	argslice := flag.Args()
	thishost := argslice[0]
	joinedargs := strings.Join(argslice[1:], " ")

	out := CreateCommandLine(thishost, joinedargs)
	fmt.Println(out)

	if *donotexecute {
		os.Exit(0)
	}

	mystrings := strings.Fields(out)
	cmd := exec.Command(mystrings[0], mystrings[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Run()

	os.Exit(0)
}
