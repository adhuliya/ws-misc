package tkvserverlite

import (
    "fmt"
    "net"
    "time"
    "codew/timedkvlite"
    "sync"
)

type KvServer struct {
  // Not in use yet : just in case
  serverId  uint32

  // The timed kv that the server provides access to.
  tkv       *timedkvlite.TimedKv

  // All client queries are put in this chan.
  queries   chan Query

  // Replies taken from this channel are sent
  // to the replyIp and replyPort.
  replies   chan Reply

  // ip and port of this server which is sent to clients.
  ip        uint32
  port      uint16

  // It caches Connection with clients for speed.
  // Key = uint64(reply.replyIp)<<16 | reply.replyPort
  connMap   map[uint64]Connection

  // Mutex for connMap above
  connMapMutex  *sync.Mutex

  // Count of number of messages received.
  msgCount  uint64

  // Number of receivers
  receivers uint8

  // Number of senders
  senders   uint8
}

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
    connMapMutex: &sync.Mutex{},
    msgCount    : 0,
    receivers   : 4,
    senders     : 4,
  }

  return server
}

func (self *KvServer) GetMsgCount() uint64 {
  return self.msgCount
}

func (self *KvServer) incMsgCount() {
  self.msgCount += 1
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

  // Signal that the server has started listening.
  ch <- 1

  // Start receivers
  for i := 0; i < self.receivers; i++ {
    go self.processQueries()
  }

  // Start senders
  for i := 0; i < self.senders; i++ {
    go self.sendReplies()
  }

  // Receive serialized version of Query
  buf := make([]byte, 32)

  for {
      n, _, err := serverConn.ReadFromUDP(buf)
      checkError(err)

      if n != 32 {
        fmt.Println("Error: Query length not 32 bytes, its ", n)
      }

      self.incMsgCount()
      query := DeSerializeQuery(buf)
      fmt.Printf("Received %#v\n", query)

      // send it for processing
      self.queries <- query
  }
}

// Goroutine to process incoming queries.
// Multiple instances of this goroutine are run.
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
          timeToLive  : ttl,    // approximate value (obviously)
          requestId   : query.requestId,
          serverIp    : self.ip,
          serverPort  : self.port,
          replyIp     : query.replyIp,
          replyPort   : query.replyPort,
          // newIp, newPort are set when finally sending reply
          newIp       : 0,
          newPort     : 0,
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

// Goroutine to send replies.
// Multiple instances of this goroutine are run.
func (self *KvServer) sendReplies() {

  for reply := range self.replies {

    conn := self.createConnection(reply.replyIp,
        reply.replyPort, false)

    replyBytes := reply.Serialize()

    _, err := conn.conn.Write(replyBytes)
    for err != nil {
      conn = self.createConnection(reply.replyIp,
          reply.replyPort, true)
      _, err = conn.conn.Write(replyBytes)

    }
  }
}

// It overwrites existing value.
func (self *KvServer) addConnection(key uint64, conn Connection) {
  self.connMapMutex.Lock()
  defer self.connMapMutex.Unlock()

  self.connMap[key] = conn
}

func (self *KvServer) getConnection(key uint64) (Connection, bool) {
  self.connMapMutex.Lock()
  defer self.connMapMutex.Unlock()

  return self.connMap[key]
}

// Generates a Connection object if none exists.
// When forced, it reestablishes the connection with the host.
func (self *KvServer) createConnection(replyIp uint32, replyPort uint16, force bool) Connection {

  // Generate the key to the connMap.
  replySocket := uint64(reply.replyIp)<<16 | uint64(reply.replyPort)

  connection, isPresent := self.getConnection(replySocket)

  if !isPresent {
    // Set up a brand new connection object.
    localAddr, err := net.ResolveUDPAddr("udp",
        getSocketString(self.ip, self.port))
    checkError(err)

    // When replying, remote client is the server.
    remoteAddr, err := net.ResolveUDPAddr("udp",
        getSocketString(replyIp, replyPort))
    checkError(err)

    conn, err := net.DialUDP("udp", localAddr, remoteAddr)
    checkError(err)

    connection := Connection {
      localAddr   : localAddr,
      remoteAddr  : remoteAddr,
      conn        : conn,
    }

    self.addConnection(replySocket, connection)
  } else {
    if force {
      // Reestablish the old connection.
      connection.conn, err := net.DialUDP("udp",
          connection.localAddr,
          connection.remoteAddr)
      checkError(err)

      self.addConnection(connection)
    }
  }

  return connection
}


