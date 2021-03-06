package hostconfig

import (
  "errors"
  "encoding/json"
  "io/ioutil"
)

type HostConfigFile struct {
	Meta  string `json:"meta"`
	Hosts []struct {
		Name string `json:"name"`
		Behind string `json:"behind"`
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

    raw, err := ioutil.ReadFile(pathtoconfigfile)
    if err != nil {
        return
    }
    err = json.Unmarshal(raw, &myhostconfig)
    if err != nil {
        return
    }
    return
}
func String() string {

    b, err := json.MarshalIndent(myhostconfig, "","  ")
    if err != nil {
        return "Bad JSON"
    }
    return string(b)

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
func HostGetBehind(hostname string) (port string) {

  for i := 0; i < len(myhostconfig.Hosts); i++ {
    if hostname == myhostconfig.Hosts[i].Name {
        return myhostconfig.Hosts[i].Behind
    }
  }
  return ""
}
