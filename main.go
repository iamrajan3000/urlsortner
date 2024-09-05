package main

import (
	"github.com/gin-gonic/gin"

	handlers "urlsortner/handler"
	"urlsortner/services"
	"urlsortner/store"
)

func main() {
	s := store.New()
	urlService := services.New(s)
	urlHandler := handlers.New(urlService)

	r := gin.Default()

	r.POST("/shorten", urlHandler.ShortenHandler)
	r.GET("/:shortURL", urlHandler.RedirectHandler)

	r.Run(":8080")
}
