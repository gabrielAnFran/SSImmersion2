package main

import (
	"cadastro/src/crud"
	"log"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := gin.Default()
	crud.ImigracaoInicial()

	r.POST("/usuarios", crud.CreateUser)
	r.GET("/usuarios", crud.GetUsers )
	r.GET("/usuarios/:id", crud.GetUser )
	//r.GET("/usuarios/:cep", crud.GetUsersByCity)
	r.PUT("/usuarios/:id",crud.UpdateUser)
	r.DELETE("/usuarios/:id", crud.DeleteUser)
	

	if erro := r.Run(":8081"); erro != nil {
		log.Fatal(erro.Error())
	}

}


