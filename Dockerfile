# Etapa 1: Compilar el binario
FROM golang:1.25.3-alpine AS build
WORKDIR /app

# Copiar dependencias y descargar módulos
COPY go.mod go.sum ./
RUN go mod download

# Copiamos todo el proyecto
COPY . .

# Compilar el binario
RUN go build -o main .

# Etapa 2: Imagen final
FROM alpine:latest
WORKDIR /root/

# Copiar binario y .env
COPY --from=build /app/main .
COPY .env .env

# Exponer el puerto 8090
EXPOSE 8090

# Ejecutar la aplicación
CMD ["./main"]

