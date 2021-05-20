package models

import (
	"errors"
	"net/http"
	"time"
)

type Item struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Creation    time.Time `json:"created_at"`
}

type ItemList struct {
	Items []Item `json:"items"`
}

// Validate reads the input and validates it
func (item *Item) Validate(r *http.Request) error {
	if item.Name == "" {
		return errors.New("name is required")
	}

	if item.Description == "" {
		return errors.New("description is required")
	}

	return nil
}

// Render renders Item
func (*Item) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// Render renders ItemList
func (*ItemList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
