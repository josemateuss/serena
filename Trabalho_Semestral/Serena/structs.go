package main

type Player struct {
	FirstName string `json:"firstname" validate:"fistname-valid,required"`
	CPF       string `json:"cpf" validate:"cpf-valid,required"`
	Birthdate string `json:"birthdate" validate:"required"`
	Phone     string `json:"phone" validate:"required"`
	Email     string `json:"email" validate:"required"`
	Password  string `json:"pass" validate:"required"`
	Group     string `json:"group" validate:"required"`
	Class     string `json:"class" validate:"required"`
	Category  string `json:"category" validate:"required"`
	Active    bool
}

type PlayerCPF struct {
	FirstName string `json:"firstname" validate:"fistname-valid,required"`
	CPF       string `json:"cpf" validate:"cpf-valid,required"`
}

type Championship struct {
	Name         string      `json:"name" validate:"required"`
	Players      []PlayerCPF `json:"players, string" validate:"required,min=2"`
	InitialDate  string      `json:"idate" validate:"required"`
	NumberOfSets int         `json:"sets" validate:"required"`
	Admin        PlayerCPF   `json:"admin, string" validate:"required"`
	Active       bool
}

type Table struct {
	NumberOfTable int          `json:"number" validate:"required"`
	Matches       []Match      `json:"matches" validate:"required,min=2"`
	Champ         Championship `json:"champ" validate:"required"`
}

type Match struct {
	NumberOfMatch int    `json:"numbermatch" validate:"required"`
	Player1       Player `json:"playerone" validate:"required"`
	Player2       Player `json:"playertwo" validate:"required"`
	PointPlayer1  int    `json:"point1"`
	GamesPlayer1  int    `json:"game1"`
	SetsPlayer1   int    `json:"set1"`
	PointPlayer2  int    `json:"point2"`
	GamesPlayer2  int    `json:"game2"`
	SetsPlayer2   int    `json:"set2"`
	MatchWinner   bool   `json:"winner"`
	Type          string
	Active        bool
}

type Ranking struct {
	Name          string             `json:"name" validate:"required"`
	Matches       []Match            `json:"matches" validate:"required"`
	Championships []ChampionshipName `json:"idate" validate:"required"`
	Players       []PlayerCPF        `json:"fdate" validate:"required"`
	Pontuation    string             `json:"pontuation" validate:"required"`
}

type Team struct {
	Name    string    `json:"name" validate:"required"`
	Player1 PlayerCPF `json:"playerone" validate:"required"`
	Player2 PlayerCPF `json:"playertwo" validate:"required"`
	Active  bool
}

type ChampionshipName struct {
	Name string `json:"name" validate:"required"`
}

type PlayerUpdate struct {
	FirstName string `json:"firstname" validate:"fistname-valid,required"`
	CPF       string `json:"cpf" validate:"cpf-valid,required"`
	Active    bool
}

type ChampionshipUpdate struct {
	Name   string `json:"name" validate:"required"`
	Active bool
}

type PlayerLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"pass" validate:"required"`
}
