package test

import (
	"bytes"
	"io"
)

func DecodeTest(body io.Reader) string {

	buf := new(bytes.Buffer)
	buf.ReadFrom(body)
	s := buf.String()

	return s
}

const (
	DECODE_ERROR_MESSAGE          = `{"code": 400, "message": "Invalid Parameters!"}`
	ENCODE_ERROR_MESSAGE          = `{"code": 400, "message": "Invalid Parameters!"}`
	VALIDATE_STRUCT_ERROR_MESSAGE = `{"code": 400, "message": "Invalid Struct Body Content"}`
	NOT_FOUND_ERROR_MESSAGE       = `{"code": 404, "message": "Not Found"}`
	SERVER_ERROR_MESSAGE          = `{"code": 500, "message": "Internal Database Error"}`
	SUCCESS_MESSAGE               = `{"code": 200, "message": "Successful!"}`
)
