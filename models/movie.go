package models

import "errors"

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

func (m *Movie) Validation() error {
	if m.Isbn == "" {
		return errors.New("isbn field empty")
	}
	if m.Title == "" {
		return errors.New("title field empty")
	}
	if err := m.Director.Validation(); err != nil {
		return err
	}
	return nil
}
