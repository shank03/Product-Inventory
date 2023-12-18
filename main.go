package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// Fetch the product - id / list of ids
// insert the product - product schema - auto generate ID

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
	ID   int
	Name string
	Type ProductType
}

var inventory = map[ProductType]Inventory{}

func main() {
	e := echo.New()
	e.GET("/product/:id", func(c echo.Context) error {
		paramProductId := c.Param("id")

		_, err := strconv.Atoi(paramProductId)
		if err != nil {
			return c.NoContent(http.StatusInternalServerError)
		}

		return c.String(http.StatusOK, paramProductId)
	})
	e.GET("/product", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
