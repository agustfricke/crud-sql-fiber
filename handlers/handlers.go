package handlers

import (
	"fmt"
	"log"

	"github.com/agustfricke/crud-sql-fiber/database"
	"github.com/agustfricke/crud-sql-fiber/models"
	"github.com/gofiber/fiber/v2"
)

func CreateAlbum(c *fiber.Ctx) error {
    album := new(models.Album)

    if err := c.BodyParser(album); err != nil {
      return err
    }

    albID, err := database.AddAlbum(models.Album{
        Title:  album.Title,
        Artist: album.Artist,
        Price:  album.Price,
    })

    if err != nil {
        log.Fatal(err)
    }

    return c.JSON(albID)
}


func Hello(c *fiber.Ctx) error {
    artist := c.Params("artist")
    fmt.Println(artist)
    albums, err := database.AlbumsByArtist(artist)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Albums found: %v\n", albums)
    return c.JSON(albums)
}
