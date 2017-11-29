package main

import (
	"os"
	"log"
	"fmt"
)

func main()  {
	//写文件
	writeFile();
}

func writeFile()  {
	file,err:=os.Create("/tmp/gofile_writer.log")
	if(err!=nil){
		log.Fatal(err)
	}
	defer file.Close()


	data:=[]byte("data is write by io.Writer")

	l,err:=file.Write(data)
	if(err!=nil){
		log.Fatal(err)
	}

	fmt.Println("use io.Writer write data %d to gofile_writer.log",l)

}
