package tkvserverlite

import (
    "fmt"
    "net"
    "time"
    "codew/timedkvlite"
)

func MakeKvServer(serverId uint32, address string) KvServer {
  // TODO replace _ below with variable
  ip, port, _ := extractIpPort(address)

  server := KvServer {
    serverId    : serverId,
    tkv         : timedkvlite.Make(),
    queries     : make(chan Query, 1024),
    replies     : make(chan Reply, 1024),
    ip          : ip,
    port        : port,
    // Maps ip:port of client to its Connection struct
    // ip:port is converted to uint64 by using lower 6 bytes
    // upper 2 bytes are zero
    connMap     : make(map[uint64]Connection),
    count       : 0,
  }

  return server
}

func (self *KvServer) GetCount() uint64 {
  return self.count
}

func (self *KvServer) incCount() {
  self.count += 1
}

// A funtion that starts the server, and returns when started.
func (self *KvServer) Start() {
  ch := make(chan int32, 1)
  go self.run(ch)
  // Wait untill the server starts.
  <-ch
}

func (self *KvServer) run(ch chan int32) {
  // Prepare server
  serverAddr, err := net.ResolveUDPAddr("udp",
    getSocketString(self.ip, self.port))
  checkError(err)

  // Now listen at selected port
  serverConn, err := net.ListenUDP("udp", serverAddr)
  checkError(err)
  defer serverConn.Close()

  go self.processQueries()
  //go self.sendReplies()

  // Receive serialized version of Query
  buf := make([]byte, 32)

  // signal the start
  ch <- 1

  for {
      n, _, err := serverConn.ReadFromUDP(buf)
      checkError(err)

      if n != 32 {
        fmt.Println("Error: Query length not 32 bytes, its ", n)
      }

      self.incCount()
      query := DeSerializeQuery(buf)
      fmt.Printf("Received %#v\n", query)

      // send it for processing
      self.queries <- query
  }
}

// goroutine to process queries
// multiple instances of this goroutine are run
func (self *KvServer) processQueries() {
  for query := range self.queries {
    shouldReply := true
    if query.replyIp == 0 || query.replyPort == 0 {
      shouldReply = false
    }

    if query.action == ACTION_GET {
      value, _ := self.tkv.Get(query.key)

      // Calculate time to live in seconds.
      now := time.Now().UnixNano()
      ttl := uint32(0)
      if value.GetDelTime() > now {
        ttl = uint32((value.GetDelTime() - now)/1000000000)
      }

      if shouldReply {
        self.replies <- Reply {
          action      : query.action,
          status      : STATUS_OK,
          key         : query.key,
          value       : value.GetValue(),
          timeToLive  : ttl,    // approximate value
          requestId   : query.requestId,
          serverIp    : self.ip,
          serverPort  : self.port,
          replyIp     : query.replyIp,
          replyPort   : query.replyPort,
          // newIp, newPort are set when extracting from chan server.replies
        }
      }
    }
    // if query.action == ACTION_DELETE {
    // }
    // if query.action == ACTION_MODIFY_VD {
    // }
    // if query.action == ACTION_MODIFY_V {
    // }
    // if query.action == ACTION_MODIFY_D {
    // }
    // if query.action == ACTION_ADD {
    // }
  }
}

// goroutine to send replies
// multiple instances of this goroutine are run
//func (self *KvServer) sendReplies() {
//  for reply := range self.replies {
//    client := uint64(reply.replyIp)<<16 | uint64(reply.replyPort)
//  }
//}


