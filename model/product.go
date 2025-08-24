package model

// config estrutura da tabela product
// doc: o `json:"id_product"` define como esse dado vai ser retornado
type Product struct {
	ID    int     `json:"id_product"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}