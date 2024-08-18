package api

import (
	"net/http"
)

type validatorInterface interface {
	validate(control *Control, response http.Response)
}

type ContentTypeValidator struct {
	validatorInterface
}

func (*ContentTypeValidator) validate(control *Control, response http.Response) {
	if response.Header.Get("Content-Type") != "application/json" {
		control.Ok = false
		control.Message = "Content-Type validation failed. " + response.Header.Get("Content-Type") + " not supported."
	}
}
