package dbUtils

import (
	"database/sql"
	"go-url-shortener/base62"

	"github.com/gin-gonic/gin"
)

type DBClient struct {
	DB *sql.DB
}

func (driver *DBClient) CreateShortURL(c *gin.Context) {
	var reqBody struct {
		URL string `json:"url"`
	}
	c.Bind(&reqBody)

	var id int

	err := driver.DB.QueryRow(`INSERT INTO URLS(URL) VALUES($1) RETURNING ID`, reqBody.URL).Scan(&id)

	if err != nil {
		c.JSON(500, gin.H{
			"error":   "Internal server error",
			"status":  500,
			"message": "Error creating short url",
		})
		return
	}

	c.JSON(201, gin.H{
		"data": map[string]interface{}{"id": base62.ToBase62(id)},
	})
}

func (driver *DBClient) GetURL(c *gin.Context) {
	id := c.Param("id")
	dbEntryId := base62.ToBase10(id)

	var url string

	err := driver.DB.QueryRow(`SELECT URL FROM URLS WHERE ID = $1 LIMIT 1`, dbEntryId).Scan(&url)

	if err != nil {
		c.JSON(404, gin.H{
			"error":  "Not Found",
			"status": 404,
			"mesage": err.Error(),
		})

		return
	}

	c.JSON(200, gin.H{
		"data": map[string]interface{}{"url": url},
	})
}
