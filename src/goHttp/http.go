package main

import (
	"fmt"
	"io"
	"net/http"
	"goLib/hello"
)

func RHello(w http.ResponseWriter, r *http.Request){
	io.WriteString(w, "Hello world!")
}

func main(){
	http.HandleFunc("/",RHello)
	hello.Hello()
	fmt.Printf("listening on port 8000\n")
	http.ListenAndServe(":8000",nil)
}
