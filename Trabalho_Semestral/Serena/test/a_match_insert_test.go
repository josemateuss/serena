package test

import (
	"net/http"
	"strings"
	"testing"
)

const (
	MATCH_REGISTER = "http://localhost:8080/matches/register"
)

func TestRegisterMatchSuccess(t *testing.T) {

	json := `{
		"numbermatch": 1,
		"player1": {
			"firstname": "Jose",
			"cpf": "055.966.891-06"
		},
		"player2": {
			"firstname": "Camila",
			"cpf": "014.091.691-11"
		},
		"point1": 1,
		"game1": 2,
		"set1": 3,
		"point2": 4,
		"game2": 5,
		"set2": 6,
		"winner1": true,
		"winner2": false,
		"type": "Basic"
	}`

	resp, err := http.Post(MATCH_REGISTER, "application/json", strings.NewReader(json))

	if err != nil {
		t.Errorf("Test error on post execution: %v", err)
	}

	response_string := DecodeTest(resp.Body)

	if response_string != SUCCESS_MESSAGE {
		t.Logf("Expected: %v  Response: %v", SUCCESS_MESSAGE, response_string)
		t.Fail()
	}
}

func TestRegisterMatchWithoutPlayers(t *testing.T) {

	json := `{
		"numbermatch": 1,
		"player1": {},
		"player2": {},
		"point1": 1,
		"game1": 2,
		"set1": 3,
		"point2": 4,
		"game2": 5,
		"set2": 6,
		"winner1": true,
		"winner2": false,
		"type": "Basic"
	}`

	resp, err := http.Post(MATCH_REGISTER, "application/json", strings.NewReader(json))

	if err != nil {
		t.Errorf("Test error on post execution: %v", err)
	}

	response_string := DecodeTest(resp.Body)

	if response_string != VALIDATE_STRUCT_ERROR_MESSAGE {
		t.Logf("Expected: %v  Response: %v", VALIDATE_STRUCT_ERROR_MESSAGE, response_string)
		t.Fail()
	}
}
func TestRegisterMatchWithoutNumber(t *testing.T) {

	json := `{
		"numbermatch": null,
		"player1": {
			"firstname": "Jose",
			"cpf": "055.966.891-06"
		},
		"player2": {
			"firstname": "Camila",
			"cpf": "014.091.691-11"
		},
		"point1": 1,
		"game1": 2,
		"set1": 3,
		"point2": 4,
		"game2": 5,
		"set2": 6,
		"winner1": true,
		"winner2": false,
		"type": "Basic"
	}`

	resp, err := http.Post(MATCH_REGISTER, "application/json", strings.NewReader(json))

	if err != nil {
		t.Errorf("Test error on post execution: %v", err)
	}

	response_string := DecodeTest(resp.Body)

	if response_string != DECODE_ERROR_MESSAGE {
		t.Logf("Expected: %v  Response: %v", DECODE_ERROR_MESSAGE, response_string)
		t.Fail()
	}
}
