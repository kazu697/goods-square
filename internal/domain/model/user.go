package model

type User struct {
	ID      int    `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Twitter string `json:"twitter,omitempty"`
	Level   int    `json:"level,omitempty"`
}
