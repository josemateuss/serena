package test

import (
	"net/http"
	"testing"
)

const (
	CHAMPIONSHIP_FIND = "http://localhost:8080/championships/search"
)

func TestGetChampionshipSuccess(t *testing.T) {
	resp, err := http.Get(CHAMPIONSHIP_FIND + "?name=Champ+1")

	if err != nil {
		t.Errorf("Test error on get execution: %v", err)
	}

	response_string := DecodeTest(resp.Body)

	if response_string != SUCCESS_MESSAGE {
		t.Logf("Expected: %v. Response: %v", SUCCESS_MESSAGE, response_string)
		t.Fail()
	}
}
