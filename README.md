# Gymma - a software for gym management

1. Configuração do ambiente:
   
- Instalar Go, Gin, GORM, e PostgreSQL.

2. Estrutura do projeto:

- main.go: Arquivo principal onde o servidor será iniciado.
- models/: Contém os modelos do banco de dados.
- controllers/: Contém a lógica de CRUD.
- routes/: Define as rotas.
- config/: Contém as configurações (incluindo conexão com o banco de dados).
- static/: Contém arquivos estáticos como CSS e JavaScript.
- templates/: Contém templates HTML (usaremos Go templates).

Estrutura completa:

```html
gym-management/
├── Config/
│   └── database.go
├── controllers/
│   └── userController.go
├── models/
│   └── user.go
├── routes/
│   └── routes.go
├── static/
│   ├── css/
│   │   └── styles.css
│   ├── js/
│   │   └── app.js
│   └── index.html
├── templates/
│   └── index.html
├── go.mod
├── go.sum
├── Dockerfile
├── docker-compose.yml
└── main.go