# Stage 1: Build
FROM golang:1.24.2-alpine AS builder

# Define o diretório de trabalho
WORKDIR /app

# Copia os arquivos go.mod e go.sum para instalar dependências
COPY go.mod go.sum ./
RUN go mod download

# Copia o código-fonte
COPY . .

# Compila a aplicação com otimizações
RUN CGO_ENABLED=0 GOOS=linux go build -o api main.go

# Stage 2: Runtime
FROM alpine:latest

# Instala certificados CA para chamadas HTTPS, se necessário
RUN apk --no-cache add ca-certificates

# Define o diretório de trabalho
WORKDIR /app

# Copia o binário compilado da stage de build
COPY --from=builder /app/api .

# Expose a porta 8080
EXPOSE 8080

# Comando para executar a API
CMD ["./api"]