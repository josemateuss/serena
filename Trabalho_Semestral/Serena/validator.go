package main

import (
	"strconv"
	"strings"
	"unicode"

	"gopkg.in/go-playground/validator.v9"
)

var validate *validator.Validate

var errormessage string

func StartValidator() {

	validate = validator.New()

	//------------------ Player Validator -----------------------------
	validate.RegisterValidation("iluminati-ursal", ValidateIlur)
	validate.RegisterValidation("fistname-valid", ValidateFirstName)
	validate.RegisterValidation("cpf-valid", ValidateCPF)

	//------------------ Championship Validator -----------------------------
	//validate.RegisterValidation("champname-valid", ValidateChampName)

}

func ValidateIlur(gender validator.FieldLevel) bool {

	g := gender.Field().String()

	if g == "iluminati" || g == "ursal" {
		errormessage = "ursal mata todos"
		return false
	}

	return true
}

func ValidateFirstName(firstname validator.FieldLevel) bool {

	for _, r := range firstname.Field().String() {
		if !unicode.IsLetter(r) && !unicode.IsSpace(r) {
			return false
			//mensagem de erro: nome com número
		}
	}
	return true
}

func ValidateCPF(cpf validator.FieldLevel) bool {

	cpff := cpf.Field().String()

	cpff = strings.Replace(cpff, ".", "", -1)
	cpff = strings.Replace(cpff, "-", "", -1)
	if len(cpff) != 11 {
		return false
	}
	var eq bool
	var dig string
	for _, val := range cpff {
		if len(dig) == 0 {
			dig = string(val)
		}
		if string(val) == dig {
			eq = true
			continue
		}
		eq = false
		break
	}
	if eq {
		return false
	}

	i := 10
	sum := 0
	for index := 0; index < len(cpff)-2; index++ {
		pos, _ := strconv.Atoi(string(cpff[index]))
		sum += pos * i
		i--
	}

	prod := sum * 10
	mod := prod % 11
	if mod == 10 {
		mod = 0
	}
	digit1, _ := strconv.Atoi(string(cpff[9]))
	if mod != digit1 {
		return false
	}
	i = 11
	sum = 0
	for index := 0; index < len(cpff)-1; index++ {
		pos, _ := strconv.Atoi(string(cpff[index]))
		sum += pos * i
		i--
	}
	prod = sum * 10
	mod = prod % 11
	if mod == 10 {
		mod = 0
	}
	digit2, _ := strconv.Atoi(string(cpff[10]))
	if mod != digit2 {
		return false
	}

	return true
}
