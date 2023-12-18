package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// Fetch the product - id / list of ids
// insert the product - product schema - auto generate ID

type Type int

const (
	food   Type = 0
	drinks Type = 1
	utils  Type = 2
)

type Inventory struct {
	Type Type
	IDs  map[int]Product
}

type Product struct {
	ID   int
	Name string
	Type Type
}

var inventory = map[Type]Inventory{}

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
