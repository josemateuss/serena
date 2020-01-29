package main

import (
	"encoding/json"
	"io"
	"log"
)

func DecodeJson(payload io.Reader, entity interface{}) (err error) {

	decoder := json.NewDecoder(payload)

	err = decoder.Decode(entity)

	if err != nil {
		log.Printf("[ERROR] could not convert json, because: %v", err)
		return
	}

	return
}

func EncodeJson(payload io.Writer, entity interface{}) (err error) {

	encoder := json.NewEncoder(payload)

	err = encoder.Encode(entity)

	if err != nil {
		log.Printf("[ERROR] could not convert json, because: %v", err)
		return
	}

	return
}