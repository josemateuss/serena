package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

var mdc_slice = []int{2, 4, 8, 16, 32, 64, 128}

func RegisterChampionship(response http.ResponseWriter, request *http.Request) {

	var err error

	body := request.Body

	championship := Championship{}

	if err = DecodeJson(body, &championship); err != nil {
		response.Write([]byte(DECODE_ERROR_MESSAGE))
		return
	}

	championship.Active = true

	if err = validate.Struct(championship); err != nil {
		response.Write([]byte(VALIDATE_STRUCT_ERROR_MESSAGE))
		return
	}

	if err = InsertChampionship(championship); err != nil {
		response.Write([]byte(SERVER_ERROR_MESSAGE))
		return
	}

	response.Write([]byte(SUCCESS_MESSAGE))

}

func SearchAllChampionships(response http.ResponseWriter, request *http.Request) {

	var err error

	championship := FindAllChampionships()

	if err = EncodeJson(response, championship); err != nil {
		response.Write([]byte(ENCODE_ERROR_MESSAGE))
		return
	}

}
func SearchChampionship(response http.ResponseWriter, request *http.Request) {

	var err error

	championshipName := request.URL.Query().Get("name")

	championshipFind, err := FindChampionship(championshipName)
	if err != nil {
		response.Write([]byte(SERVER_ERROR_MESSAGE))
		return
	}

	if err = EncodeJson(response, championshipFind); err != nil {
		response.Write([]byte(DECODE_ERROR_MESSAGE))
		return
	}
}

func EditChampionship(response http.ResponseWriter, request *http.Request) {

	var err error

	body := request.Body

	championship := ChampionshipUpdate{}

	if err = DecodeJson(body, &championship); err != nil {
		response.Write([]byte(DECODE_ERROR_MESSAGE))
		return
	}

	championship.Active = true

	if err = validate.Struct(championship); err != nil {
		response.Write([]byte(VALIDATE_STRUCT_ERROR_MESSAGE))
		return
	}

	if err = UpdateChampionship(championship); err != nil {
		response.Write([]byte(SERVER_ERROR_MESSAGE))
		return
	}

	response.Write([]byte(SUCCESS_MESSAGE))

}

func DeleteChampionship(response http.ResponseWriter, request *http.Request) {

	var err error

	championship := request.URL.Query().Get("name")

	if err = RemoveChampionship(championship); err != nil {
		response.Write([]byte(SERVER_ERROR_MESSAGE))
		return
	}

	response.Write([]byte(SUCCESS_MESSAGE))

}

func CreateChampTables(champ Championship) (err error) {

	size := len(champ.Players)
	if size%2 == 1 {
		err = fmt.Errorf("Champ does not have the correct number of players: %v", size)
		return
	}

	for _, number := range mdc_slice {
		if size != number {
			err = fmt.Errorf("Champ does not have the correct number of players: %v", size)
			return
		}
	}

	//first_stage_size := size / 2
	//match_arr := make([]Match, first_stage_size)
	players_slice := make([]PlayerCPF, size)
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for _, i := range r.Perm(size) {
		player := champ.Players[i]
		players_slice[i] = player
	}

	return
}

func Merge(players []PlayerCPF, first_element_index, quantity int) {
	if quantity < 1 {
		return
	}
	m := quantity / 2
	Merge(players, first_element_index, m)
	Merge(players, first_element_index+m, quantity-m)

}
