package handlers

import (
	"fmt"
	"log"

	"github.com/agustfricke/crud-sql-fiber/database"
	"github.com/agustfricke/crud-sql-fiber/models"
	"github.com/gofiber/fiber/v2"
)

func UpdateAlbum(c *fiber.Ctx) error {
  id := c.Params("id")

  album := new(models.Album)
  if err := c.BodyParser(album); err != nil {
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Bad request"})
  }

  err := database.UpdateAlbum(id, *album)
  if err != nil {
    log.Fatal(err)
    return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Interal Sever Error"})
  }

  return c.JSON(album)
}

func DeleteAlbum(c *fiber.Ctx) error {
  id := c.Params("id")
  err := database.DeleteAlbum(id)
  if err != nil {
    log.Fatal(err)
    return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Interal Sever Error"})
  }
  return c.JSON(fiber.Map{"message": "Album deleted successfully"})
}

func GetAll(c *fiber.Ctx) error {
  albums, err := database.GetAllAlbums()
  if err != nil {
    log.Fatal(err)
  }
  return c.JSON(albums)

}

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

func GetByID(c *fiber.Ctx) error {
    id := c.Params("id")
    albums, err := database.AlbumByID(id)
    if err != nil {
        log.Fatal(err)
    }
    return c.JSON(albums)
}


func GetByName(c *fiber.Ctx) error {
    artist := c.Params("artist")
    fmt.Println(artist)
    albums, err := database.AlbumsByArtist(artist)
    if err != nil {
        log.Fatal(err)
    }
    return c.JSON(albums)
}
