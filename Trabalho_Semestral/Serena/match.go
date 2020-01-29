package main

import (
	"net/http"
)

func RegisterMatch(response http.ResponseWriter, request *http.Request) {

	var err error

	body := request.Body

	match := Match{}

	if err = DecodeJson(body, &match); err != nil {
		response.Write([]byte(DECODE_ERROR_MESSAGE))
		return
	}
	match.Active = true

	if err = validate.Struct(match); err != nil {
		response.Write([]byte(VALIDATE_STRUCT_ERROR_MESSAGE))
		return
	}

	if err = InsertMatch(match); err != nil {
		response.Write([]byte(SERVER_ERROR_MESSAGE))
		return
	}

	response.Write([]byte(SUCCESS_MESSAGE))

}

func SearchAllMatches(response http.ResponseWriter, request *http.Request) {

	var err error

	match := FindAllMatches()

	if err = EncodeJson(response, match); err != nil {
		response.Write([]byte(ENCODE_ERROR_MESSAGE))
		return
	}
}

func SearchMatch(response http.ResponseWriter, request *http.Request) {

	var err error

	match := request.URL.Query().Get("name")

	matchFind, err := FindMatch(match)
	if err != nil {
		response.Write([]byte(SERVER_ERROR_MESSAGE))
		return
	}

	if err = EncodeJson(response, matchFind); err != nil {
		response.Write([]byte(ENCODE_ERROR_MESSAGE))
		return
	}

}

func DeleteMatch(response http.ResponseWriter, request *http.Request) {

	var err error

	match := request.URL.Query().Get("numbermatch")

	if err = RemoveMatch(match); err != nil {
		response.Write([]byte(SERVER_ERROR_MESSAGE))
		return
	}

	response.Write([]byte(SUCCESS_MESSAGE))

}

func SearchAllDeletedMatches(response http.ResponseWriter, request *http.Request) {

	var err error

	match := FindAllDeletedMatches()

	if err = EncodeJson(response, match); err != nil {
		response.Write([]byte(SERVER_ERROR_MESSAGE))
		return
	}

}
