package output

import "github.com/danyukod/chave-pix-bff/internal/application/domain/model"

type PixKeyGateway interface {
	FindPixKey() ([]model.PixKeyDomainInterface, error)
	FindPixKeyByKey(key string) (model.PixKeyDomainInterface, error)
}
