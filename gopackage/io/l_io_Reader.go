package main

import (
	"os"
	"log"
	"fmt"
	"io/ioutil"
)

func main()  {

	//1.读一个文件

	//使用io.Read()
	readerFile();

	fmt.Println("----------------------------------------")

	//使用ioutil.ReadAll()
	readerFileUseIoutilReadAll()

}

func readerFile()  {
	file,err:=os.Open("/tmp/gofile.log")
	if(err!=nil){
		log.Fatal(err)
	}
	defer file.Close()

	s:= make([]byte,100);

	l,err:=file.Read(s);
	if(err!=nil){
		log.Fatal(err)
	}

	fmt.Println("read %d byte from gofile.log,data is %s",l,string(s))
}

func readerFileUseIoutilReadAll()  {
	file,err:=os.Open("/tmp/gofile.log")
	if(err!=nil){
		log.Fatal(err)
	}
	defer file.Close()

	data,err:=ioutil.ReadAll(file)
	if(err!=nil){
		log.Fatal(err)
	}

	fmt.Println("use ioutil.ReadAll from gofile.log,data is %s",string(data))
}
