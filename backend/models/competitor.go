package models

type Competitor struct {
	Id         int     `json:"id"`
	First_name string  `json:"firstName"`
	Last_name  string  `json:"lastName"`
	Weight     float32 `json:"weight"`
}