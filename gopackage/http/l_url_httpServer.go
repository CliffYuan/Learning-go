package main

import (
	"fmt"
	"net/http"
)

type flt float32

type inventoryUrl map[string]flt

func (in inventoryUrl) ServeHTTP(w http.ResponseWriter, req *http.Request)()  {
	switch req.URL.Path{
	case "/name":
		fmt.Fprint(w,"xnd");
		break;
	case "/list":
		for k,v := range in{
			fmt.Fprintf(w,"%s:%s",k,v)
		}
		break;
	case "/query":
		item :=req.URL.Query().Get("item");
		k,ok:= in[item]
		if(!ok){
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w,"not fund query %s item",item)
			break
		}
		fmt.Fprintf(w,"%s -> %s",item,k)
		break
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w,"not fund %s",req.URL.Path);
		break
	}
}

func main()  {

	db := inventoryUrl{"apple":5.6,"fish":7.1}

	http.ListenAndServe("127.0.0.1:8082",db)
}
