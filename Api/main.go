package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Monster struct { //criando como se fosse uma classe
	ID     string `json:"id"`
	Name   string `json:"name"`
	Damage string `json:"damage"`
}

var monsters = []Monster{ //criando o formato do arquivo json
	{ID: "1", Name: "Bee", Damage: "d20"},
	{ID: "2", Name: "Ted Bear", Damage: "d8"},
	{ID: "3", Name: "Monkey", Damage: "d4"},
}

func getMonsters(context *gin.Context) { // get
	context.IndentedJSON(http.StatusOK, monsters) //hhtp status para dizer que chegou e não nenhum outro erro, monsters que é a estrutura json que vai mandar
}

func getMonster(context *gin.Context) {
	id := context.Param("id")
	monster, error := getMonsterById(id)
	if error != nil {
		context.IndentedJSON(http.StatusNoContent, gin.H{"message": "Monster Not Found"})
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

func addMonster(context *gin.Context) {
	var newMonster Monster //criadno uma variável do tipo Monster

	if err := context.BindJSON(&newMonster); err != nil {
		return // como se desse um break caso a estrutura não bata com a do formato json de monster
	}

	monsters = append(monsters, newMonster) // adiciona newMonster no final de monsters

	context.IndentedJSON(http.StatusOK, newMonster) //hhtp status para dizer que chegou e não nenhum outro erro, newMonster é a variável que vai ter seu valor atribuído
}

func main() {
	router := gin.Default()
	router.GET("/monsters", getMonsters)
	router.POST("/monsters", addMonster)
	router.GET("/monsters/:id", getMonster)
	router.Run("localhost:9090")
}
