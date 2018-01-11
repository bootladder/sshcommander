package hostconfig

import (
  "errors"
  "encoding/json"
  "io/ioutil"
  "fmt"
)

type HostConfigFile struct {
	Meta  string `json:"meta"`
	Hosts []struct {
		Name string `json:"name"`
		Host struct {
			Hostname string `json:"hostname"`
			Port     string `json:"port"`
			Key      string `json:"key"`
			User     string `json:"user"`
		} `json:"host"`
	} `json:"hosts"`
}

var myhostconfig HostConfigFile

func Load(pathtoconfigfile string) (err error) {
    fmt.Println("Loading Conf File into Memory")
    raw, err := ioutil.ReadFile(pathtoconfigfile)
    if err != nil {
        return
    }
    json.Unmarshal(raw, &myhostconfig)
    return
}
func LookupHostname(hostname string) (err error) {
  for i := 0; i < len(myhostconfig.Hosts); i++ {
    if hostname == myhostconfig.Hosts[i].Name {
        return
    }
  }
  return errors.New("Hostname Not Found")
}


func HostGetPort(hostname string) (port string) {

  for i := 0; i < len(myhostconfig.Hosts); i++ {
    if hostname == myhostconfig.Hosts[i].Name {
        return myhostconfig.Hosts[i].Host.Port
    }
  }
  return ""
}
func HostGetHostname(hostname string) (port string) {

  for i := 0; i < len(myhostconfig.Hosts); i++ {
    if hostname == myhostconfig.Hosts[i].Name {
        return myhostconfig.Hosts[i].Host.Hostname
    }
  }
  return ""
}
func HostGetUser(hostname string) (port string) {

  for i := 0; i < len(myhostconfig.Hosts); i++ {
    if hostname == myhostconfig.Hosts[i].Name {
        return myhostconfig.Hosts[i].Host.User
    }
  }
  return ""
}
func HostGetKey(hostname string) (port string) {

  for i := 0; i < len(myhostconfig.Hosts); i++ {
    if hostname == myhostconfig.Hosts[i].Name {
        return myhostconfig.Hosts[i].Host.Key
    }
  }
  return ""
}
