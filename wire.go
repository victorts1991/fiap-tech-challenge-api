// wire.go
//go:build wireinject

package main

import (
	"fiap-tech-challenge-api/internal/adapters/http"
	"fiap-tech-challenge-api/internal/adapters/http/handlers"
	"fiap-tech-challenge-api/internal/adapters/repository"
	"fiap-tech-challenge-api/internal/core/usecases"
	"fiap-tech-challenge-api/internal/util"
	"github.com/google/wire"
)

func InitializeWebServer() (*http.Server, error) {
	wire.Build(repository.NewMySQLConnector,
		util.NewCustomValidator,
		repository.NewClienteRepo,
		usecases.NewCadastraCliente,
		usecases.NewPesquisarClientePorCpf,
		handlers.NewCliente,
		handlers.NewHealthCheck,
		http.NewAPIServer)
	return &http.Server{}, nil
}
