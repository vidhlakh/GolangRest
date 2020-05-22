package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	var router = mux.NewRouter() //create router
	const port string = ":8000"
	// create handler with 2 params 1. root path 2. Func Resp and request
	router.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(resp, "Up and Running .. .")
	})
	log.Println("Server listening on port ", port)
	log.Fatalln(http.ListenAndServe(port, router))

}
