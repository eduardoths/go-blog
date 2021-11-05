package input

import (
	"errors"
	"regexp"
)

func ValidateNameField(name string) error {
	match, _ := regexp.Match("^([A-Za-z]|[ ]){3,256}$", []byte(name))
	if !match {
		return errors.New("name.invalid")
	}
	return nil
}

func ValidateEmailField(email string) error {
	match, _ := regexp.Match("^[A-Za-z0-9+_.-]+@(.+)$", []byte(email))
	if !match {
		return errors.New("email.invalid")
	}
	return nil
}
