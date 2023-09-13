package output

import "github.com/danyukod/chave-pix-bff/internal/application/domain/model"

type AuthGateway interface {
	GenerateToken(model.AuthDomainInterface) (string, error)
}
