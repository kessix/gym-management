# Gymma - a software for gym management

1. Configuração do ambiente:
   
- Instalar Go e PostgreSQL.

2. Estrutura do projeto:

- main.go: Arquivo principal onde o servidor será iniciado. (Atualmente tudo está aqui, por enquanto!)
- static/: Contém arquivos estáticos como CSS e JavaScript.
- templates/: Contém templates HTML (usaremos Go templates).

Estrutura completa:

```html
gym-management/
│
├── main.go
├── index.html
└── static
    ├── css
    │   └── styles.css
    └── js
        └── main.js
