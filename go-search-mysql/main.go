package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"

	"github.com/bxcodec/faker/v3"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	Id          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Price       int    `json:"price"`
}

func main() {
	db, err := gorm.Open(mysql.Open("golang:golang@/go_search"), &gorm.Config{})
	if err != nil {
		panic("could not connect to the database")
	}

	db.AutoMigrate(&Product{})

	app := fiber.New()

	app.Use(cors.New())

	app.Post("/api/products/populate", func(c *fiber.Ctx) error {
		for i := 0; i < 50; i++ {
			db.Create(&Product{
				Title:       faker.Word(),
				Description: faker.Paragraph(),
				Image:       fmt.Sprintf("https://picsum.photos/720/720?%s", faker.UUIDDigit()),
				Price:       rand.Intn(90) + 10,
			})
		}

		return c.JSON(fiber.Map{
			"message": "success",
		})
	})

	app.Get("/api/products/frontend", func(c *fiber.Ctx) error {
		var products []Product

		db.Find(&products)

		return c.JSON(products)
	})

	app.Get("/api/products/backend", func(c *fiber.Ctx) error {
		var products []Product

		sql := "SELECT * FROM products"

		if s := c.Query("s"); s != "" {
			sql = fmt.Sprintf("%s WHERE title LIKE '%%%s%%' OR description LIKE '%%%s%%'", sql, s, s)
		}

		if sort := c.Query("sort"); sort != "" {
			sql = fmt.Sprintf("%s ORDER BY price %s", sql, sort)
		}

		page, _ := strconv.Atoi(c.Query("page", "1"))
		perPage := 9
		var total int64

		db.Raw(sql).Count(&total)

		sql = fmt.Sprintf("%s LIMIT %d OFFSET %d", sql, perPage, (page-1)*perPage)

		db.Raw(sql).Scan(&products)

		return c.JSON(fiber.Map{
			"data":      products,
			"total":     total,
			"page":      page,
			"last_page": math.Ceil(float64(total) / float64(perPage)),
		})
	})

	app.Listen(":8000")
}
