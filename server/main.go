package main

import (
	"fmt"
	"log"
	"net/http"
)

func main()  {
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(homePage))
	fmt.Println("begin server!")
	log.Fatal(http.ListenAndServe("47.115.134.61:80", mux))
}

func homePage(write http.ResponseWriter,request *http.Request)  {
	s := "<h1>Welcome to yx.com</h1>"
	write.Write([]byte(s))
}