package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/monsters", getMonsters)
	router.POST("/monsters", addMonster)
	router.GET("/monsters/:id", getMonster)
	router.Run("localhost:9090")
}
