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

func main() {
	e := echo.New()

	e.GET("/products", func(c echo.Context) error {
		paramProductId := c.QueryParam("id")

		if len(paramProductId) == 0 {
			return c.JSON(http.StatusOK, productList)
		}

		prodId, err := strconv.Atoi(paramProductId)
		if err != nil {
			return c.String(http.StatusBadRequest, "Invalid product ID")
		}

		product, ok := productList[prodId]
		if !ok {
			return c.JSON(http.StatusNoContent, "Product not found")
		}

		return c.JSON(http.StatusOK, product)
	})


	e.Logger.Fatal(e.Start(":1323"))
}

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
