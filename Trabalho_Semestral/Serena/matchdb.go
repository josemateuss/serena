package main

import (
	"log"

	"github.com/globalsign/mgo/bson"
)

func InsertMatch(match Match) (err error) {

	conn, err := GetConnection()

	if err != nil {
		log.Printf(SERVER_ERROR_MESSAGE)
		return
	}

	collection := conn.DB(DATABASE_NAME).C(MATCHES_COLLECTION)

	if err = collection.Insert(match); err != nil {
		log.Printf(SERVER_ERROR_MESSAGE)
		return
	}

	return
}

func FindAllMatches() (match []interface{}) {

	conn, err := GetConnection()

	if err != nil {
		log.Printf(SERVER_ERROR_MESSAGE)
		return
	}

	collection := conn.DB(DATABASE_NAME).C(MATCHES_COLLECTION)

	if err = collection.Find(bson.M{"active": true}).All(&match); err != nil {
		log.Printf(SERVER_ERROR_MESSAGE)
		return
	}
	return
}

func FindMatch(name string) (match Match, err error) {
	conn, err := GetConnection()
	collection := conn.DB(DATABASE_NAME).C(MATCHES_COLLECTION)

	query := bson.M{
		"name":   name,
		"active": true,
	}

	if err = collection.Find(query).One(&match); err != nil {
		log.Printf("%v", err)
	}

	return
}

func RemoveMatch(Name string) (err error) {

	conn, err := GetConnection()

	if err != nil {
		log.Printf(SERVER_ERROR_MESSAGE)
		return
	}

	collection := conn.DB(DATABASE_NAME).C(MATCHES_COLLECTION)

	if err = collection.Update(bson.M{"name": Name}, bson.M{"$set": bson.M{"active": false}}); err != nil {
		log.Printf(SERVER_ERROR_MESSAGE)
		return
	}

	return
}

func FindAllDeletedMatches() (match []interface{}) {

	conn, err := GetConnection()

	if err != nil {
		log.Printf(SERVER_ERROR_MESSAGE)
		return
	}

	collection := conn.DB(DATABASE_NAME).C(MATCHES_COLLECTION)

	if err = collection.Find(bson.M{"active": false}).All(&match); err != nil {
		log.Printf(SERVER_ERROR_MESSAGE)
		return
	}

	return
}
