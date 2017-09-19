package kvserverlite

import (
  "codew/timedkvlite"
  "encoding/binary"
  "strings"
  "strconv"
)

type KvServer struct {
  serverId  uint32
  tkv       timedkvlite.TimedKv
  queries   chan Query
  replies   chan Reply
  ip        uint32
  port      uint16
  connMap   map[uint64]Connection
}

// 32 bytes
type Query struct {
  action1     uint8   // byte 0
  empty       uint8   // byte 1
  replyIp     uint32  // byte 2:6
  replyPort   uint16  // byte 6:8
  key         uint64  // byte 8:16
  value       uint64  // byte 16:24
  duration    uint32  // byte 24:28 (seconds)
  requestId   uint32  // byte 28:32
}

// 32 bytes
type Reply struct {
  action      uint8   // byte 0
  status      uint8   // byte 1
  serverIp    uint32  // byte 2:6
  serverPort  uint16  // byte 6:8
  key         uint64  // byte 8:16
  value       uint64  // byte 16:24
  duration    uint32  // byte 24:28 (seconds)
  requestId   uint32  // byte 28:32
}

type Connection struct {
  localAddr   net.UDPAddr,
  remoteAddr  net.UDPAddr,
  conn        net.UDPConn,
}


const (
  ACTION_ADD       byte = 1
  ACTION_DELETE    byte = 2
  ACTION_MODIFY_VD byte = 3 // modify value and duration
  ACTION_MODIFY_V  byte = 4 // modiry value
  ACTION_MODIFY_D  byte = 5 // modiry duration
  ACTION_GET       byte = 6

  STATUS_SUCCESS   byte = 0
  STATUS_NOKEY     byte = 1 // no key found
)

// serialize Reply into bytes
func (reply *Reply) Serialize() []byte {
  bytes := make([]byte, 32, 32)

  bytes[0] = reply.action
  bytes[1] = reply.status
  binary.BigEndian.PutUint32(bytes[2:6], reply.serverIp)
  binary.BigEndian.PutUint16(bytes[6:8], reply.serverPort)
  binary.BigEndian.PutUint64(bytes[8:16], reply.key)
  binary.BigEndian.PutUint64(bytes[16:24], reply.value)
  binary.BigEndian.PutUint32(bytes[24:28], reply.duration)
  binary.BigEndian.PutUint32(bytes[28:32], reply.requestId)

  return bytes
}

// deserialize the Reply from bytes
// Given bytes must be exactly 32 bytes long.
func DeSerializeReply(bytes []byte) Reply {
  if len(bytes) != 32 {
    fmt.Println("Error : bytes length is not 32. Its ", len(bytes))
  }

  return Reply {
    action    : binary.BigEndian.Uint16(bytes[0]),
    status    : 0,
    serverIp  : binary.BinEndian.Uint32(bytes[2:6]),
    serverPort: binary.BinEndian.Uint16(bytes[6:8]),
    key       : binary.BinEndian.Uint64(bytes[8:16]),
    value     : binary.BinEndian.Uint64(bytes[16:24]),
    duration  : binary.BinEndian.Uint32(bytes[24:28]),
    requestId : binary.BinEndian.Uint32(bytes[28:32]),
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
    action    : binary.BigEndian.Uint16(bytes[0]),
    empty     : 0,
    replyIp   : binary.BinEndian.Uint32(bytes[2:6]),
    replyPort : binary.BinEndian.Uint16(bytes[6:8]),
    key       : binary.BinEndian.Uint64(bytes[8:16]),
    value     : binary.BinEndian.Uint64(bytes[16:24]),
    duration  : binary.BinEndian.Uint32(bytes[24:28]),
    requestId : binary.BinEndian.Uint32(bytes[28:32]),
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
func extractIpPort(address string) uint32, uint16, bool {
  // defaults
  ip    := ipUint32(127,0,0,1)
  port  := uint16(10001)

  address = strings.Trim(address, " \t")

  if len(address) == 0 {
    // No address given, hence assume defaults
    return ip, port, true
  }

  if strings.HasPrefix(":") {
    // Only port number given
    port = strconv.Atoi(address[1:])
    return ip, port, true
  }

  splitIpPort := strings.Split(address, ":")

  if len(splitIpPort) == 1 {
    // Only ip address given
    splitIp := strings.Split(address, ".")
    ip = ipUint32(uint8(strconv.Atoi(splitIp[0])),
                  uint8(strconv.Atoi(splitIp[1])),
                  uint8(strconv.Atoi(splitIp[2])),
                  uint8(strconv.Atoi(splitIp[3])))

    return ip, port, true
  } else {
    // both ip:port given
    splitIp := strings.Split(splitIpPort[0], ".")
    ip = ipUint32(uint8(strconv.Atoi(splitIp[0])),
                  uint8(strconv.Atoi(splitIp[1])),
                  uint8(strconv.Atoi(splitIp[2])),
                  uint8(strconv.Atoi(splitIp[3])))
    port = uint16(strconv.Atoi(splitIpPort[1]))
    return ip, port, true
  }

  return ip, port, false
}

// converts four separate bytes to a single uint32 ip
func ipUint32(byte0 uint8, byte1 uint8, byte2 uint8, byte3 uint8) {
  var ip uint32 = 0

  ip |= uint32(byte0) << 24
  ip |= uint32(byte1) << 16
  ip |= uint32(byte2) << 8
  ip |= uint32(byte3)

  return ip
}


