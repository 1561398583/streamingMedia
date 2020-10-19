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
	log.Fatal(http.ListenAndServe(":8001", mux))
}

func homePage(write http.ResponseWriter,request *http.Request)  {
	s := "<h1>deploy</h1>"
	write.Write([]byte(s))
}