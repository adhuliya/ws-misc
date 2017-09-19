package tkvserverlite

import (
  "encoding/binary"
  "encoding/json"
  "fmt"
)

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

func MakeQuery(action uint8, requestId uint32) Query {
  return Query {
    action    : action,
    requestId : requestId,
  }
}

func (self *Query) String() string {
  out, _ := json.Marshal(self)
  return string(out)
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

