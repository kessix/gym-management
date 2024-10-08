# Etapa de construção
FROM golang:1.22-alpine AS builder

# Instalar dependências necessárias
RUN apk update && apk add --no-cache git

# Definir o diretório de trabalho
WORKDIR /app

# Copiar go.mod e go.sum para instalação das dependências
COPY go.mod go.sum ./
RUN go mod download

# Copiar o restante do código
COPY . .

# Compilar a aplicação com CGO desativado para binário estático
RUN CGO_ENABLED=0 GOOS=linux go build -o gym-management

# Etapa de produção
FROM alpine:3.17

# Instalar dependências necessárias para rodar a aplicação
RUN apk update && apk add --no-cache ca-certificates

# Definir variáveis de ambiente
ENV GIN_MODE=release

# Criar um diretório para a aplicação
WORKDIR /app

# Copiar o binário da etapa de construção para o contêiner final
COPY --from=builder /app/gym-management .
COPY --from=builder /app/static ./static
COPY --from=builder /app/templates ./templates

# Expor a porta da aplicação
EXPOSE 8080

# Comando para rodar a aplicação
CMD ["./gym-management"]