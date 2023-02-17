package main

import (
	"errors"

	"net/http"

	"github.com/gin-gonic/gin"
)

func getMonster(context *gin.Context) {
	id := context.Param("id")
	monster, error := getMonsterById(id)
	if error != nil {
		context.IndentedJSON(http.StatusNoContent, gin.H{"message": "Monster Not Found"})
		return
	}
	context.IndentedJSON(http.StatusOK, monster)

}

func getMonsterById(id string) (*Monster, error) { // getById
	for i, t := range monsters {
		if t.ID == id {
			return &monsters[i], nil
		}
	}

	return nil, errors.New("monster not found")
}
