package test

import (
	"net/http"
	"testing"
)

const (
	MATCH_FIND = "http://localhost:8080/matches/search/"
)

func TestGetMatchSuccess(t *testing.T) {
	resp, err := http.Get(MATCH_FIND + "?numbermatch=1")

	if err != nil {
		t.Errorf("Test error on get execution: %v", err)
	}

	response_string := DecodeTest(resp.Body)

	if response_string != SUCCESS_MESSAGE {
		t.Logf("Expected: %v. Response: %v", SUCCESS_MESSAGE, response_string)
		t.Fail()
	}
}

func TestGetMatchNotFound(t *testing.T) {
	resp, err := http.Get(MATCH_FIND + "?numbermatch=55")

	if err != nil {
		t.Errorf("Test error on get execution: %v", err)
	}

	response_string := DecodeTest(resp.Body)

	if response_string != NOT_FOUND_ERROR_MESSAGE {
		t.Logf("Expected: %v. Response: %v", NOT_FOUND_ERROR_MESSAGE, response_string)
		t.Fail()
	}
}
