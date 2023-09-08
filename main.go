package main

import (
	"log"
	"net/http"
	"time"

	"go-url-shortener/dbUtils"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := dbUtils.InitDB()
	defer db.Close()

	if err != nil {
		log.Println(err)
	}

	dbClient := &dbUtils.DBClient{DB: db}

	router := gin.Default()

	router.GET("/:id", dbClient.GetURL)
	router.PUT("/create", dbClient.CreateShortURL)

	s := &http.Server{
		Addr:           ":4000",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
