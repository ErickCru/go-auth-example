package main

import (
	"main/database"
	"main/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"

)

func read_db_env() string {
	var envs map[string]string
    envs, err := godotenv.Read(".env")

    if err != nil {
        panic("Error loading .env file")
    }

	user := envs["db_user"]
	pass := envs["db_pass"]
	db_name := envs["db_name"]

	url_db := user + ":" + pass + "@/" + db_name

	return url_db
}

func main() {

	url_db := read_db_env()

	database.Connect(url_db)

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.Setup(app)

	app.Listen(":8000")
}
