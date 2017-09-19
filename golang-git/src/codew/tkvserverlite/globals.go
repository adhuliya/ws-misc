package tkvserverlite

import (
  "codew/timedkvlite"
  "encoding/binary"
  "encoding/json"
  "strings"
  "strconv"
  "net"
  "fmt"
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


  // Count of number of messages received.
  count     uint64
}

// 32 bytes
// fields ordered for better memory alignment
type Query struct {
  key         uint64  // byte 8:16
  value       uint64  // byte 16:24
  duration    uint32  // byte 24:28 (seconds)
  requestId   uint32  // byte 28:32
  replyIp     uint32  // byte 2:6
  replyPort   uint16  // byte 6:8
  empty       uint8   // byte 1
  action      uint8   // byte 0
}

// 44 bytes (serialized to 38 bytes)
// fields ordered for better memory alignment
type Reply struct {
  key         uint64  // byte 8:16
  value       uint64  // byte 16:24
  timeToLive  uint32  // byte 24:28 (seconds)
  requestId   uint32  // byte 28:32
  // new ip:port of server where to send query
  newIp       uint32  // byte 32:36
  newPort     uint16  // byte 36:38
  // ip:port where query was sent
  serverIp    uint32  // byte 2:6
  serverPort  uint16  // byte 6:8
  replyIp     uint32  // Not serialized
  replyPort   uint16  // Not serialized
  status      uint8   // byte 1
  action      uint8   // byte 0
}

type Connection struct {
  localAddr   *net.UDPAddr
  remoteAddr  *net.UDPAddr
  conn        *net.UDPConn
}


const (
  ACTION_ADD       byte = 1
  ACTION_DELETE    byte = 2
  ACTION_MODIFY_VD byte = 3 // modify value and duration
  ACTION_MODIFY_V  byte = 4 // modiry value
  ACTION_MODIFY_D  byte = 5 // modiry duration
  ACTION_GET       byte = 6

  STATUS_OK        byte = 0
  STATUS_NOKEY     byte = 1 // no key found
)

func (self *Query) String() string {
  out, _ := json.Marshal(self)
  return string(out)
}

func (self *Reply) String() string {
  out, _ := json.Marshal(self)
  return string(out)
}

func MakeQuery(action uint8, requestId uint32) Query {
  return Query {
    action    : action,
    requestId : requestId,
  }
}

func MakeConnection(localIp uint32, localPort uint16,
    remoteIp uint32, remotePort uint16) Connection {
  localAddr, _  := net.ResolveUDPAddr("udp", getSocketString(localIp, localPort))
  remoteAddr, _ := net.ResolveUDPAddr("udp", getSocketString(remoteIp, remotePort))
  conn, _       := net.DialUDP("udp", localAddr, remoteAddr)

  return Connection {
    localAddr   : localAddr,
    remoteAddr  : remoteAddr,
    conn        : conn,
  }
}

// serialize Reply into bytes
func (reply *Reply) Serialize() []byte {
  bytes := make([]byte, 38, 38)

  bytes[0] = reply.action
  bytes[1] = reply.status
  binary.BigEndian.PutUint32(bytes[2:6], reply.serverIp)
  binary.BigEndian.PutUint16(bytes[6:8], reply.serverPort)
  binary.BigEndian.PutUint64(bytes[8:16], reply.key)
  binary.BigEndian.PutUint64(bytes[16:24], reply.value)
  binary.BigEndian.PutUint32(bytes[24:28], reply.timeToLive)
  binary.BigEndian.PutUint32(bytes[28:32], reply.requestId)
  binary.BigEndian.PutUint32(bytes[32:36], reply.newIp)
  binary.BigEndian.PutUint16(bytes[36:38], reply.newPort)

  return bytes
}

// deserialize the Reply from bytes
// Given bytes must be exactly 32 bytes long.
func DeSerializeReply(bytes []byte) Reply {
  if len(bytes) != 38 {
    fmt.Println("Error : bytes length is not 38. Its ", len(bytes))
  }

  return Reply {
    action    : bytes[0],
    status    : 0,
    serverIp  : binary.BigEndian.Uint32(bytes[2:6]),
    serverPort: binary.BigEndian.Uint16(bytes[6:8]),
    key       : binary.BigEndian.Uint64(bytes[8:16]),
    value     : binary.BigEndian.Uint64(bytes[16:24]),
    timeToLive: binary.BigEndian.Uint32(bytes[24:28]),
    requestId : binary.BigEndian.Uint32(bytes[28:32]),
    newIp     : binary.BigEndian.Uint32(bytes[32:36]),
    newPort   : binary.BigEndian.Uint16(bytes[36:38]),
    replyIp   : 0,
    replyPort : 0,
  }
}

// serialize query into bytes
func (query *Query) Serialize() []byte {
  bytes := make([]byte, 32, 32)

  bytes[0] = query.action
  bytes[1] = query.empty
  binary.BigEndian.PutUint32(bytes[2:6], query.replyIp)
  binary.BigEndian.PutUint16(bytes[6:8], query.replyPort)
  binary.BigEndian.PutUint64(bytes[8:16], query.key)
  binary.BigEndian.PutUint64(bytes[16:24], query.value)
  binary.BigEndian.PutUint32(bytes[24:28], query.duration)
  binary.BigEndian.PutUint32(bytes[28:32], query.requestId)

  return bytes
}

// deserialize the Query from bytes
// Given bytes must be exactly 32 bytes long.
func DeSerializeQuery(bytes []byte) Query {
  if len(bytes) != 32 {
    fmt.Println("Error : bytes length is not 32. Its ", len(bytes))
  }

  return Query {
    action    : bytes[0],
    empty     : 0,
    replyIp   : binary.BigEndian.Uint32(bytes[2:6]),
    replyPort : binary.BigEndian.Uint16(bytes[6:8]),
    key       : binary.BigEndian.Uint64(bytes[8:16]),
    value     : binary.BigEndian.Uint64(bytes[16:24]),
    duration  : binary.BigEndian.Uint32(bytes[24:28]),
    requestId : binary.BigEndian.Uint32(bytes[28:32]),
  }
}

func getSocketString(ip uint32, port uint16) string {
  return fmt.Sprintf("%d.%d.%d.%d:%d", 
    (ip & 0xFF000000) >> 24,
    (ip & 0x00FF0000) >> 16,
    (ip & 0x0000FF00) >> 8,
    (ip & 0x000000FF),
    port)
}

// address in the format 127.0.0.1:18202
// or 127.0.9.8 hence default port 10001
// or :18202 hence default ip 127.0.0.1
func extractIpPort(address string) (uint32, uint16, bool) {
  // defaults
  ip    := ipUint32(127,0,0,1)
  port  := uint16(10001)

  address = strings.Trim(address, " \t")

  if len(address) == 0 {
    // No address given, hence assume defaults
    return ip, port, true
  }

  if strings.HasPrefix(address, ":") {
    // Only port number given
    p, _ := strconv.Atoi(address[1:])
    return ip, uint16(p), true
  }

  splitIpPort := strings.Split(address, ":")

  if len(splitIpPort) == 1 {
    // Only ip address given
    splitIp := strings.Split(address, ".")
    b1, _ := strconv.Atoi(splitIp[0])
    b2, _ := strconv.Atoi(splitIp[1])
    b3, _ := strconv.Atoi(splitIp[2])
    b4, _ := strconv.Atoi(splitIp[3])

    ip = ipUint32(uint8(b1), uint8(b2), uint8(b3), uint8(b4))

    return ip, port, true
  } else {
    // both ip:port given
    splitIp := strings.Split(splitIpPort[0], ".")
    b1, _ := strconv.Atoi(splitIp[0])
    b2, _ := strconv.Atoi(splitIp[1])
    b3, _ := strconv.Atoi(splitIp[2])
    b4, _ := strconv.Atoi(splitIp[3])

    ip = ipUint32(uint8(b1), uint8(b2), uint8(b3), uint8(b4))
    p, _ := strconv.Atoi(splitIpPort[1])

    return ip, uint16(p), true
  }

  return ip, port, false
}

// converts four separate bytes to a single uint32 ip
func ipUint32(byte0 uint8, byte1 uint8, byte2 uint8, byte3 uint8) uint32 {
  var ip uint32 = 0

  ip |= uint32(byte0) << 24
  ip |= uint32(byte1) << 16
  ip |= uint32(byte2) << 8
  ip |= uint32(byte3)

  return ip
}

// Error check
func checkError(err error) {
    if err  != nil {
        fmt.Println("Error: " , err)
    }
}


