# Go Service Template

Este es un proyecto base generado por Backstage para servicios en Go.

## Estructura
- `go.mod`: Módulo Go
- `cmd/`: Punto de entrada principal
- `internal/handler/`: Handlers HTTP o lógicos
- `internal/service/`: Lógica de negocio

## Cómo compilar

```sh
go build -o bin/app ./cmd
```

## Cómo ejecutar

```sh
./bin/app
```
