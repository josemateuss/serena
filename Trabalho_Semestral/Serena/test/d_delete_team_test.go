package test

import (
	"net/http"
	"strings"
	"testing"
)

const (
	TEAM_DELETE = "http://localhost:8080/team/delete"
)

func DeleteRemoveTeam(name string) (resp *http.Response, err error) {
	client := &http.Client{}
	req, _ := http.NewRequest("DELETE", TEAM_DELETE+"?name="+name, strings.NewReader(""))

	resp, err = client.Do(req)

	return
}

func TestDeleteTeamSuccess(t *testing.T) {
	resp, err := DeleteRemoveTeam("Team One")

	if err != nil {
		t.Errorf("Test error on delete execution: %v", err)
	}

	response_string := DecodeTest(resp.Body)

	if response_string != SUCCESS_MESSAGE {
		t.Logf("Expected: %v. Response: %v", SUCCESS_MESSAGE, response_string)
		t.Fail()
	}
}

/*func TestDeleteTeamNotFound(t *testing.T) {
	resp, err := DeleteRemoveTeam("Team Vasco")

	if err != nil {
		t.Errorf("Test error on delete execution: %v", err)
	}

	response_string := DecodeTest(resp.Body)

	if response_string != NOT_FOUND_ERROR_MESSAGE {
		t.Logf("Expected: %v. Response: %v", NOT_FOUND_ERROR_MESSAGE, response_string)
		t.Fail()
	}
}*/
