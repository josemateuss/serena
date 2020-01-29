package main

import (
	"log"

	"github.com/globalsign/mgo/bson"
)

func InsertPlayer(player Player) (err error) {

	conn, err := GetConnection()

	if err != nil {
		log.Printf(SERVER_ERROR_MESSAGE)
		return
	}

	collection := conn.DB(DATABASE_NAME).C(PLAYERS_COLLECTION)

	if err = collection.Insert(player); err != nil {
		log.Printf(SERVER_ERROR_MESSAGE)
		return
	}

	return
}

func FindAllPlayers() (player []interface{}) {

	conn, err := GetConnection()

	if err != nil {
		log.Printf(SERVER_ERROR_MESSAGE)
		return
	}

	collection := conn.DB(DATABASE_NAME).C(PLAYERS_COLLECTION)

	if err = collection.Find(bson.M{"active": true}).All(&player); err != nil {
		log.Printf(SERVER_ERROR_MESSAGE)
		return
	}

	return
}

func FindPlayer(cpf string) (player Player, err error) {
	conn, err := GetConnection()
	collection := conn.DB(DATABASE_NAME).C(PLAYERS_COLLECTION)

	query := bson.M{
		"cpf":    cpf,
		"active": true,
	}

	if err = collection.Find(query).One(&player); err != nil {
		log.Printf("%v", err)
	}
	/* 	result = []Player{}

	   	if err != nil {
	   		log.Printf("[ERROR] could not get BD connection")
	   		return
	   	}


	   	var list []bson.M
	   	list = append(list, bson.M{"cpf": cpf})
	   	list = append(list, bson.M{"active": true})

	   	query := bson.M{"$and": list}



	   	if err != nil {
	   		log.Printf(`{"code": 400, "message": "error: body content"}`, err)
	   		return
	   	}

		   return */

	return
}

func UpdatePlayer(player PlayerUpdate) (err error) {

	conn, err := GetConnection()

	if err != nil {
		log.Printf(SERVER_ERROR_MESSAGE)
		return
	}

	collection := conn.DB(DATABASE_NAME).C(PLAYERS_COLLECTION)

	if err = collection.Update(bson.M{"$set": bson.M{"firstname": player.FirstName}}, player); err != nil {
		log.Printf(SERVER_ERROR_MESSAGE)
		return
	}

	return
}

func RemovePlayer(CPF string) (err error) {

	conn, err := GetConnection()

	if err != nil {
		log.Printf(SERVER_ERROR_MESSAGE)
		return
	}

	collection := conn.DB(DATABASE_NAME).C(PLAYERS_COLLECTION)

	if err = collection.Update(bson.M{"cpf": CPF}, bson.M{"$set": bson.M{"active": false}}); err != nil {
		log.Printf(SERVER_ERROR_MESSAGE)
		return
	}

	return
}

func FindAllDeletedPlayers() (player []interface{}) {

	conn, err := GetConnection()

	if err != nil {
		log.Printf(SERVER_ERROR_MESSAGE)
		return
	}

	collection := conn.DB(DATABASE_NAME).C(PLAYERS_COLLECTION)

	if err = collection.Find(bson.M{"active": false}).All(&player); err != nil {
		log.Printf(SERVER_ERROR_MESSAGE)
		return
	}

	return
}

func ActivePlayer(CPF string) (err error) {

	conn, err := GetConnection()

	if err != nil {
		log.Printf(SERVER_ERROR_MESSAGE)
		return
	}

	collection := conn.DB(DATABASE_NAME).C(PLAYERS_COLLECTION)

	if err = collection.Update(bson.M{"cpf": CPF}, bson.M{"$set": bson.M{"active": true}}); err != nil {
		log.Printf(SERVER_ERROR_MESSAGE)
		return
	}

	return
}

func FindPlayerInterface(cpf string, player interface{}) error {
	var err error

	conn, err := GetConnection()
	collection := conn.DB(DATABASE_NAME).C(PLAYERS_COLLECTION)

	query := bson.M{
		"cpf":    cpf,
		"active": true,
	}

	if err = collection.Find(query).One(player); err != nil {
		log.Printf("%v", err)
	}

	return err
}

func FindPlayerBool(cpf string) bool {
	var err error

	conn, err := GetConnection()
	collection := conn.DB(DATABASE_NAME).C(PLAYERS_COLLECTION)

	query := bson.M{
		"cpf":    cpf,
		"active": true,
	}

	result := Player{}

	if err = collection.Find(query).One(&result); err != nil {
		return false
	}

	return true

}

func FindPlayerLogin(query, result interface{}) (err error) {

	conn, err := GetConnection()

	collection := conn.DB(DATABASE_NAME).C(PLAYERS_COLLECTION)

	err = collection.Find(query).All(result)

	return
}
