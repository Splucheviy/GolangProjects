package models

type Goods struct {
	Id       string    `json:"Id"`
	Name     string `json:"Name"`
	Provider int    `json:"Provider"`
	Price    int    `json:"Price"`
	Quantity int    `json:"Quantity"`
}
