package main

import (
	"net/http"
	"fmt"
)
/**
  基本操作
 */
func main()  {

	db := inventory{"fish":23.98,"apple":5.6}
	http.ListenAndServe("127.0.0.1:8081",db);

}

type flt float32


func (f flt) String() string {
	return fmt.Sprintf("$%.2f", f);
}

type inventory map[string]flt

func (ku inventory) ServeHTTP(writer http.ResponseWriter, req *http.Request)  {
	for k,v := range ku{
		fmt.Fprintf(writer,"%s:%s \n",k,v);
	}
}

