package handler

import (
	"database/sql"
	"net/http"

	"example/web-service/internal/models"
	"github.com/gin-gonic/gin"
)

type AlbumHandler struct {
	DB *sql.DB
}


func (h *AlbumHandler) GetAlbums(c *gin.Context) {
	rows, err := h.DB.Query(
		`SELECT id, title, artist, price FROM albums`,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to fetch albums",
		})
		return
	}
	defer rows.Close()

	var albums []models.Album

	for rows.Next() {
		var a models.Album
		if err := rows.Scan(&a.ID, &a.Title, &a.Artist, &a.Price); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "failed to read album",
			})
			return
		}
		albums = append(albums, a)
	}

	c.JSON(http.StatusOK, albums)
}



func (h *AlbumHandler) CreateAlbum(c *gin.Context) {
	var input models.Album

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	err := h.DB.QueryRow(
		`INSERT INTO albums (title, artist, price)
		 VALUES ($1, $2, $3)
		 RETURNING id`,
		input.Title,
		input.Artist,
		input.Price,
	).Scan(&input.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to create album",
		})
		return
	}

	c.JSON(http.StatusCreated, input)
}
func (h *AlbumHandler) GetAlbumByID(c *gin.Context) {
	id := c.Param("id")

	var a models.Album

	err := h.DB.QueryRow(
		`SELECT id, title, artist, price FROM albums WHERE id = $1`,
		id,
	).Scan(&a.ID, &a.Title, &a.Artist, &a.Price)

	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "album not found",
		})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to fetch album",
		})
		return
	}

	c.JSON(http.StatusOK, a)
}
