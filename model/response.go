package model

import "fmt"

type ErrorV1 struct {
	Detail  string `json:"detail"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type ErrorV1s []ErrorV1

func (e ErrorV1) Error() string {
	return fmt.Sprintf("Detail: %s, Message: %s, Code: %d", e.Detail, e.Message, e.Code)
}
