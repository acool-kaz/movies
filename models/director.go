package models

import "errors"

type Director struct {
	FirstName string `josn:"firstname"`
	LastName  string `json:"lastname"`
}

func (d *Director) Validation() error {
	if d.FirstName == "" {
		return errors.New("first name field empty")
	}
	if d.LastName == "" {
		return errors.New("last name field empty")
	}
	return nil
}
