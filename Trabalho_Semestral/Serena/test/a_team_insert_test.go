package test

import (
	"net/http"
	"strings"
	"testing"
)

const (
	TEAM_REGISTER = "http://localhost:8080/team/register"
)

func TestRegisterTeamSuccess(t *testing.T) {

	json := `{
		"name": "Team One",
		"playerone": {
			"firstname": "Jose",
			"cpf": "055.966.891-06"
		},
		"playertwo": {
			"firstname": "Camila",
			"cpf": "014.091.691-11"
		}
	}`

	resp, err := http.Post(TEAM_REGISTER, "application/json", strings.NewReader(json))

	if err != nil {
		t.Errorf("Test error on post execution: %v", err)
	}

	response_string := DecodeTest(resp.Body)

	if response_string != SUCCESS_MESSAGE {
		t.Logf("Expected: %v  Response: %v", SUCCESS_MESSAGE, response_string)
		t.Fail()
	}
}

func TestRegisterTeamWithoutTeamName(t *testing.T) {

	json := `{
		"name": "",
		"playerone": {
			"firstname": "Jose",
			"cpf": "055.966.891-06"
		},
		"playertwo": {
			"firstname": "Camila",
			"cpf": "014.091.691-11"
		}
	}`

	resp, err := http.Post(TEAM_REGISTER, "application/json", strings.NewReader(json))

	if err != nil {
		t.Errorf("Test error on post execution: %v", err)
	}

	response_string := DecodeTest(resp.Body)

	if response_string != VALIDATE_STRUCT_ERROR_MESSAGE {
		t.Logf("Expected: %v  Response: %v", VALIDATE_STRUCT_ERROR_MESSAGE, response_string)
		t.Fail()
	}
}

func TestRegisterTeamWithoutPlayerOne(t *testing.T) {

	json := `{
		"name": "Team One",
		"playerone": {},
		"playertwo": {
			"firstname": "Camila",
			"cpf": "014.091.691-11"
		}
	}`

	resp, err := http.Post(TEAM_REGISTER, "application/json", strings.NewReader(json))

	if err != nil {
		t.Errorf("Test error on post execution: %v", err)
	}

	response_string := DecodeTest(resp.Body)

	if response_string != VALIDATE_STRUCT_ERROR_MESSAGE {
		t.Logf("Expected: %v  Response: %v", VALIDATE_STRUCT_ERROR_MESSAGE, response_string)
		t.Fail()
	}
}

func TestRegisterWithoutPlayerTwo(t *testing.T) {

	json := `{
		"name": "Team One",
		"playerone": {
			"firstname": "Jose",
			"cpf": "055.966.891-06"
		},
		"playertwo": {}
	}`

	resp, err := http.Post(TEAM_REGISTER, "application/json", strings.NewReader(json))

	if err != nil {
		t.Errorf("Test error on post execution: %v", err)
	}

	response_string := DecodeTest(resp.Body)

	if response_string != VALIDATE_STRUCT_ERROR_MESSAGE {
		t.Logf("Expected: %v  Response: %v", VALIDATE_STRUCT_ERROR_MESSAGE, response_string)
		t.Fail()
	}
}
