package main

import (
	"fmt"
)

func main()  {

	test_type_convert();


	fmt.Printf("---------------------xxx-----------------------------\n")


	var v_i int;

	var v_s string;           // type is string

	v_s_ptr:=&v_s             // type is *string

	var v_inter interface{}   // type is nil

	var v_typeint typeint     // type is

	classifier(v_i,v_s,v_s_ptr,v_inter,v_typeint)

}

type typeint int

func classifier(items ...interface{})  {

	for i,x:=range items {
		switch x.(type) {
		case bool:
			fmt.Printf("Param #%d is a bool\n",i)
		case float32:
			fmt.Printf("Param #%d is a float32\n",i)
		case int:
			fmt.Printf("Param #%d is a int\n",i)
		case typeint:
			fmt.Printf("Param #%d is a typeint\n",i)
		case nil:
			fmt.Printf("Param #%d is a nil\n",i)                   //todo default interface is nil
		case string:
			fmt.Printf("Param #%d is a string\n",i)
		case *string:
			fmt.Printf("Param #%d is a *string\n",i)               //todo ptr type
		default:
			fmt.Printf("Param #%d is unknown\n",i)
		}
	}


}

/**
 测试type自定义基本类型，转化后不会创建新值
 */
func test_type_convert()  {
	type String string

	s:="string value"

	fmt.Println(&s)

	ss:=String(s)

	fmt.Println(&ss)

	//output:
	//0xc4200761b0
	//0xc4200761c0

}



