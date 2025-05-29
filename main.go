package main

import (
	"net/http"

	server "github.com/anicse37/RestFullAPI/Server"
)

func main() {

	http.ListenAndServe(":8080", http.HandlerFunc(server.GetUserData))

}
