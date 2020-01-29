package main

import (
	"net/http"

	"github.com/globalsign/mgo/bson"
)

func RegisterPlayer(response http.ResponseWriter, request *http.Request) {

	var err error

	body := request.Body

	player := Player{}

	if err = DecodeJson(body, &player); err != nil {
		response.Write([]byte(DECODE_ERROR_MESSAGE))
		return
	}
	player.Active = true

	if err = validate.Struct(player); err != nil {
		response.Write([]byte(VALIDATE_STRUCT_ERROR_MESSAGE))
		return
	}

	if err = InsertPlayer(player); err != nil {
		response.Write([]byte(SERVER_ERROR_MESSAGE))
		return
	}

	response.Write([]byte(SUCCESS_MESSAGE))
}

func SearchAllPlayers(response http.ResponseWriter, request *http.Request) {

	var err error

	player := FindAllPlayers()

	if err = EncodeJson(response, player); err != nil {
		response.Write([]byte(ENCODE_ERROR_MESSAGE))
		return
	}

}

func SearchPlayer(response http.ResponseWriter, request *http.Request) {

	var err error

	player := request.URL.Query().Get("cpf")

	playerFind, err := FindPlayer(player)
	if err != nil {
		response.Write([]byte(SERVER_ERROR_MESSAGE))
		return
	}

	if err = EncodeJson(response, playerFind); err != nil {
		response.Write([]byte(ENCODE_ERROR_MESSAGE))
		return
	}

}

func EditPlayer(response http.ResponseWriter, request *http.Request) {

	var err error

	body := request.Body

	player := PlayerUpdate{}

	if err = DecodeJson(body, &player); err != nil {
		response.Write([]byte(DECODE_ERROR_MESSAGE))
		return
	}

	player.Active = true

	if err = validate.Struct(player); err != nil {
		response.Write([]byte(VALIDATE_STRUCT_ERROR_MESSAGE))
		return
	}

	if err = UpdatePlayer(player); err != nil {
		response.Write([]byte(SERVER_ERROR_MESSAGE))
		return
	}

	response.Write([]byte(SUCCESS_MESSAGE))

}

func DeletePlayer(response http.ResponseWriter, request *http.Request) {

	var err error

	player := request.URL.Query().Get("cpf")

	if err = RemovePlayer(player); err != nil {
		response.Write([]byte(SERVER_ERROR_MESSAGE))
		return
	}

	response.Write([]byte(SUCCESS_MESSAGE))

}

func SearchAllDeletedPlayers(response http.ResponseWriter, request *http.Request) {

	var err error

	player := FindAllDeletedPlayers()

	if err = EncodeJson(response, player); err != nil {
		response.Write([]byte(ENCODE_ERROR_MESSAGE))
		return
	}

}

func ActivatePlayer(response http.ResponseWriter, request *http.Request) {

	var err error

	playerCPF := request.URL.Query().Get("cpf")

	if err = RemovePlayer(playerCPF); err != nil {
		response.Write([]byte(SERVER_ERROR_MESSAGE))
		return
	}

	response.Write([]byte(SUCCESS_MESSAGE))

}

func RegisterPlayersInChampionship(response http.ResponseWriter, request *http.Request) {

	var err error

	body := request.Body

	champ := Championship{}

	if err = DecodeJson(body, &champ); err != nil {
		response.Write([]byte(DECODE_ERROR_MESSAGE))
		return
	}

	for _, player := range champ.Players {
		if ok := FindPlayerBool(player.CPF); ok == false {
			response.Write([]byte(NOT_FOUND_ERROR_MESSAGE))

			return
		}
	}

	for _, player := range champ.Players {
		tmp := Player{}
		FindPlayerInterface(player.CPF, &tmp)
		if err := InsertPlayer(tmp); err != nil {
			response.Write([]byte(SERVER_ERROR_MESSAGE))

			return
		}
	}
}

func LoginPlayer(response http.ResponseWriter, request *http.Request) {

	var err error

	body := request.Body

	player := PlayerLogin{}

	if err = DecodeJson(body, &player); err != nil {
		response.Write([]byte(DECODE_ERROR_MESSAGE))

		return
	}

	if err = validate.Struct(player); err != nil {
		response.Write([]byte(VALIDATE_STRUCT_ERROR_MESSAGE))

		return
	}

	players := []Player{}

	query := bson.M{"email": player.Email, "pass": player.Password}
	if FindPlayerLogin(query, &players); len(players) == 0 {

		return
	}

	response.Write([]byte(SUCCESS_MESSAGE))

}
