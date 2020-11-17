package api

import (
	"net/http"
	"video_server/session"
)


func init()  {
	http.HandleFunc("/user_login", user_login)
}

func user_login(w http.ResponseWriter,r *http.Request)  {
	
}

