package main

const (
	SERVER_HOST   = "localhost:8080"
	MONGO_HOST    = "localhost:27017"
	DATABASE_NAME = "Serena"
)

const (
	PLAYERS_COLLECTION       = "players"
	CHAMPIONSHIPS_COLLECTION = "championships"
	MATCHES_COLLECTION       = "matches"
	TEAMS_COLLECTION         = "teams"
)

const (
	DECODE_ERROR_MESSAGE          = `{"code": 400, "message": "Invalid Parameters!"}`
	ENCODE_ERROR_MESSAGE          = `{"code": 400, "message": "Invalid Parameters!"}`
	VALIDATE_STRUCT_ERROR_MESSAGE = `{"code": 400, "message": "Invalid Struct Body Content"}`
	NOT_FOUND_ERROR_MESSAGE       = `{"code": 404, "message": "Not Found"}`
	CONFLICT_ERROR_MESSAGE        = `{"code": 409, "message": "Json already exists"}`
	SERVER_ERROR_MESSAGE          = `{"code": 500, "message": "Internal Database Error"}`
	SUCCESS_MESSAGE               = `{"code": 200, "message": "Successful!"}`
)
