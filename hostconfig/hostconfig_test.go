package hostconfig_test

import(
  "github.com/bootladder/sshcommander/hostconfig"
  "github.com/stretchr/testify/assert"
  "fmt"
  "testing"
)

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

func TestGettersForValidHostname( t *testing.T) {

  err := hostconfig.Load("./samplehostconfig.json")
  if err != nil {
    fmt.Println(err)
  }

  assert.Equal(t, "22", hostconfig.HostGetPort("myfirsthostname"))
  assert.Equal(t, "22", hostconfig.HostGetPort("mysecondhostname"))
  assert.Equal(t, "localhost", hostconfig.HostGetHostname("mysecondhostname"))
  assert.Equal(t, "steve", hostconfig.HostGetUser("myfirsthostname"))
  assert.Equal(t, "", hostconfig.HostGetKey("myfirsthostname"))
  assert.Equal(t, "myfirsthostname", hostconfig.HostGetBehind("myfirstnestedhostname"))

}
