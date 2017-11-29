package main

import (
	"fmt"
	"errors"
)

func main()  {

	e:=errors.New("no data");

	fmt.Printf("error type is %T",e)

}
