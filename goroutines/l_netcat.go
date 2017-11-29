package main

import (
	"net"
	"log"
	"os"
	"io"
)

func main()  {
	conn,err:=net.Dial("tcp","127.0.0.1:8081")
	if(err!=nil){
		log.Fatal(err)
	}

	defer conn.Close()

	showMsg(os.Stdout,conn)
}

func showMsg(writer io.Writer,conn io.Reader)  {
	if _,err:=io.Copy(writer,conn);err!=nil{
		log.Fatal(err)
	}

}
