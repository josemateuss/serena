package main

import (
	"log"

	"github.com/globalsign/mgo/bson"
)

func InsertChampionship(championship Championship) (err error) {

	conn, err := GetConnection()

	if err != nil {
		log.Printf(SERVER_ERROR_MESSAGE)
		return
	}

	collection := conn.DB(DATABASE_NAME).C(CHAMPIONSHIPS_COLLECTION)

	if err = collection.Insert(championship); err != nil {
		log.Printf(SERVER_ERROR_MESSAGE)
		return
	}

	return
}

func FindAllChampionships() (championship []interface{}) {

	conn, err := GetConnection()

	if err != nil {
		log.Printf(SERVER_ERROR_MESSAGE)
		return
	}

	collection := conn.DB(DATABASE_NAME).C(CHAMPIONSHIPS_COLLECTION)

	if err = collection.Find(bson.M{"active": true}).All(&championship); err != nil {
		log.Printf(SERVER_ERROR_MESSAGE)
		return
	}

	return
}

func FindChampionship(name string) (result []Championship, err error) {

	result = []Championship{}

	conn, err := GetConnection()

	if err != nil {
		log.Printf(SERVER_ERROR_MESSAGE)
		return
	}

	collection := conn.DB(DATABASE_NAME).C(CHAMPIONSHIPS_COLLECTION)

	var list []bson.M
	list = append(list, bson.M{"name": name})
	list = append(list, bson.M{"active": true})

	query := bson.M{"$and": list}

	err = collection.Find(query).All(&result)

	if err != nil {
		log.Printf("%v", err)
		return
	}

	return
}

func UpdateChampionship(championship ChampionshipUpdate) (err error) {

	conn, err := GetConnection()

	if err != nil {
		log.Printf(SERVER_ERROR_MESSAGE)
		return
	}

	collection := conn.DB(DATABASE_NAME).C(CHAMPIONSHIPS_COLLECTION)

	if err = collection.Update(bson.M{"$set": bson.M{"name": championship.Name}}, championship); err != nil {
		log.Printf(SERVER_ERROR_MESSAGE)
		return
	}

	return
}

func RemoveChampionship(Name string) (err error) {

	conn, err := GetConnection()

	if err != nil {
		log.Printf(SERVER_ERROR_MESSAGE)
		return
	}

	collection := conn.DB(DATABASE_NAME).C(CHAMPIONSHIPS_COLLECTION)

	if err = collection.Update(bson.M{"name": Name}, bson.M{"$set": bson.M{"active": false}}); err != nil {
		log.Printf(SERVER_ERROR_MESSAGE)
		return
	}

	return
}
