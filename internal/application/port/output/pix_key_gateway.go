package output

import (
	"github.com/danyukod/chave-pix-bff/internal/application/domain/model"
	"github.com/danyukod/chave-pix-bff/internal/application/queries/dto"
)

type PixKeyGateway interface {
	FindPixKey() ([]dto.FindPixKeyQueryDTO, error)
	FindPixKeyByKey(key string) (model.PixKeyDomainInterface, error)
	CreatePixKey(pixKey model.PixKeyDomainInterface) (model.PixKeyDomainInterface, error)
}
