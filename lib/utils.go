package lib

import (
	"errors"

	"alilacream/socialx/lib/helpers"
)

func ParseEmail(s string) error {
	if helpers.HasUnwantedChar(s) {
		return errors.New("Invalid Email")
	}
	return nil
}

func ParseUsername(s string) error {
	if helpers.HasUnwantedChar(s) {
		return errors.New("Invalid Username")
	}
	return nil
}
