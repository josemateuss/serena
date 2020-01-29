package test

import (
	"net/http"
	"testing"
)

const (
	TEAM_FIND = "http://localhost:8080/team/search"
)

func TestGetTeamSuccess(t *testing.T) {

	resp, err := http.Get(TEAM_FIND + "?name=Team+One")

	if err != nil {
		t.Errorf("Test error on get execution: %v", err)
	}

	response_string := DecodeTest(resp.Body)

	if response_string != SUCCESS_MESSAGE {
		t.Logf("Expected: %v. Response: %v", SUCCESS_MESSAGE, response_string)
		t.Fail()
	}
}

func TestGetTeamNotFound(t *testing.T) {
	resp, err := http.Get(TEAM_FIND + "?name=Team+Not")

	if err != nil {
		t.Errorf("Test error on get execution: %v", err)
	}

	response_string := DecodeTest(resp.Body)

	if response_string != NOT_FOUND_ERROR_MESSAGE {
		t.Logf("Expected: %v. Response: %v", NOT_FOUND_ERROR_MESSAGE, response_string)
		t.Fail()
	}
}
