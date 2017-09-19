package tkvserverlite

import (
    "fmt"
    "net"
    "os"
    "encoding/binary"
)

func CreateServer(serverId uint32, address string) {
  ip, port, _ := extractIpPort(address)

  server := KvServer {
    serverId    : serverId,
    tkv         : timedkvlite.Make(),
    queries     : make(chan Query, 1024),
    replies     : make(chan Reply, 1024),
    ip          : ip,
    port        : port,
    connMap     : make(map[uint64]Connection),
  }

  return server
}

// a goroutine that goes into an indefinite loop
func (server *KvServer) Start() {
  // Prepare server
  ServerAddr, err := net.ResolveUDPAddr("udp",server.address)
  checkError(err)

  // Now listen at selected port
  ServerConn, err := net.ListenUDP("udp", ServerAddr)
  checkError(err)
  defer ServerConn.Close()

  // Receive serialized version of Query
  buf := make([]byte, 32)

  for {
      n, addr, err := ServerConn.ReadFromUDP(buf)
      chechError(err)

      if n != 32 {
        fmt.Println("Error: Query length not 32, its ", n)
      }

      query := DeSerializeQuery(buf)

      // send it for processing
      server.queries <- query
  }

}

// goroutine to process queries
// multiple instances of this goroutine are run
func (server *KvServer) processQueries() {
  for query := range server.queries {
    if query.action == ACTION_GET {
    }
    if query.action == ACTION_DELETE {
    }
    if query.action == ACTION_MODIFY_VD {
    }
    if query.action == ACTION_MODIFY_V {
    }
    if query.action == ACTION_MODIFY_D {
    }
    if query.action == ACTION_ADD {
    }
  }
}

/* A Simple function to verify error */
func checkError(err error) {
    if err  != nil {
        fmt.Println("Error: " , err)
        os.Exit(0)
    }
}


