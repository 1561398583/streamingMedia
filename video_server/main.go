package main

import (
	_ "gorm.io/gorm"
	"log"
	"net/http"
)

func main()  {
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

