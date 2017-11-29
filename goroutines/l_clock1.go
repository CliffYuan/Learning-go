package main

import (
	"net"
	"log"
	"io"
	"time"
	"fmt"
)

//基于单goroutine的服务端
func main()  {
	lis,err:=net.Listen("tcp","127.0.0.1:8081")
	if(err!=nil){
		log.Fatal(err)
	}

	for {
		conn,err:=lis.Accept()
		if(err!=nil){
			log.Println(err)
			continue
		}

		handle(conn)

	}
}

func handle(conn net.Conn)  {
	defer conn.Close()

	for  {
		l,err:=io.WriteString(conn,time.Now().Format("15:04:05"))

		if(err!=nil){
			log.Println(err)
			return
		}
		fmt.Printf("write %d byte to client",l)
		time.Sleep(1000*time.Millisecond)
	}
}
