package tkvserverlite

import (
  "encoding/binary"
  "encoding/json"
  "fmt"
)

// 44 bytes (serialized to 38 bytes)
// fields ordered for better memory alignment
type Reply struct {
  key         uint64  // byte 8:16
  value       uint64  // byte 16:24
  // Time in seconds in which the key-value will be discarded.
  timeToLive  uint32  // byte 24:28 (seconds)
  // Request Id of the query this reply is for.
  requestId   uint32  // byte 28:32
  // New ip:port of server where to send next query
  newIp       uint32  // byte 32:36
  newPort     uint16  // byte 36:38
  // ip:port where the query was sent.
  serverIp    uint32  // byte 2:6
  serverPort  uint16  // byte 6:8
  // Reply ip and port are used by KvServer's reply goroutine.
  replyIp     uint32  // Not serialized (used internally)
  replyPort   uint16  // Not serialized (used internally)
  status      uint8   // byte 1
  action      uint8   // byte 0
}

func (self *Reply) String() string {
  out, _ := json.Marshal(self)
  return string(out)
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

// Equality test. Used mostly in testing.
func (self *Reply) isEqual(other *Reply) bool {
  return (self.action     == other.action     &&
          self.status     == other.status     &&
          self.serverIp   == other.serverIp   &&
          self.serverPort == other.serverPort &&
          self.key        == other.key        &&
          self.value      == other.value      &&
          self.timeToLive == other.timeToLive &&
          self.requestId  == other.requestId  &&
          // self.newIp      == other.newIp      &&
          // self.newPort    == other.newPort    &&
          self.replyIp    == other.replyIp    &&
          self.replyPort  == other.replyPort)
}


