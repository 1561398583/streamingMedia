package api

import "net/http"

func init()  {
	http.HandleFunc("", user1)
}
