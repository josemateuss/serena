package main

import (
	"encoding/json"
	"net/http"

	"github.com/globalsign/mgo/bson"
)

type User struct {
	Name      string `json:"name"`
	Pass      string `json:"pass"`
	Birthdate string `json:"birthdate"`
	Gender    string `json:"gender"`
}

func RegisterUser(response http.ResponseWriter, request *http.Request) {

	body := request.Body

	user := User{}

	DecodeJson(body, &user)

	err := SaveUser(user)

	if err != nil {
		response.Write([]byte("400"))
		return
	}

	response.Write([]byte("200"))

}

func SaveUser(user User) (err error) {

	conn, err := GetConnection()

	collection := conn.DB("reprova").C("users")

	collection.Insert(user)

	return
}

func SearchUser(response http.ResponseWriter, request *http.Request) {
	nameUser := request.URL.Query().Get("name")
	genderUser := request.URL.Query().Get("gender")

	var user User

	user.Name = nameUser
	user.Gender = genderUser

	userFind, err := FindUser(user)
	if err != nil {
		response.Write([]byte("400"))
		return
	}

	respUser, _ := json.Marshal(userFind)

	response.Write(respUser)
	return

}

func FindUser(user User) (err error) {

	conn, err := GetConnection()

	collection := conn.DB("reprova").C("users")

	queryUser := bson.M{"$or": []bson.M{bson.M{"name": name}, bson.M{"gender": gender}}}
	err = collection.Find(queryUser).All(&user)

	return
}
