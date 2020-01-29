package test

import (
	"net/http"
	"strings"
	"testing"
)

const (
	CHAMPIONSHIP_DELETE = "http://localhost:8080/championships/delete"
)

func DeleteRemoveChamp(name string) (resp *http.Response, err error) {
	client := &http.Client{}
	req, _ := http.NewRequest("DELETE", CHAMPIONSHIP_DELETE+"?name="+name, strings.NewReader(""))

	resp, err = client.Do(req)

	return
}

func TestDeleteChampSuccess(t *testing.T) {
	resp, err := DeleteRemoveChamp("Champ1")

	if err != nil {
		t.Errorf("Test error on delete execution: %v", err)
	}

	response_string := DecodeTest(resp.Body)

	if response_string != SUCCESS_MESSAGE {
		t.Logf("Expected: %v. Response: %v", SUCCESS_MESSAGE, response_string)
		t.Fail()
	}
}

/*func TestDeleteChampNotFound(t *testing.T) {
	resp, err := DeleteRemoveChamp("Champ 44")

	if err != nil {
		t.Errorf("Test error on delete execution: %v", err)
	}

	response_string := DecodeTest(resp.Body)

	if response_string != NOT_FOUND_ERROR_MESSAGE {
		t.Logf("Expected: %v. Response: %v", NOT_FOUND_ERROR_MESSAGE, response_string)
		t.Fail()
	}
}*/
