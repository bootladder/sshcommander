package hostconfig

import (
  "errors"
  "encoding/json"
  "io/ioutil"
  "fmt"
)

type HostConfigFile struct {
	Name string `json:"name"`
	Host struct {
		Hostname string `json:"hostname"`
		Port     int    `json:"port"`
		Key      string `json:"key"`
		User     string `json:"user"`
	} `json:"host"`
}

var myhostconfig HostConfigFile

func Load(pathtoconfigfile string) (err error) {
    fmt.Println("Loading Conf File into Memory")
    raw, err := ioutil.ReadFile(pathtoconfigfile)
    if err != nil {
        return
    }
    fmt.Printf("%s",raw)
    json.Unmarshal(raw, &myhostconfig)
    fmt.Printf("Results: %v\n", myhostconfig)
    fmt.Printf("name: %v\n", myhostconfig.Name)
    //return errors.New("i'm failing here")
    return nil
}
func LookupHostname(hostname string) (err error) {
  if hostname == myhostconfig.Name {
      return
  }
  return errors.New("i'm failing here")
}
