package main

import ("net/http"
		"github.com/gorilla/mux")


func main() {

	StartServer()
}

func StartServer() {

	r := mux.NewRouter()

	r.HandleFunc("/users", RegisterUser).Methods("POST")

	r.HandleFunc("/users", SearchUser).Methods("GET")

	http.ListenAndServe(SERVER_HOST, r)
}