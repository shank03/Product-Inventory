package main

type ProductDto struct {
	Name string      `json:"name"`
	Type ProductType `json:"type"`
}
