package main

import (
	"fmt"
	"time"
)

func main()  {
	
	varInit();

	fmt.Println("---------------------------------opt not init var--------------------------------------------")

	optNotInitVar();

}

func varInit()  {
	//int
	var varint int
	fmt.Println("int:",varint)

	//string
	var varstring string
	fmt.Println("string:",varstring)

	//array
	var vararray [5]int
	fmt.Println("array:",vararray)

	//struct
	type person struct {
		age int
		name string
	}
	var varstruct person
	fmt.Println("struct:",varstruct)


	fmt.Println("---------------------------------default is nil--------------------------------------------")

	//interface
	var varinterface interface{}
	if varinterface==nil  {
		fmt.Println("nil,interface{}:",varinterface)
	}else {
		fmt.Println("interface{}:",varinterface)
	}

	//map
	var varmap map[string]int
	if varmap==nil {
		fmt.Println("nil,map:",varmap)
	}else {
		fmt.Println("map:",varmap)
	}

	//slice
	var varslice []int
	if  varslice==nil{
		fmt.Println("nil,slice:",varslice)
	}else {
		fmt.Println("slice:",varslice)
	}

	//channel
	var varchanreceive chan int
	if varchanreceive==nil{
		fmt.Println("nil,chan:",varchanreceive)
	}else {
		fmt.Println("chan:",varchanreceive)
	}


}

func optNotInitVar()  {

	var varslice []int
	varslice2:=append(varslice,1)
	fmt.Println("varslice:",varslice)
	fmt.Println("varslice2:",varslice2)



	//interface{}
	var varinterface interface{}
	d,ok:=varinterface.(interface{})
	fmt.Println("interface{},value:",d,",ok:",ok)

	//var varmap map[string]int
	//varmap["one"]=1                       //panic: assignment to entry in nil map


	//func



	var varchan chan int
	//fmt.Println("send value to varchan")
	//varchan<-1
	//fmt.Println("send value to varchan finish")

	go func() {
		fmt.Println("send value to varchan")
		varchan<-2
		fmt.Println("send value to varchan finish")
	}()
	time.Sleep(time.Second)
	chanval:= <-varchan
	fmt.Println("........................channel after..................................",chanval)

}


