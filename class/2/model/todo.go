package model

type Todo struct {
	Id          int    `json:"id"`
	Description string `json:"description,omitempty"`
	Completed   bool   `json:"completed"`
}
