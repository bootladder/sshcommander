package hostconfig_test

import(
  "github.com/bootladder/sshcommander/hostconfig"
  "github.com/stretchr/testify/assert"
  "os/user"
  "fmt"
  "testing"
)

func TestLoadDefaultConfig( t *testing.T) {

  usr, _ := user.Current()
  dir := usr.HomeDir
  err := hostconfig.Load(dir + "/.sshcommander/hostconfig.json")
  assert.Equal(t, nil, err)
}

func TestLoadConfigInvalidFilename( t *testing.T) {

  err := hostconfig.Load("/my/crappy/filename")
  assert.NotEqual(t, nil, err)
}

func TestLookupHostnameInvalidReturnsError( t *testing.T) {

  err := hostconfig.Load("./samplehostconfig.json")
  err = hostconfig.LookupHostname("invalidhostname")
  assert.NotEqual(t, nil, err)
}

func TestLookupHostnameValid( t *testing.T) {

  err := hostconfig.Load("./samplehostconfig.json")
  if err != nil {
    fmt.Println(err)
  }
  err = hostconfig.LookupHostname("myfirsthostname")
  assert.Equal(t, nil, err)
  err = hostconfig.LookupHostname("mysecondhostname")
  assert.Equal(t, nil, err)
}
