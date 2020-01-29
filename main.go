package main

import (
		"fmt"
		"time"
		"net/http"
		"strconv"
		"github.com/gorilla/mux"
		)

func main() {

	fmt.Println("Starting server...")
	
	r := mux.NewRouter()

	r.HandleFunc("/time/", Time).Methods("GET")
	r.HandleFunc("/time/now/", TimeNow).Methods("GET")

	http.ListenAndServe("localhost:8080", r)
}

func Time(response http.ResponseWriter, request *http.Request){
	
	now := time.Now()
	response.Write([]byte("A hora do servidor é: " + now.Format("02/01/2006 15:04:05")))

}

func TimeNow(response http.ResponseWriter, request *http.Request){

	gmt := request.URL.Query()["gmt"][0]
	gmtNumber,_:=strconv.Atoi(gmt)
	locationTime := time.FixedZone("", gmtNumber*60*60)
	hora := time.Now().In(locationTime).Format("02/01/2006 15:04:05")
	response.Write([]byte("A hora do servidor é: " + hora))
}