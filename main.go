package main

import (
	"log"

	"github.com/agustfricke/crud-sql-fiber/database"
	"github.com/agustfricke/crud-sql-fiber/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)


func main() {

    database.ConnectDB()

    app := fiber.New()

    app.Use(cors.New(cors.Config{
        AllowOrigins: "http://localhost:5173",
        AllowMethods: "GET, POST, PUT, DELETE",
        AllowCredentials: true,
        AllowHeaders: "Origin, Content-Type, Accept",
    }))

    router.SetupRoutes(app)
    log.Fatal(app.Listen(":6969"))
    
    /*
    albums, err := albumsByArtist("John Coltrane")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Albums found: %v\n", albums)

    // Hard-code ID 2 here to test the query.
    alb, err := albumByID(2)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Album found: %v\n", alb)

    albID, err := addAlbum(Album{
        Title:  "The Modern Sound of Betty Carter",
        Artist: "Betty Carter",
        Price:  49.99,
    })
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("ID of added album: %v\n", albID)
    */

}

