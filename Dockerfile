# Use uma imagem base com Go instalado
FROM golang:1.20 AS builder

# Defina o diretório de trabalho
WORKDIR /app

# Copie o go.mod e go.sum e faça o download das dependências
COPY go.mod go.sum ./
RUN go mod tidy

# Copie o restante do código fonte
COPY . .

# Construa o binário
RUN go build -o main ./cmd

# Use uma imagem base mínima para o runtime
FROM alpine:latest

# Instale o pacote necessário para o binário
RUN apk --no-cache add ca-certificates

# Copie o binário da etapa de build
COPY --from=builder /app/main /app/main

# Defina o comando de inicialização
ENTRYPOINT ["/app/main"]

# Exponha a porta que o aplicativo vai usar
EXPOSE 8080
