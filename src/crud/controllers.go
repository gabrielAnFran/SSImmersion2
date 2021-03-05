package crud

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/badoux/checkmail"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //
)

var db *gorm.DB
//Usuario é o modelo
type Usuario struct{
	gorm.Model
	Nome 		string
	Cpf			int
	Telefone	int
	Email		string
	Rua			string
	Bairro		string
	Complemento string
	Cep			int 

}
//ImigracaoInicial inicia a conexao com bando de dados
func ImigracaoInicial(){
	stringConexao :=  "golang:golang@/cadastroCompleto?charset=utf8&parseTime=True&loc=Local"
	db, erro := gorm.Open("mysql", stringConexao)
	if erro != nil{
		fmt.Println(erro.Error())
		panic("Falha ao conectar ao banco de dados")
	}
	defer db.Close()

	db.AutoMigrate(&Usuario{})
}

//CreateUser cria um usuario
func CreateUser(c *gin.Context){
	stringConexao :=  "golang:golang@/cadastroCompleto?charset=utf8&parseTime=True&loc=Local"
	db, erro := gorm.Open("mysql", stringConexao)
	if erro != nil{
		fmt.Println(erro.Error())
		panic("Falha ao conectar ao banco de dados")
	}
	defer db.Close()

	fmt.Println("endpoint hit")
 	var reqBody Usuario
	if erro := c.ShouldBindJSON(&reqBody); erro != nil{
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error" : true,
			"message" : "erro ao receber corpo da requisição",
		})
		return
	}

	erro = checkmail.ValidateFormat(reqBody.Email)
	if erro != nil{
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error" : true,
			"message" : "email em formato invalido",
		})
	} else{
		db.Create(&Usuario{Nome: reqBody.Nome, Cpf: reqBody.Cpf, Telefone: reqBody.Telefone, Email: reqBody.Email, Rua: reqBody.Rua, Bairro: reqBody.Bairro , Complemento: reqBody.Complemento, Cep: reqBody.Cep})

	c.JSON(http.StatusCreated, gin.H{
		"message": "Criado com sucesso",
	})
	}

	fmt.Println("endpoint hit 2")

	
	
}

//GetUsers query por todos usuarios
func GetUsers(c *gin.Context){
	stringConexao :=  "golang:golang@/cadastroCompleto?charset=utf8&parseTime=True&loc=Local"
	db, erro := gorm.Open("mysql", stringConexao)
	if erro != nil{
		fmt.Println(erro.Error())
		panic("Falha ao conectar ao banco de dados")
	}
	defer db.Close()

	var usuarios []Usuario

	db.Find(&usuarios)

	c.JSON(http.StatusOK, usuarios)
	

}
//GetUser busca um usuario
func GetUser(c *gin.Context){
	Param := c.Param("id")
	id,_ := strconv.Atoi(Param)
	stringConexao :=  "golang:golang@/cadastroCompleto?charset=utf8&parseTime=True&loc=Local"
	db, erro := gorm.Open("mysql", stringConexao)
	if erro != nil{
		fmt.Println(erro.Error())
		panic("Falha ao conectar ao banco de dados")
	}
	defer db.Close()

	var usuario Usuario

	db.Find(&usuario, id)

	c.JSON(http.StatusOK, usuario)
	

}
//DeleteUser deleta
func DeleteUser(c *gin.Context){
	Param := c.Param("id")
	id,_ := strconv.Atoi(Param)
	stringConexao :=  "golang:golang@/cadastroCompleto?charset=utf8&parseTime=True&loc=Local"
	db, erro := gorm.Open("mysql", stringConexao)
	if erro != nil{
		fmt.Println(erro.Error())
		panic("Falha ao conectar ao banco de dados")
	}
	defer db.Close()

	db.Delete(&Usuario{}, id); 


	c.JSON(http.StatusOK, gin.H{
		"message": "deletado com sucesso",
	})
	

}
//UpdateUser atualiza	
func UpdateUser(c *gin.Context){
	Param := c.Param("id")
	id,_ := strconv.Atoi(Param)
	stringConexao :=  "golang:golang@/cadastroCompleto?charset=utf8&parseTime=True&loc=Local"
	db, erro := gorm.Open("mysql", stringConexao)
		if erro != nil{
			fmt.Println(erro.Error())
			panic("Falha ao conectar ao banco de dados")
		}
	defer db.Close()

		fmt.Println("endpoint hit")
 	var reqBody Usuario
	if erro := c.ShouldBindJSON(&reqBody); erro != nil{
	c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error" : true,
			"message" : "erro ao receber corpo da requisição",
		})
		return
	}

		fmt.Println("endpoint hit 2")
		var usuario Usuario
		db.Find(&usuario, id)

		db.Model(&usuario).Updates(&Usuario{Nome: reqBody.Nome, Cpf: reqBody.Cpf,
			 Telefone: reqBody.Telefone, Email: reqBody.Email, Rua: reqBody.Rua,
			  Bairro: reqBody.Bairro , Complemento: reqBody.Complemento, Cep: reqBody.Cep})

		c.JSON(http.StatusOK, gin.H{
				"message": "Atualizado com sucesso",
			})
			

		
}






