# Usar una imagen base de Go
FROM golang:1.20 AS builder

# Establecer el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiar los archivos del proyecto al contenedor
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Compilar la aplicación
RUN go build -o url-checker .

# Imagen final más ligera
FROM alpine:latest

# Copiar el binario compilado desde el contenedor de compilación
WORKDIR /app
COPY --from=builder /app/url-checker .

# Exponer el puerto 3000
EXPOSE 3000

# Comando para ejecutar la aplicación
CMD ["./url-checker"]
