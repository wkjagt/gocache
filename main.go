package main

import (
    "fmt"
    "net"
    "os"
    "github.com/wkjagt/gocache/cache"
)

const (
    connHost = "localhost"
    connPort = "3333"
    connType = "tcp"
)

func main() {
    listener, err := net.Listen(connType, fmt.Sprintf("%s:%s", connHost, connPort))
    if err != nil {
        fmt.Println("Error listening:", err.Error())
        os.Exit(1)
    }
    defer listener.Close()

    fmt.Println("Listening on " + connHost + ":" + connPort)

    for {
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println("Error accepting: ", err.Error())
            os.Exit(1)
        }
        go handleRequest(conn)
    }
}

func handleRequest(conn net.Conn) {
  defer conn.Close()

	payLoad, _ := readPayload(conn)
	command, err := cache.ExtractCommand(payLoad)

  if err != nil {
    conn.Write([]byte(err.Error()));
    return
  }

	res, err := command.Handle()

	if err != nil {
		conn.Write([]byte(err.Error()));
	} else {
		conn.Write([]byte(res));
	}
}

func readPayload(conn net.Conn) (string, error) {
  buffer := make([]byte, 1024) // a slice with size and capacity of 1024

  len, err := conn.Read(buffer)

  if err != nil {
    return "", fmt.Errorf("Error reading:%s", err.Error())
  }
	return string(buffer[:len-1]), nil
}
