package main

import (
	"go-restapi/controller"
	"go-restapi/db"
	"go-restapi/repository"
	"go-restapi/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	// config server com framework gin
	server := gin.Default()

	// conectando ao banco
	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	// doc: camada de repository - conecta ao banco
	// instanciando usecase de produtos
	ProductRepository := repository.NewProductRepository(dbConnection)
	
	// doc: camada de usecase - conecta controllers aos repositories e trata regras de negócio
	// instanciando usecase de produtos
	ProductUseCase := usecase.NewProductUseCase(ProductRepository)

	// doc: camada de controllers - retorna os dados
	// instanciando controller de produtos
	ProductController := controller.NewProductController(ProductUseCase)

	// rota para testar server
	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// rota para retornar produtos e função do controller responsável por fazer isso
	server.GET("/products", ProductController.GetProducts)

	// rota para criar novo produto
	server.POST("/product", ProductController.CreateProduct)
	
	// rota para buscar um produto
	server.GET("/product/:productId", ProductController.GetProductById)

	// TODO: add rota de delete e put (editar)
	// TODO: add autenticação jwt

	// rodando o projeto na porta indicada
	server.Run(":8000")
}