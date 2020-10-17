package main

import (
	"log"
	"net/http"
)

func main()  {
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(homePage))
	log.Fatal(http.ListenAndServe("localhost:8100", mux))
}

func homePage(write http.ResponseWriter,request *http.Request)  {
	s := "<h1>Welcome to yx.com</h1>"
	write.Write([]byte(s))
}