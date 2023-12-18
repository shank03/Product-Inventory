package main

type ProductType int

const (
	food   ProductType = 0
	drinks ProductType = 1
	utils  ProductType = 2
)

type Inventory struct {
	Type ProductType
	IDs  map[int]Product
}

type Product struct {
	ID   int         `json:"id"`
	Name string      `json:"name"`
	Type ProductType `json:"type"`
}
