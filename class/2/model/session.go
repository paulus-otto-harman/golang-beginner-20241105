package model

type Session struct {
	Id   string `json:"id"`
	User User   `json:"user,omitempty"`
}
