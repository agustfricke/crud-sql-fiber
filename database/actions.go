package database

import (
	"database/sql"
	"fmt"

	"github.com/agustfricke/crud-sql-fiber/models"
)

func AddAlbum(alb models.Album) (models.Album, error) {
    result, err := DB.Exec("INSERT INTO album (title, artist, price) VALUES (?, ?, ?)", alb.Title, alb.Artist, alb.Price)
    if err != nil {
        return models.Album{}, fmt.Errorf("addAlbum: %v", err)
    }

    id, err := result.LastInsertId()
    if err != nil {
        return models.Album{}, fmt.Errorf("addAlbum: %v", err)
    }

    insertedAlbum := models.Album{
      ID: id,
      Title: alb.Title,
      Artist: alb.Artist,
      Price: alb.Price,
    }

    return insertedAlbum, nil
}

// albumsByArtist queries for albums that have the specified artist name.
func AlbumsByArtist(name string) ([]models.Album, error) {
    // An albums slice to hold data from returned rows.
    var albums []models.Album

    rows, err := DB.Query("SELECT * FROM album WHERE artist = ?", name)
    if err != nil {
        return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
    }
    defer rows.Close()
    // Loop through rows, using Scan to assign column data to struct fields.
    for rows.Next() {
        var alb models.Album
        if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
            return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
        }
        albums = append(albums, alb)
    }
    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
    }
    return albums, nil
}

// albumByID queries for the album with the specified ID.
func albumByID(id int64) (models.Album, error) {
    // An album to hold data from the returned row.
    var alb models.Album

    row := DB.QueryRow("SELECT * FROM album WHERE id = ?", id)
    if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
        if err == sql.ErrNoRows {
            return alb, fmt.Errorf("albumsById %d: no such album", id)
        }
        return alb, fmt.Errorf("albumsById %d: %v", id, err)
    }
    return alb, nil
}
