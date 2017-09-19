package tkvserverlite

import (
  "os"
  "testing"
  "time"
  "net"
  "fmt"
)

var testingServer KvServer
var requestId uint32 = 0

func TestMain(m *testing.M) {
  x := m.Run()

  os.Exit(x)
}

// Add element to the server.
func Test_addElement(t *testing.T) {
  setupKvServer()
  testingServer.Start()

  // t.Error("key", "12", "is present!")
  query := MakeQuery(ACTION_ADD, nextRequestId())
  query.value = 12

  serverAddr,err := net.ResolveUDPAddr("udp","127.0.0.1:10001")
  checkError(err)

  localAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:10002")
  checkError(err)

  conn, err := net.DialUDP("udp", localAddr, serverAddr)
  checkError(err)

  fmt.Printf("Sending %#v\n", query)
  _, err = conn.Write(query.Serialize())
  if err != nil {
      fmt.Println(err)
  }

  time.Sleep(time.Second)

  defer conn.Close()
}


func setupKvServer() {
  testingServer = MakeKvServer(0, "127.0.0.1:10001")
}

func nextRequestId() uint32 {
  requestId += 1
  return requestId
}


