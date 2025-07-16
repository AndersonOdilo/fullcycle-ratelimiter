package entity

import "context"

type ClienteRepositoryInterface interface {

	Obtem(ctx context.Context, chave string) (Cliente, error)

	Grava(ctx context.Context, cliente Cliente) error
}

 
