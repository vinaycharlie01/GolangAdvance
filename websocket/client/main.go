package main

import (
	"fmt"
	"log"
	"net"
)

func WriteandRead(con net.Conn) {
	con.Write([]byte("Vinay"))

	// a := []int{1, 2, 4, 5, 6, 7}
	reader := make([]byte, 1012)
	n, err := con.Read(reader)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Reciver :", string(reader[:n]))
}

const (
	N = "tcp"
	A = "localhost:8080"
)

func main() {

	conn, err := net.Dial(N, A)
	if err != nil {
		log.Fatal(err)
	}

	WriteandRead(conn)

}
