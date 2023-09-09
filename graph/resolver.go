package graph

import "github.com/danyukod/chave-pix-bff/internal/application/port/input"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	input.FindPixKeyUsecase
	input.FindPixKeyByKeyUsecase
	input.CreatePixKeyUsecase
}
