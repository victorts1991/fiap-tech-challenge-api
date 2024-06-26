// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"fiap-tech-challenge-api/internal/adapters/http"
	"fiap-tech-challenge-api/internal/adapters/http/handlers"
	repository2 "fiap-tech-challenge-api/internal/adapters/repository"
	"fiap-tech-challenge-api/internal/core/usecase"
	"github.com/rhuandantas/fiap-tech-challenge-commons/pkg/db/mysql"
	"github.com/rhuandantas/fiap-tech-challenge-commons/pkg/middlewares/auth"
	"github.com/rhuandantas/fiap-tech-challenge-commons/pkg/util"
)

// Injectors from wire.go:

func InitializeWebServer() (*http.Server, error) {
	healthCheck := handlers.NewHealthCheck()
	dbConnector := repository.NewMySQLConnector()
	clienteRepo := repository2.NewClienteRepo(dbConnector)
	cadastrarClienteUseCase := usecase.NewCadastraCliente(clienteRepo)
	pesquisarClientePorCPF := usecase.NewPesquisarCliente(clienteRepo)
	lgpd := usecase.NewLGPD(clienteRepo)
	validator := util.NewCustomValidator()
	token := auth.NewJwtToken()
	cliente := handlers.NewCliente(cadastrarClienteUseCase, pesquisarClientePorCPF, lgpd, validator, token)
	produtoRepo := repository2.NewProdutoRepo(dbConnector)
	cadastrarProduto := usecase.NewCadastraProduto(produtoRepo)
	pegarProdutoPorCategoria := usecase.NewPegaProdutoPorCategoria(produtoRepo)
	apagarProduto := usecase.NewApagaProduto(produtoRepo)
	atualizarProduto := usecase.NewAtualizaProduto(produtoRepo)
	pegaPorIDS := usecase.NewPegaPorIDS(produtoRepo)
	produto := handlers.NewProduto(validator, cadastrarProduto, pegarProdutoPorCategoria, apagarProduto, atualizarProduto, pegaPorIDS)
	login := handlers.NewLogin(pesquisarClientePorCPF, token)
	server := http.NewAPIServer(healthCheck, cliente, produto, login)
	return server, nil
}
