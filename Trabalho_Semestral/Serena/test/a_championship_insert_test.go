package test

import (
	"net/http"
	"strings"
	"testing"
)

const (
	CHAMPIONSHIP_REGISTER = "http://localhost:8080/championships/register"
)

func TestRegisterChampionshipSuccess(t *testing.T) {

	json := `{
		"name": "Champ1",
		"players": [{
				"firstname": "Jose",
				"cpf": "055.966.891-06"
			},
			{
				"firstname": "Camila",
				"cpf": "014.091.691-11"
			}
		],
		"idate": "26-11-2018",
		"sets": 3,
		"admin": {
			"firstname": "Jose",
			"cpf": "055.966.891-06"
		}
	}`

	resp, err := http.Post(CHAMPIONSHIP_REGISTER, "application/json", strings.NewReader(json))

	if err != nil {
		t.Errorf("Test error on post execution: %v", err)
	}

	response_string := DecodeTest(resp.Body)

	if response_string != SUCCESS_MESSAGE {
		t.Logf("Expected: %v  Response: %v", SUCCESS_MESSAGE, response_string)
		t.Fail()
	}
}

func TestRegisterChampionshipWithoutName(t *testing.T) {

	json := `{
		"name": "",
		"players": [{
				"firstname": "Jose",
				"cpf": "055.966.891-06"
			},
			{
				"firstname": "Camila",
				"cpf": "014.091.691-11"
			}
		],
		"idate": "26-11-2018",
		"sets": 3,
		"admin": {
			"firstname": "Jose",
			"cpf": "055.966.891-06"
		}
	}`

	resp, err := http.Post(CHAMPIONSHIP_REGISTER, "application/json", strings.NewReader(json))

	if err != nil {
		t.Errorf("Test error on post execution: %v", err)
	}

	response_string := DecodeTest(resp.Body)

	if response_string != VALIDATE_STRUCT_ERROR_MESSAGE {
		t.Logf("Expected: %v  Response: %v", VALIDATE_STRUCT_ERROR_MESSAGE, response_string)
		t.Fail()
	}
}

func TestRegisterChampionshipWithoutPlayers(t *testing.T) {

	json := `{
		"name": "Champ 1",
		"players": [],
		"idate": "26-11-2018",
		"sets": 3,
		"admin": {
			"firstname": "Jose",
			"cpf": "055.966.891-06"
		}
	}`

	resp, err := http.Post(CHAMPIONSHIP_REGISTER, "application/json", strings.NewReader(json))

	if err != nil {
		t.Errorf("Test error on post execution: %v", err)
	}

	response_string := DecodeTest(resp.Body)

	if response_string != VALIDATE_STRUCT_ERROR_MESSAGE {
		t.Logf("Expected: %v  Response: %v", VALIDATE_STRUCT_ERROR_MESSAGE, response_string)
		t.Fail()
	}
}

func TestRegisterChampionshipWithoutInitialDate(t *testing.T) {

	json := `{
		"name": "Champ 1",
		"players": [{
				"firstname": "Jose",
				"cpf": "055.966.891-06"
			},
			{
				"firstname": "Camila",
				"cpf": "014.091.691-11"
			}
		],
		"idate": "",
		"sets": 3,
		"admin": {
			"firstname": "Jose",
			"cpf": "055.966.891-06"
		}
	}`

	resp, err := http.Post(CHAMPIONSHIP_REGISTER, "application/json", strings.NewReader(json))

	if err != nil {
		t.Errorf("Test error on post execution: %v", err)
	}

	response_string := DecodeTest(resp.Body)

	if response_string != VALIDATE_STRUCT_ERROR_MESSAGE {
		t.Logf("Expected: %v  Response: %v", VALIDATE_STRUCT_ERROR_MESSAGE, response_string)
		t.Fail()
	}
}

func TestRegisterChampionshipWithoutAdmin(t *testing.T) {

	json := `{
		"name": "Champ 1",
		"players": [{
				"firstname": "Jose",
				"cpf": "055.966.891-06"
			},
			{
				"firstname": "Camila",
				"cpf": "014.091.691-11"
			}
		],
		"idate": "26-11-2018",
		"sets": 3,
		"admin": {}
	}`

	resp, err := http.Post(CHAMPIONSHIP_REGISTER, "application/json", strings.NewReader(json))

	if err != nil {
		t.Errorf("Test error on post execution: %v", err)
	}

	response_string := DecodeTest(resp.Body)

	if response_string != VALIDATE_STRUCT_ERROR_MESSAGE {
		t.Logf("Expected: %v  Response: %v", VALIDATE_STRUCT_ERROR_MESSAGE, response_string)
		t.Fail()
	}
}

func TestRegisterChampionshipNumberSets(t *testing.T) {

	json := `{
		"name": "Champ 1",
		"players": [{
				"firstname": "Jose",
				"cpf": "055.966.891-06"
			},
			{
				"firstname": "Camila",
				"cpf": "014.091.691-11"
			}
		],
		"idate": "26-11-2018",
		"sets": null,
		"admin": {
			"firstname": "Jose",
			"cpf": "055.966.891-06"
		}
	}`

	resp, err := http.Post(CHAMPIONSHIP_REGISTER, "application/json", strings.NewReader(json))

	if err != nil {
		t.Errorf("Test error on post execution: %v", err)
	}

	response_string := DecodeTest(resp.Body)

	if response_string != VALIDATE_STRUCT_ERROR_MESSAGE {
		t.Logf("Expected: %v  Response: %v", VALIDATE_STRUCT_ERROR_MESSAGE, response_string)
		t.Fail()
	}
}
