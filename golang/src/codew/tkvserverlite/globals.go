package tkvserverlite

import (
  "strings"
  "strconv"
  "net"
  "fmt"
)

const (
  ACTION_ADD       byte = 1
  ACTION_DELETE    byte = 2
  ACTION_MODIFY_VD byte = 3 // modify value and duration
  ACTION_MODIFY_V  byte = 4 // modify value
  ACTION_MODIFY_D  byte = 5 // modify duration
  ACTION_GET       byte = 6

  STATUS_OK        byte = 0
  STATUS_NOKEY     byte = 1 // no key found
)

type Connection struct {
  localAddr   *net.UDPAddr
  remoteAddr  *net.UDPAddr
  conn        *net.UDPConn
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


