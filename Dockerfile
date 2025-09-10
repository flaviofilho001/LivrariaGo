# Use a imagem oficial do Go como base
FROM golang:1.21-alpine AS builder

# Instalar git (necessário para baixar dependências)
RUN apk add --no-cache git

# Definir diretório de trabalho
WORKDIR /app

# Copiar go.mod e go.sum (se existirem)
COPY go.* ./

# Baixar dependências
RUN go mod download

# Copiar código fonte
COPY . .

# Compilar a aplicação
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Usar uma imagem Alpine mínima para produção
FROM alpine:latest

# Instalar ca-certificates para conexões HTTPS
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copiar o binário compilado
COPY --from=builder /app/main .

# Expor a porta da aplicação
EXPOSE 8080

# Comando para executar a aplicação
CMD ["./main"]