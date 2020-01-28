package model

type MenuItem struct {
	Name  string `json:"name"`
	Price string `json:"price,omitempty"`
}
