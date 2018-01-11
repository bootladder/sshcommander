package hostconfig

import (
  //"errors"
  "io/ioutil"
  "fmt"
)


func Load(pathtoconfigfile string) (err error) {
    raw, err := ioutil.ReadFile(pathtoconfigfile)
    if err != nil {
        fmt.Println(err.Error())
        return
    }
    fmt.Println("%s",raw)

  //return errors.New("i'm failing here")
  return nil
}
func LookupHostname(hostname string) (err error) {
  return nil
}
