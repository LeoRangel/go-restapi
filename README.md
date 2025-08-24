# 🛍️ Product Management API (Go + Gin)

API REST de teste para **gerenciamento de produtos**, construída em **Go** com **Gin** e **Postgres**. Está preparada para rodar **localmente** (Go) ou via **Docker Compose**.

---

## 🚀 Ferramentas

- **Go** — linguagem principal.
- **Gin** — framework web (roteamento, middleware, handlers).
- **Postgres** — banco de dados relacional.
- **Docker & Docker Compose** — containerização da aplicação e do banco.

---

## 📐 Arquitetura (visão geral)

Arquitetura em camadas, inspirada em _Clean Architecture_:

- **model/** — entidades e estruturas de dados do domínio (ex.: `Product`).
- **repository/** — acesso a dados (consultas/commands no Postgres).
- **usecase/** — regras de negócio (ex.: criar produto, buscar por id).
- **controller/** — camada HTTP (handlers do Gin chamando os use cases).

Fluxo típico: **Controller → Use Case → Repository → Database**.

---

## 📂 Estrutura de pastas

.
├── cmd/
│ └── main.go
├── controller/
│ └── product_controller.go
├── model/
│ └── product.go
├── repository/
│ └── product_repository.go
├── usecase/
│ └── product_usecase.go
├── docker-compose.yml
└── Dockerfile

---

## 📌 Rotas

### Produtos

- `GET /products` — lista todos os produtos
- `POST /product` — cria um novo produto
- `GET /product/:productId` — obtém um produto pelo `id`

### Exemplos rápidos (cURL)

```bash
# Listar produtos
curl -i http://localhost:8000/products

# Criar produto (exemplo de payload)
curl -i -X POST http://localhost:8000/product \
  -H "Content-Type: application/json" \
  -d '{"name":"TV", "price":1999.90}'

# Buscar por ID
curl -i http://localhost:8000/product/1
```

---

## ▶️ Como rodar

### Opção 1 — Local (Go)

1. Instalar dependências:

```bash
go mod tidy
```

2. Subir apenas o Postgres (usando o serviço do Compose):

```bash
docker-compose up -d go_db
```

1. Rodar a aplicação:

```bash
go run cmd/main.go
```

- API: `http://localhost:8000`
- Postgres: `localhost:5432` (user: `postgres`, pass: `1234`, db: `postgres`)

---

### Opção 2 — Docker Compose (App + DB)

Subir tudo com build:

```bash
docker-compose up --build
```

- API: `http://localhost:8000`
- Postgres dentro da rede do Compose: host `go_db:5432`  
  (externo: `localhost:5432`)

---

## ✅ Checklist / Próximos passos

- [ ] Implementar **DELETE** e **PUT**
- [ ] Adicionar **JWT** (autenticação/autorização)
- [ ] Criar testes (unitários/integrados)
- [ ] Documentar a API (Swagger ou Postman)
