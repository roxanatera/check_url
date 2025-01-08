# Etapa de construcción
FROM golang:1.20 AS builder  
# Establecer el directorio de trabajo
WORKDIR /app

# Copiar los archivos necesarios
COPY go.mod go.sum ./
RUN go mod download
COPY . .

# Compilar la aplicación para Linux (binario estático)
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o url-checker .

# Etapa de producción
FROM alpine:3.18.3  

# Instalar certificados raíz
RUN apk add --no-cache ca-certificates

# Establecer el directorio de trabajo
WORKDIR /app

# Copiar el binario y la carpeta de plantillas desde la etapa de construcción
COPY --from=builder /app/url-checker .
COPY --from=builder /app/templates ./templates

# Exponer el puerto 3000
EXPOSE 3000

# Ejecutar la aplicación
CMD ["./url-checker"]
