package test

import (
	"net/http"
	"strings"
	"testing"
)

const (
	MATCH_DELETE = "http://localhost:8080/matches/delete"
)

func DeleteRemoveMatch(numbermatch string) (resp *http.Response, err error) {
	client := &http.Client{}
	req, _ := http.NewRequest("DELETE", MATCH_DELETE+"?numbermatch="+numbermatch, strings.NewReader(""))

	resp, err = client.Do(req)

	return
}

func TestDeleteMatchSuccess(t *testing.T) {
	resp, err := DeleteRemoveMatch("Match One")

	if err != nil {
		t.Errorf("Test error on delete execution: %v", err)
	}

	response_string := DecodeTest(resp.Body)

	if response_string != SUCCESS_MESSAGE {
		t.Logf("Expected: %v. Response: %v", SUCCESS_MESSAGE, response_string)
		t.Fail()
	}
}

/*func TestDeleteMatchNotFound(t *testing.T) {
	resp, err := DeleteRemoveMatch("MatchNotFound")

	if err != nil {
		t.Errorf("Test error on delete execution: %v", err)
	}

	response_string := DecodeTest(resp.Body)

	if response_string != NOT_FOUND_ERROR_MESSAGE {
		t.Logf("Expected: %v. Response: %v", NOT_FOUND_ERROR_MESSAGE, response_string)
		t.Fail()
	}
}*/
