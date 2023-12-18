package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// Fetch the product - id / list of ids
// insert the product - product schema - auto generate ID

var productList = map[int]Product{}
var inventory = map[ProductType]Inventory{}

func init() {
	// populate product list
	productList[1] = Product{
		ID:   1,
		Name: "Coke",
		Type: drinks,
	}
	productList[3] = Product{
		ID:   3,
		Name: "Sprite",
		Type: drinks,
	}
	productList[4] = Product{
		ID:   4,
		Name: "Lays",
		Type: food,
	}
	productList[6] = Product{
		ID:   6,
		Name: "Cheetos",
		Type: food,
	}
	productList[7] = Product{
		ID:   7,
		Name: "Toilet Paper",
		Type: utils,
	}

	// populate inventory
	inventory[drinks] = Inventory{
		Type: drinks,
		IDs: map[int]Product{
			1: productList[1],
			3: productList[3],
		},
	}
	inventory[food] = Inventory{
		Type: food,
		IDs: map[int]Product{
			4: productList[4],
			6: productList[6],
		},
	}
	inventory[utils] = Inventory{
		Type: utils,
		IDs: map[int]Product{
			7: productList[7],
		},
	}
}

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
