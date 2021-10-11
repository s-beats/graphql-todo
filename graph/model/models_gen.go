// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Item struct {
	ID         string `json:"id"`
	UnitPrice  int    `json:"unitPrice"`
	Amount     int    `json:"amount"`
	TotalPrice int    `json:"totalPrice"`
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type BuyItemInput struct {
	ID string `json:"id"`
}
