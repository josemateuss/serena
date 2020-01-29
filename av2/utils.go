package main

import("io"
	   "log" 	 
	   "encoding/json")

func DecodeJson(payload io.Reader, entity interface{}) (err error){

	decoder := json.NewDecoder(payload)

	err = decoder.Decode(entity)

	if err != nil {
		log.Printf("[ERROR] could not convert json, because: %v", err)
		return
	}

	return
}