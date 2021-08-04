package model

type Athlete struct {
	ID         int    `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	BackNumber int    `json:"back_number,omitempty"`
}

type Athletes []Athlete
