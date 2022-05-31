package main

import "net"
import "fmt"
import "bufio"

func main() {
  fmt.Println("Starting socket server")

  // We are listening on port 8000
  ln, _ := net.Listen("tcp", ":8000")

  // Accepting the connection
  conn, _ := ln.Accept()

  // Runing loop forever (until ctrl-c)
  for {
    // get message, output
    message, _ := bufio.NewReader(conn).ReadString('\n')
    fmt.Print("Message Received:", string(message))
  }
}
