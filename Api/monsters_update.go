package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func updateMonster(context *gin.Context) {

	var updatedMonster Monster

	id := context.Param("id")
	monster, error := getMonsterById(id)

	if error != nil {
		return
	}

	if err := context.Bind(&updatedMonster); err != nil {
		return
	}

	monster.ID = id
	monster.Name = updatedMonster.Name
	monster.Damage = updatedMonster.Damage

	context.IndentedJSON(http.StatusOK, monster)

}
