package api

import (
	"net/http"
)

type Control struct {
	validators []validatorInterface
	Ok         bool   `json:"ok"`
	Message    string `json:"message"`
}

func NewControl() *Control {
	control := Control{
		Ok:         true,
		Message:    "Forxy pass OK.",
		validators: make([]validatorInterface, 0),
	}

	control.validators = append(control.validators, new(ContentTypeValidator))

	return &control
}

func (control *Control) Validate(response http.Response) {
	for _, validator := range control.validators {
		validator.validate(control, response)
	}
}
