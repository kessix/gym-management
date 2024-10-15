# Gymma - a software for gym management

1. Configuração do ambiente:
   
- Instalar Go, Gin, GORM, e PostgreSQL.

2. Estrutura do projeto:

- main.go: Arquivo principal onde o servidor será iniciado. (Atualmente tudo está aqui!)
(Para implementar!) - models/: Contém os modelos do banco de dados.
(Para implementar!) - controllers/: Contém a lógica de CRUD.
(Para implementar!) - routes/: Define as rotas.
(Para implementar!) - config/: Contém as configurações (incluindo conexão com o banco de dados).
- static/: Contém arquivos estáticos como CSS e JavaScript.
- templates/: Contém templates HTML (usaremos Go templates).

Estrutura completa:

```html
gym-management/
<<<<<<< HEAD
│
├── main.go
├── index.html
└── static
    ├── css
    │   └── styles.css
    └── js
        └── main.js
=======
├── go.mod
├── go.sum
├── Dockerfile
├── docker-compose.yml
└── main.go
>>>>>>> 76f1bac29e20a11c2a049c6dbf44f38ed699005e
