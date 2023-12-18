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
			products := []Product{}
			for _, v := range productList {
				products = append(products, v)
			}

			return c.JSON(http.StatusOK, products)
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


	e.GET("/inventory", func(c echo.Context) error {
		paramInventorytype := c.QueryParam("type")

		if len(paramInventorytype) == 0 {
			inventoryItems := []Inventory{}
			for _, v := range inventory {
				inventoryItems = append(inventoryItems, v)
			}

			return c.JSON(http.StatusOK, inventoryItems)
		}

		inventoryType, err := strconv.Atoi(paramInventorytype)
		if err != nil {
			return c.String(http.StatusBadRequest, "Invalid inventory type")
		}

		productType, ok := toType[inventoryType]
		if !ok {
			return c.JSON(http.StatusNoContent, "Product type not found")
		}

		return c.JSON(http.StatusOK, inventory[productType])
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
	productList[2] = Product{
		ID:   2,
		Name: "Sprite",
		Type: drinks,
	}
	productList[3] = Product{
		ID:   3,
		Name: "Lays",
		Type: food,
	}
	productList[4] = Product{
		ID:   4,
		Name: "Cheetos",
		Type: food,
	}
	productList[5] = Product{
		ID:   5,
		Name: "Toilet Paper",
		Type: utils,
	}

	// populate inventory
	inventory[drinks] = Inventory{
		Type: drinks,
		Products: map[int]Product{
			1: productList[1],
			2: productList[2],
		},
	}
	inventory[food] = Inventory{
		Type: food,
		Products: map[int]Product{
			3: productList[3],
			4: productList[4],
		},
	}
	inventory[utils] = Inventory{
		Type: utils,
		Products: map[int]Product{
			5: productList[5],
		},
	}
}
