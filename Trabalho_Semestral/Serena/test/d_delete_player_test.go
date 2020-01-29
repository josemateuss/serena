package test

import (
	"net/http"
	"strings"
	"testing"
)

const (
	PLAYER_DELETE = "http://localhost:8080/players/delete"
)

func DeleteRemovePlayer(cpf string) (resp *http.Response, err error) {
	client := &http.Client{}
	req, _ := http.NewRequest("DELETE", PLAYER_DELETE+"?cpf="+cpf, strings.NewReader(""))

	resp, err = client.Do(req)

	return
}

func TestDeletePlayerSuccess(t *testing.T) {
	resp, err := DeleteRemovePlayer("055.966.891-06")

	if err != nil {
		t.Errorf("Test error on delete execution: %v", err)
	}

	response_string := DecodeTest(resp.Body)

	if response_string != SUCCESS_MESSAGE {
		t.Logf("Expected: %v. Response: %v", SUCCESS_MESSAGE, response_string)
		t.Fail()
	}
}

/*func TestDeletePlayerNotFound(t *testing.T) {
	resp, err := DeleteRemovePlayer("056.696.901-76")

	if err != nil {
		t.Errorf("Test error on delete execution: %v", err)
	}

	response_string := DecodeTest(resp.Body)

	if response_string != NOT_FOUND_ERROR_MESSAGE {
		t.Logf("Expected: %v. Response: %v", NOT_FOUND_ERROR_MESSAGE, response_string)
		t.Fail()
	}
}*/
