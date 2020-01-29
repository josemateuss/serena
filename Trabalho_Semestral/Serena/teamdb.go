package main

import (
	"log"

	"github.com/globalsign/mgo/bson"
)

func InsertTeam(team Team) (err error) {

	conn, err := GetConnection()

	if err != nil {
		log.Printf(SERVER_ERROR_MESSAGE)
		return
	}

	collection := conn.DB(DATABASE_NAME).C(TEAMS_COLLECTION)

	if err = collection.Insert(team); err != nil {
		log.Printf(SERVER_ERROR_MESSAGE)
		return
	}

	return
}

func FindAllTeams() (team []interface{}) {

	conn, err := GetConnection()

	if err != nil {
		log.Printf(SERVER_ERROR_MESSAGE)
		return
	}

	collection := conn.DB(DATABASE_NAME).C(TEAMS_COLLECTION)

	if err = collection.Find(bson.M{"active": true}).All(&team); err != nil {
		log.Printf(SERVER_ERROR_MESSAGE)
		return
	}

	return
}

func FindTeam(name string) (team Team, err error) {
	conn, err := GetConnection()
	collection := conn.DB(DATABASE_NAME).C(TEAMS_COLLECTION)

	query := bson.M{
		"name":   name,
		"active": true,
	}

	if err = collection.Find(query).One(&team); err != nil {
		log.Printf("%v", err)
	}

	return
}

func UpdateTeam(team Team) (err error) {

	conn, err := GetConnection()

	if err != nil {
		log.Printf(SERVER_ERROR_MESSAGE)
		return
	}

	collection := conn.DB(DATABASE_NAME).C(TEAMS_COLLECTION)

	if err = collection.Update(bson.M{"$set": bson.M{"name": team.Name}}, team); err != nil {
		log.Printf(SERVER_ERROR_MESSAGE)
		return
	}

	return
}

func RemoveTeam(Name string) (err error) {

	conn, err := GetConnection()

	if err != nil {
		log.Printf(SERVER_ERROR_MESSAGE)
		return
	}

	collection := conn.DB(DATABASE_NAME).C(TEAMS_COLLECTION)

	if err = collection.Update(bson.M{"name": Name}, bson.M{"$set": bson.M{"active": false}}); err != nil {
		log.Printf(SERVER_ERROR_MESSAGE)
		return
	}

	return
}
