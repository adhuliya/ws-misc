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
var queryList []Query
var replyList []Reply

func TestMain(m *testing.M) {
  x := m.Run()

  os.Exit(x)
}

// Add element to the server.
func Test_addElement(t *testing.T) {
  // initialize test data
  initializeQueryReply1()

  // Setup and start the KvServer.
  setupKvServer()

  conn := setupClientSender()
  defer conn.Close()

  // Start the channel and receiver that 
  // accepts replies from KvServer.
  replies := make(chan Reply, 1)
  go receiveFromKvServer(replies)

  time.Sleep(time.Second)

  // fmt.Printf("Sending %#v\n", query)
  _, err := conn.Write(queryList[0].Serialize())
  if err != nil {
      fmt.Println(err)
  }

  reply := <-replies

  if ! replyList[0].isEqual(&reply) {
    t.Error("Unexpected Reply : Expected ", fmt.Sprintf("%#v", replyList[0]), "Got ", fmt.Sprintf("%#v", reply))
  }
}

func setupClientSender() *net.UDPConn {
  serverAddr,err := net.ResolveUDPAddr("udp","127.0.0.1:10001")
  checkError(err)

  localAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:10003")
  checkError(err)

  conn, err := net.DialUDP("udp", localAddr, serverAddr)
  checkError(err)

  return conn
}

// Goroutine to receive replies from the KvServer
// DeSerialize and send it into the chan replies.
func receiveFromKvServer(replies chan Reply) {
  // Prepare server
  serverAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:10004")
  checkError(err)

  // Now listen at selected port
  serverConn, err := net.ListenUDP("udp", serverAddr)
  checkError(err)
  defer serverConn.Close()

  // Receive serialized version of Reply
  buf := make([]byte, 38)

  for {
      n, _, err := serverConn.ReadFromUDP(buf)
      checkError(err)

      if n != 38 {
        fmt.Println("Error: Query length not 32 bytes, its ", n)
      }

      reply := DeSerializeReply(buf)
      //fmt.Printf("Received %#v\n", reply)

      // send it for processing
      replies <- reply
  }
}

func setupKvServer() {
  testingServer = MakeKvServer(0, "127.0.0.1:10001")
  testingServer.Start()
}

func nextRequestId() uint32 {
  requestId += 1
  return requestId
}

func initializeQueryReply1() {
  queryList = make([]Query, 1)
  replyList = make([]Reply, 1)

  queryList[0] = Query {
    action      : ACTION_ADD,
    empty       : 0,
    value       : 12,
    key         : 0,
    timeToLive  : 10,
    requestId   : 1,
    replyIp     : ipUint32(127,0,0,1),
    replyPort   : uint16(10004),
  }

  replyList[0] = Reply {
    action      : ACTION_ADD,
    status      : STATUS_OK,
    value       : 12,
    key         : 1,    // first elemtn will have key 1
    timeToLive  : 10,
    requestId   : 1,
    serverIp    : ipUint32(127,0,0,1),
    serverPort  : uint16(10001),
    newIp       : ipUint32(127,0,0,1),
    newPort     : uint16(10001),
    // ReplyIp and replyPort are used internally by the server.
    // Hence they are not checked.
  }
}
