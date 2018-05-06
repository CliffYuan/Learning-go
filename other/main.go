package main

import "fmt"

//只有这个init执行了，init.go中的未执行
func init()  {
	fmt.Println("pkg main init 1...")
}

func main()  {
	fmt.Println("声明变量不使用")
}
