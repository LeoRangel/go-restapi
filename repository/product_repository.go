package repository

import (
	"database/sql"
	"fmt"
	"go-restapi/model"
)

// doc: struct é um tipo de dado composto, semelhante as classes encontradas em outras linguagens
// doc: os structs são como "objetos sem métodos" e podemos adicionar métodos a elas usando receivers
// doc: ao add o receiver (pr *ProductRepository) as funções se tornam métodos de ProductRepository
type ProductRepository struct {
	connection *sql.DB
}

func NewProductRepository(connection *sql.DB) ProductRepository {
	return ProductRepository{
		connection: connection,
	}
}

// função responsável por pegar dados de produtos no banco
func (pr *ProductRepository) GetProducts() ([]model.Product, error) {
	query := "SELECT id, product_name, price FROM product"
	// recupera rows de dados do banco
	rows, err := pr.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.Product{}, err
	}

	var productList []model.Product
	var productObj model.Product

	// itera rows e salva os dados de produtos de cada uma na lista, se não houver erro
	for rows.Next() {
		// doc: o & indica que é para usar o endereço de memória
		err = rows.Scan(
			&productObj.ID,
			&productObj.Name,
			&productObj.Price)

		if err != nil {
			fmt.Println(err)
			return []model.Product{}, err
		}

		productList = append(productList, productObj)
	}

	rows.Close()

	return productList, nil
}

func (pr *ProductRepository) CreateProduct(product model.Product) (int, error) {
	var id int

	query, err := pr.connection.Prepare("INSERT INTO product" +
		"(product_name, price)" +
		" VALUES ($1, $2) RETURNING id")
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	
	// o Scan salva o id retornado na variável (&id indica o endereço da variável), se query tiver sucesso
	err = query.QueryRow(product.Name, product.Price).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	query.Close()
	return id, nil
}