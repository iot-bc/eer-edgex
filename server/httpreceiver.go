package server

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func Receiver() (string, string) {
	deviceName := ""
	cmd := ""
	return deviceName, cmd
}

func handleConnection(conn net.Conn) {
	br := bufio.NewReader(conn)
	for {
		data, err := br.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Printf("%s", data)
		fmt.Fprintf(conn, "OK\n")
	}
	conn.Close()
}
func main() {
	ln, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal("get client connection error: ", err)
		}
		go handleConnection(conn)
	}
}
