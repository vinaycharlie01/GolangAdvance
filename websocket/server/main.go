package main

import (
	"fmt"
	"log"
	"net"
	"sync"
)

const (
	N = "tcp"
	A = "localhost:8080"
)

func ReadAndWrite(conn net.Conn) {
	reader := make([]byte, 0)
	n, err := conn.Read(reader)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Recived form client", string(reader[:n]))
	conn.Write([]byte("Hello Hi" + string(reader[:n])))
}

func main() {
	lis, err := net.Listen(N, A)
	if err != nil {
		log.Fatal(err)
	}
	// for {
	conn, err := lis.Accept()

	if err != nil {
		log.Fatal(err)
	}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		ReadAndWrite(conn)
	}()
	wg.Wait()

	// }

}
