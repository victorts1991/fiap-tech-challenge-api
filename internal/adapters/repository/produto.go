package repository

import (
	"context"
	"fiap-tech-challenge-api/internal/core/commons"
	"fiap-tech-challenge-api/internal/core/domain"
	"fmt"
	"github.com/joomcode/errorx"
	"xorm.io/xorm"
)

const tableNameProduto string = "produto"

type produto struct {
	session *xorm.Session
}

type ProdutoRepo interface {
	Insere(ctx context.Context, produto *domain.Produto) (*domain.Produto, error)
	PesquisaPorCategoria(ctx context.Context, produto *domain.Produto) ([]*domain.Produto, error)
	PesquisaPorID(ctx context.Context, produto *domain.Produto) (*domain.Produto, error)
	Apaga(ctx context.Context, produto *domain.Produto) error
	Atualiza(ctx context.Context, produto *domain.Produto, id int64) error
}

func NewProdutoRepo(connector DBConnector) ProdutoRepo {
	session := connector.GetORM().Table(tableNameProduto)
	return &produto{
		session: session,
	}
}

func (p *produto) Insere(ctx context.Context, produto *domain.Produto) (*domain.Produto, error) {
	_, err := p.session.Context(ctx).Insert(produto)
	if err != nil {
		if commons.IsDuplicatedEntryError(err) {
			return nil, commons.BadRequest.New("produto já existe")
		}

		return nil, err
	}

	return produto, nil
}

func (p *produto) PesquisaPorCategoria(ctx context.Context, filter *domain.Produto) ([]*domain.Produto, error) {
	produtos := make([]*domain.Produto, 0)
	err := p.session.Context(ctx).Find(&produtos, filter)
	if err != nil {
		return nil, err
	}

	return produtos, nil
}

func (p *produto) PesquisaPorID(ctx context.Context, produto *domain.Produto) (*domain.Produto, error) {
	has, err := p.session.Context(ctx).Get(produto)
	if err != nil {
		return nil, err
	}

	if !has {
		return nil, commons.NotFound.New("produto não encontrado")
	}

	return produto, nil
}

func (p *produto) Apaga(ctx context.Context, produto *domain.Produto) error {
	_, err := p.PesquisaPorID(ctx, produto)
	if err != nil {
		return err
	}

	_, err = p.session.Context(ctx).ID(produto.Id).Delete(&domain.Produto{})
	if err != nil {
		return errorx.InternalError.New(fmt.Sprintf("falha ao deletar produto: %s", err.Error()))
	}

	return nil
}

func (p *produto) Atualiza(ctx context.Context, new *domain.Produto, id int64) error {
	_, err := p.session.Context(ctx).ID(id).Update(new)
	if err != nil {
		return errorx.InternalError.New(err.Error())
	}

	return nil
}
