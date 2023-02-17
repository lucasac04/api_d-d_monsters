package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func addMonster(context *gin.Context) {
	var newMonster Monster //criadno uma variável do tipo Monster

	if err := context.BindJSON(&newMonster); err != nil {
		return // como se desse um break caso a estrutura não bata com a do formato json de monster
	}

	monsters = append(monsters, newMonster) // adiciona newMonster no final de monsters

	context.IndentedJSON(http.StatusOK, newMonster) //hhtp status para dizer que chegou e não nenhum outro erro, newMonster é a variável que vai ter seu valor atribuído
}
