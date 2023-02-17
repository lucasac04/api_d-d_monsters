package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getMonsters(context *gin.Context) { // get
	context.IndentedJSON(http.StatusOK, monsters) //hhtp status para dizer que chegou e não nenhum outro erro, monsters que é a estrutura json que vai mandar
}
