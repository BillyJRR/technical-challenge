FROM golang:1.22-alpine AS builder

WORKDIR /app

# Copiar archivos de dependencias
COPY go.mod go.sum ./
RUN go mod download

# Copiar el código fuente
COPY . .

# Compilar el binario estático sin CGO
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/api/main.go

# --- Etapa 2: Ejecución (Production) ---
FROM alpine:3.19

# Instalar certificados CA para permitir peticiones HTTPS salientes
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copiar únicamente el binario compilado de la etapa anterior
COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]