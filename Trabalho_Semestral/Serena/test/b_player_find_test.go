package test

import (
	"net/http"
	"testing"
)

const (
	PLAYER_FIND = "http://localhost:8080/players/search"
)

func TestGetPlayerSuccess(t *testing.T) {
	resp, err := http.Get(PLAYER_FIND + "?cpf=055.966.891-06")

	if err != nil {
		t.Errorf("Test error on get execution: %v", err)
	}

	response_string := DecodeTest(resp.Body)

	if response_string != SUCCESS_MESSAGE {
		t.Logf("Expected: %v. Response: %v", SUCCESS_MESSAGE, response_string)
		t.Fail()
	}
}
