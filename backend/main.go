package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "github.com/gin-contrib/cors"
)

type album struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}

var albums = []album{
    {ID: "1", Title: "November Rain", Artist: "Guns n' Roses", Price: 56.99},
    {ID: "2", Title: "Levels", Artist: "Avicii", Price: 17.99},
    {ID: "3", Title: "Rap God", Artist: "Eminem", Price: 39.99},
}

func main() {
  router := gin.Default()
  config := cors.DefaultConfig()
  config.AllowAllOrigins = true
  router.Use(cors.New(config))
  router.GET("/albums", getAlbums)

  router.Run()
}

func getAlbums(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, albums)
}
