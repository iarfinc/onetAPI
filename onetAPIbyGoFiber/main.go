package main

import (
	"github.com/gofiber/fiber/v2"
)

type Listing struct {
	ID     int     `json:"id"`
	Nama   string  `json:"nama"`
	Harga  float64 `json:"harga"`
	Gender string  `json:"gender"`
	Size   string  `json:"size"`
}

var listings = []Listing{
	{ID: 1, Nama: "Nike Court Trackzip", Harga: 420000, Gender: "Hommes", Size: "L"},
	{ID: 2, Nama: "Puma x BMW Hooded Bomber", Harga: 550000, Gender: "Hommes", Size: "M"},
	{ID: 3, Nama: "Adidas Terrex", Harga: 605000, Gender: "Hommes", Size: "XL"},
	{ID: 4, Nama: "Baselayer Rebook", Harga: 100000, Gender: "Unisex", Size: "S"},
	{ID: 5, Nama: "Kappa Trackpants", Harga: 75000, Gender: "Femmes", Size: "S"},
	{ID: 6, Nama: "Ellese Remini Tracktop", Harga: 650000, Gender: "Unisex", Size: "M"},
}

func main() {
	app := fiber.New()

	app.Get("/listing", func(c *fiber.Ctx) error {
		return c.JSON(listings)
	})

	app.Listen(":8080")
}
