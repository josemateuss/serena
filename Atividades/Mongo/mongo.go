package main

import ("fmt"
		"github.com/globalsign/mgo"
		"github.com/globalsign/mgo/bson")

type User struct {
	Nome  string
	Idade string
	Email string
	Senha string
}

func main() {
	busca := User{}
	session, err := mgo.Dial("localhost")
	c := session.DB("banco").C("users")
	defer session.Close()

	err = c.Insert(&User{Nome: "Jose", Idade: "50", Email: "jose@email.com", Senha: "123456"})
	if err != nil {
		panic(err)
	}

	err = c.Find(bson.M{"Nome": "Jose"}).One(&busca)
	if err != nil {
		panic(err)
	}

	fmt.Println(busca)
}