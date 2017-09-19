package main

import (
    "fmt"
    "net"
    "os"
    "time"
//    "strconv"
)

/* A Simple function to verify error */
func CheckError(err error) {
    if err  != nil {
        fmt.Println("Error: " , err)
        os.Exit(0)
    }
}

func main() {
    /* Lets prepare a address at any address at port 10001*/
    ServerAddr,err := net.ResolveUDPAddr("udp","127.0.0.1:10001")
    CheckError(err)

    /* Now listen at selected port */
    ServerConn, err := net.ListenUDP("udp", ServerAddr)
    CheckError(err)
    defer ServerConn.Close()

    buf := make([]byte, 1024)

    go client()

    for {
        n,addr,err := ServerConn.ReadFromUDP(buf)
        fmt.Println("Received ",string(buf[0:n]), " from ",addr)

        if err != nil {
            fmt.Println("Error: ",err)
        }
    }
}


func client() {
   ServerAddr,err := net.ResolveUDPAddr("udp","127.0.0.1:10001")
   CheckError(err)

   LocalAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:10002")
   CheckError(err)

   Conn, err := net.DialUDP("udp", LocalAddr, ServerAddr)
   CheckError(err)

   defer Conn.Close()
   buf := make([]byte, 1, 1)
   buf[0] = 65
   //i := 0
   for {
       //msg := strconv.Itoa(i)
       //i++
       // buf := []byte(msg)
       _,err := Conn.Write(buf)
       if err != nil {
           //fmt.Println(msg, err)
           fmt.Println(err)
       }
       time.Sleep(time.Second * 1)
   }
}


