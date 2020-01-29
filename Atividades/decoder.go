package main

import ("fmt"
		"io/ioutil"
		"net/http"
		"encoding/json"
		"log"
		"github.com/gorilla/mux")

func main() {

	fmt.Println("Starting server...")
	
	r := mux.NewRouter()

	r.HandleFunc("/users", CreateUser).Methods("POST")

	http.ListenAndServe("localhost:8080", r)
}

type User struct {
	Name  string   "json:name"
	Age   int      "json:age"
	Phone []string "json:phones"
	Valid bool     "json:valid"
}

func CreateUser(response http.ResponseWriter, request *http.Request){

	user := User{}

	DecodeJson(request.Body, user)

}

func DecodeJson(body io.Reader, entity interface{}){

	decodificador := json.NewDecoder(body)

	for {
		if err := decodificador.Decode(&entity); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Request: %v", entity)
	}

}