package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/monsters", getMonsters) //read
	router.GET("/monsters/:id", getMonster)
	router.POST("/monsters", addMonster)       //create
	router.PUT("/monsters/:id", updateMonster) //update
	//router.DELETE("/monsters/:id", getMonster)   //delete

	router.Run("localhost:9090")
}
