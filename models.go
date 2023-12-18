package main

type ProductType int

const (
	food   ProductType = 0
	drinks ProductType = 1
	utils  ProductType = 2
)

var toType = map[int]ProductType{
	0: food,
	1: drinks,
	2: utils,
}

type Inventory struct {
	Type ProductType     `json:"type"`
	IDs  map[int]Product `json:"products"`
}

type Product struct {
	ID   int         `json:"id"`
	Name string      `json:"name"`
	Type ProductType `json:"type"`
}
