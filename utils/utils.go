package utils

import (
	"strconv"
	"github.com/goinggo/tracelog"
)

func StringToFloat32(text string) (number float32) {

	value, err := strconv.ParseFloat(text, 32)
	if err != nil {
		tracelog.Errorf(err, "utils", "stringToFloat32", "It was not possible to convert text into float32")
	}
	number = float32(value)

	return
}

func StringToBool(text string) (bool) {

	response := false

	if text == "true" {
		response = true
	}

	return response
}

func BoolToString(isBool bool) (string) {

	response := "false"

	if isBool {
		response = "true"
	}

	return response
}