package util

import (
	"encoding/json"
)

type responseTop struct {
	Count    int      `json:"count"`
	TopWords []string `json:"top_words"`
}

type responseError struct {
	Code  int    `json:"Code"`
	Error string `json:"Error"`
}

type responseMessage struct {
	Code  int    `json:"Code"`
	Error string `json:"Message"`
}

func ErrorJson(code int, s string) (string, error) {
	resp := &responseError{
		Code:  code,
		Error: s}
	if bytes, err := json.Marshal(resp); err != nil {
		Log("Error err-json create", err)
		return "", err
	} else {
		return (string(bytes)), nil
	}
}

func MessageJson(code int, s string) (string, error) {
	resp := &responseMessage{
		Code:  code,
		Error: s}
	if bytes, err := json.Marshal(resp); err != nil {
		Log("Error message-json create", err)
		return "", err
	} else {
		return (string(bytes)), nil
	}
}

func TopWordsJson(list []string) (string, error) {
	resp := &responseTop{
		Count:    len(list),
		TopWords: list}
	if bytes, err := json.Marshal(resp); err != nil {
		Log("Error topwords-json create", err)
		return "", err
	} else {
		return (string(bytes)), nil
	}
}
