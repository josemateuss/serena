package main

import (
	"net/http"
)

func RegisterTeam(response http.ResponseWriter, request *http.Request) {

	var err error

	body := request.Body

	team := Team{}

	if err = DecodeJson(body, &team); err != nil {
		response.Write([]byte(DECODE_ERROR_MESSAGE))
		return
	}

	team.Active = true

	if err = validate.Struct(team); err != nil {
		response.Write([]byte(VALIDATE_STRUCT_ERROR_MESSAGE))
		return
	}

	if err = InsertTeam(team); err != nil {
		response.Write([]byte(SERVER_ERROR_MESSAGE))
		return
	}

	response.Write([]byte(SUCCESS_MESSAGE))
}

func SearchAllTeams(response http.ResponseWriter, request *http.Request) {

	var err error

	team := FindAllTeams()

	if err = EncodeJson(response, team); err != nil {
		response.Write([]byte(ENCODE_ERROR_MESSAGE))
		return
	}

}

func SearchTeam(response http.ResponseWriter, request *http.Request) {

	var err error

	team := request.URL.Query().Get("name")

	teamFind, err := FindTeam(team)
	if err != nil {
		response.Write([]byte(SERVER_ERROR_MESSAGE))
		return
	}

	if err = EncodeJson(response, teamFind); err != nil {
		response.Write([]byte(ENCODE_ERROR_MESSAGE))
		return
	}

}

func EditTeam(response http.ResponseWriter, request *http.Request) {

	var err error

	body := request.Body

	team := Team{}

	if err = DecodeJson(body, &team); err != nil {
		response.Write([]byte(DECODE_ERROR_MESSAGE))
		return
	}

	team.Active = true

	if err = validate.Struct(team); err != nil {
		response.Write([]byte(VALIDATE_STRUCT_ERROR_MESSAGE))
		return
	}

	if err = UpdateTeam(team); err != nil {
		response.Write([]byte(SERVER_ERROR_MESSAGE))
		return
	}

	response.Write([]byte(SUCCESS_MESSAGE))

}

func DeleteTeam(response http.ResponseWriter, request *http.Request) {

	var err error

	team := request.URL.Query().Get("name")

	if err = RemoveTeam(team); err != nil {
		response.Write([]byte(SERVER_ERROR_MESSAGE))
		return
	}

	response.Write([]byte(SUCCESS_MESSAGE))

}
