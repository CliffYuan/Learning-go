package main

import (
	"flag"
	"fmt"
)

func main()  {
	//默认值
	var flag_i=flag.Int("xnd",1,"flag int");
	fmt.Println(flag_i);
	fmt.Println(*flag_i);

	//output:
	//0xc42001c110
	//1


	var flag_int=flag.Int("int",2,"flag int 2");
	fmt.Println(flag_int);
	fmt.Println(&flag_int);
	fmt.Println(*flag_int);

	//output:
	//0xc42001c130
	//0xc42000c030
	//2


	//output:
}
