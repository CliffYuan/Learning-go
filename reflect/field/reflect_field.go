package main

import (
	"reflect"
	"github.com/golang/go/src/pkg/fmt"
)

type NotKnownType struct {
	s1, s2, s3 string
	k int32
}

func (n NotKnownType) String() string {
	return n.s1+"-"+n.s2+"-"+n.s3
}

func main()  {

	var secret interface{} =NotKnownType{"Ada","Go","Oberon",32}

	value:=reflect.ValueOf(secret)
	typ:=reflect.TypeOf(secret)

	fmt.Println(typ);
	fmt.Println(value);
	fmt.Println(typ.Kind())
	fmt.Println(value.Kind())

	fmt.Println("------------------show field---------------------")
	for i:=0;i<value.NumField() ;i++  {
		fmt.Printf("Field %d,name:s% = %v\n",i,typ.Field(i).Name, value.Field(i))
	}

	fmt.Println("------------------method field---------------------")
	for i:=0;i<value.NumMethod() ;i++  {
		fmt.Printf("Method %d:%v\n",i,value.Method(i))
		result:=value.Method(i).Call(nil);
		fmt.Println(result)
	}


}
