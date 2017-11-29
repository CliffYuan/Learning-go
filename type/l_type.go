package main

import (
	"fmt"
	"io/ioutil"
)

func main()  {

	test_type_convert();



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



