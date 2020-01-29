package test

import (
	"net/http"
	"strings"
	"testing"
)

const (
	PLAYER_REGISTER = "http://localhost:8080/players/register"
)

func TestRegisterPlayerSuccess(t *testing.T) {

	json := `{
		"firstname": "Jose",
		"cpf": "055.966.891-06",
		"birthdate": "18-12-1995",
		"phone": "992838470",
		"email": "mateus.18.santos@sempreceub.com",
		"pass": "123456",
		"group": "1",
		"class": "1ª",
		"category": "A"
	}`

	resp, err := http.Post(PLAYER_REGISTER, "application/json", strings.NewReader(json))

	if err != nil {
		t.Errorf("Test error on register execution: %v", err)
	}

	response_string := DecodeTest(resp.Body)

	if response_string != SUCCESS_MESSAGE {
		t.Logf("Expected: %v  Response: %v", SUCCESS_MESSAGE, response_string)
		t.Fail()
	}
}

func TestRegisterPlayerWithoutName(t *testing.T) {

	json := `{
		"firstname": "",
		"cpf": "055.966.891-06",
		"birthdate": "18-12-1995",
		"phone": "992838470",
		"email": "mateus.18.santos@sempreceub.com",
		"pass": "123456",
		"group": "1",
		"class": "1ª",
		"category": "A"
	}`

	resp, err := http.Post(PLAYER_REGISTER, "application/json", strings.NewReader(json))

	if err != nil {
		t.Errorf("Test error on register execution: %v", err)
	}

	response_string := DecodeTest(resp.Body)

	if response_string != VALIDATE_STRUCT_ERROR_MESSAGE {
		t.Logf("Expected: %v  Response: %v", VALIDATE_STRUCT_ERROR_MESSAGE, response_string)
		t.Fail()
	}
}

func TestRegisterPlayerWithoutCPF(t *testing.T) {

	json := `{
		"firstname": "Jose",
		"cpf": "",
		"birthdate": "18-12-1995",
		"phone": "992838470",
		"email": "mateus.18.santos@sempreceub.com",
		"pass": "123456",
		"group": "1",
		"class": "1ª",
		"category": "A"
	}`

	resp, err := http.Post(PLAYER_REGISTER, "application/json", strings.NewReader(json))

	if err != nil {
		t.Errorf("Test error on register execution: %v", err)
	}

	response_string := DecodeTest(resp.Body)

	if response_string != VALIDATE_STRUCT_ERROR_MESSAGE {
		t.Logf("Expected: %v  Response: %v", VALIDATE_STRUCT_ERROR_MESSAGE, response_string)
		t.Fail()
	}
}

func TestRegisterPlayerWithInvalidCPF(t *testing.T) {

	json := `{
		"firstname": "Jose",
		"cpf": "111.222.333-44",
		"birthdate": "18-12-1995",
		"phone": "992838470",
		"email": "mateus.18.santos@sempreceub.com",
		"pass": "123456",
		"group": "1",
		"class": "1ª",
		"category": "A"
	}`

	resp, err := http.Post(PLAYER_REGISTER, "application/json", strings.NewReader(json))

	if err != nil {
		t.Errorf("Test error on register execution: %v", err)
	}

	response_string := DecodeTest(resp.Body)

	if response_string != VALIDATE_STRUCT_ERROR_MESSAGE {
		t.Logf("Expected: %v  Response: %v", VALIDATE_STRUCT_ERROR_MESSAGE, response_string)
		t.Fail()
	}
}

func TestRegisterPlayerWithNumberInName(t *testing.T) {

	json := `{
		"firstname": "Jose454",
		"cpf": "055.966.891-06",
		"birthdate": "18-12-1995",
		"phone": "992838470",
		"email": "mateus.18.santos@sempreceub.com",
		"pass": "123456",
		"group": "1",
		"class": "1ª",
		"category": "A"
	}`

	resp, err := http.Post(PLAYER_REGISTER, "application/json", strings.NewReader(json))

	if err != nil {
		t.Errorf("Test error on register execution: %v", err)
	}

	response_string := DecodeTest(resp.Body)

	if response_string != VALIDATE_STRUCT_ERROR_MESSAGE {
		t.Logf("Expected: %v  Response: %v", VALIDATE_STRUCT_ERROR_MESSAGE, response_string)
		t.Fail()
	}
}

func TestRegisterPlayerWithoutBirthdate(t *testing.T) {

	json := `{
		"firstname": "Jose",
		"cpf": "055.966.891-06",
		"birthdate": "",
		"phone": "992838470",
		"email": "mateus.18.santos@sempreceub.com",
		"pass": "123456",
		"group": "1",
		"class": "1ª",
		"category": "A"
	}`

	resp, err := http.Post(PLAYER_REGISTER, "application/json", strings.NewReader(json))

	if err != nil {
		t.Errorf("Test error on register execution: %v", err)
	}

	response_string := DecodeTest(resp.Body)

	if response_string != VALIDATE_STRUCT_ERROR_MESSAGE {
		t.Logf("Expected: %v  Response: %v", VALIDATE_STRUCT_ERROR_MESSAGE, response_string)
		t.Fail()
	}
}

func TestRegisterPlayerWithoutPhone(t *testing.T) {

	json := `{
		"firstname": "Jose",
		"cpf": "055.966.891-06",
		"birthdate": "18-12-1995",
		"phone": "",
		"email": "mateus.18.santos@sempreceub.com",
		"pass": "123456",
		"group": "1",
		"class": "1ª",
		"category": "A"
	}`

	resp, err := http.Post(PLAYER_REGISTER, "application/json", strings.NewReader(json))

	if err != nil {
		t.Errorf("Test error on register execution: %v", err)
	}

	response_string := DecodeTest(resp.Body)

	if response_string != VALIDATE_STRUCT_ERROR_MESSAGE {
		t.Logf("Expected: %v  Response: %v", VALIDATE_STRUCT_ERROR_MESSAGE, response_string)
		t.Fail()
	}
}

func TestRegisterPlayerWithoutEmail(t *testing.T) {

	json := `{
		"firstname": "Jose",
		"cpf": "055.966.891-06",
		"birthdate": "18-12-1995",
		"phone": "992838470",
		"email": "",
		"pass": "123456",
		"group": "1",
		"class": "1ª",
		"category": "A"
	}`

	resp, err := http.Post(PLAYER_REGISTER, "application/json", strings.NewReader(json))

	if err != nil {
		t.Errorf("Test error on register execution: %v", err)
	}

	response_string := DecodeTest(resp.Body)

	if response_string != VALIDATE_STRUCT_ERROR_MESSAGE {
		t.Logf("Expected: %v  Response: %v", VALIDATE_STRUCT_ERROR_MESSAGE, response_string)
		t.Fail()
	}
}

func TestRegisterPlayerWithoutPassword(t *testing.T) {

	json := `{
		"firstname": "Jose",
		"cpf": "055.966.891-06",
		"birthdate": "18-12-1995",
		"phone": "992838470",
		"email": "mateus.18.santos@sempreceub.com",
		"pass": "",
		"group": "1",
		"class": "1ª",
		"category": "A"
	}`

	resp, err := http.Post(PLAYER_REGISTER, "application/json", strings.NewReader(json))

	if err != nil {
		t.Errorf("Test error on register execution: %v", err)
	}

	response_string := DecodeTest(resp.Body)

	if response_string != VALIDATE_STRUCT_ERROR_MESSAGE {
		t.Logf("Expected: %v  Response: %v", VALIDATE_STRUCT_ERROR_MESSAGE, response_string)
		t.Fail()
	}
}

func TestRegisterPlayerWithoutGroup(t *testing.T) {

	json := `{
		"firstname": "Jose",
		"cpf": "055.966.891-06",
		"birthdate": "18-12-1995",
		"phone": "992838470",
		"email": "mateus.18.santos@sempreceub.com",
		"pass": "123456",
		"group": "",
		"class": "1ª",
		"category": "A"
	}`

	resp, err := http.Post(PLAYER_REGISTER, "application/json", strings.NewReader(json))

	if err != nil {
		t.Errorf("Test error on register execution: %v", err)
	}

	response_string := DecodeTest(resp.Body)

	if response_string != VALIDATE_STRUCT_ERROR_MESSAGE {
		t.Logf("Expected: %v  Response: %v", VALIDATE_STRUCT_ERROR_MESSAGE, response_string)
		t.Fail()
	}
}

func TestRegisterPlayerWithoutClass(t *testing.T) {

	json := `{
		"firstname": "Jose",
		"cpf": "055.966.891-06",
		"birthdate": "18-12-1995",
		"phone": "992838470",
		"email": "mateus.18.santos@sempreceub.com",
		"pass": "123456",
		"group": "1",
		"class": "",
		"category": "A"
	}`

	resp, err := http.Post(PLAYER_REGISTER, "application/json", strings.NewReader(json))

	if err != nil {
		t.Errorf("Test error on register execution: %v", err)
	}

	response_string := DecodeTest(resp.Body)

	if response_string != VALIDATE_STRUCT_ERROR_MESSAGE {
		t.Logf("Expected: %v  Response: %v", VALIDATE_STRUCT_ERROR_MESSAGE, response_string)
		t.Fail()
	}
}

func TestRegisterPlayerWithoutCategory(t *testing.T) {

	json := `{
		"firstname": "Jose",
		"cpf": "055.966.891-06",
		"birthdate": "18-12-1995",
		"phone": "992838470",
		"email": "mateus.18.santos@sempreceub.com",
		"pass": "123456",
		"group": "1",
		"class": "1ª",
		"category": ""
	}`

	resp, err := http.Post(PLAYER_REGISTER, "application/json", strings.NewReader(json))

	if err != nil {
		t.Errorf("Test error on register execution: %v", err)
	}

	response_string := DecodeTest(resp.Body)

	if response_string != VALIDATE_STRUCT_ERROR_MESSAGE {
		t.Logf("Expected: %v  Response: %v", VALIDATE_STRUCT_ERROR_MESSAGE, response_string)
		t.Fail()
	}
}

func TestRegisterPlayerWithoutPhoneAndBirth(t *testing.T) {

	json := `{
		"firstname": "Jose",
		"cpf": "055.966.891-06",
		"birthdate": "",
		"phone": "",
		"email": "mateus.18.santos@sempreceub.com",
		"pass": "123456",
		"group": "1",
		"class": "1ª",
		"category": "A"
	}`

	resp, err := http.Post(PLAYER_REGISTER, "application/json", strings.NewReader(json))

	if err != nil {
		t.Errorf("Test error on register execution: %v", err)
	}

	response_string := DecodeTest(resp.Body)

	if response_string != VALIDATE_STRUCT_ERROR_MESSAGE {
		t.Logf("Expected: %v  Response: %v", VALIDATE_STRUCT_ERROR_MESSAGE, response_string)
		t.Fail()
	}
}

func TestRegisterPlayerWithoutNameAndCPF(t *testing.T) {

	json := `{
		"firstname": "",
		"cpf": "",
		"birthdate": "18-12-1995",
		"phone": "992838470",
		"email": "mateus.18.santos@sempreceub.com",
		"pass": "123456",
		"group": "1",
		"class": "1ª",
		"category": "A"
	}`

	resp, err := http.Post(PLAYER_REGISTER, "application/json", strings.NewReader(json))

	if err != nil {
		t.Errorf("Test error on register execution: %v", err)
	}

	response_string := DecodeTest(resp.Body)

	if response_string != VALIDATE_STRUCT_ERROR_MESSAGE {
		t.Logf("Expected: %v  Response: %v", VALIDATE_STRUCT_ERROR_MESSAGE, response_string)
		t.Fail()
	}
}
