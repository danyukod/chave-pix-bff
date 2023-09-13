package pix_key

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/danyukod/chave-pix-bff/internal/application/domain/model"
	"io"
	"net/http"
)

type PixKeyClient struct {
}

func NewPixKeyClient() *PixKeyClient {
	return &PixKeyClient{}
}

func (p *PixKeyClient) FindPixKeyByKey(key string) (model.PixKeyDomainInterface, error) {
	resp, err := http.Get("http://localhost:8082/api/v1/pix-keys/" + key)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	var dto PixKeyResponseDTO

	err = json.Unmarshal(body, &dto)
	if err != nil {
		return nil, err
	}

	domain, err := dto.ToDomain()
	if err != nil {
		return nil, err
	}

	return domain, nil
}

func (p *PixKeyClient) FindPixKey() ([]model.PixKeyDomainInterface, error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://localhost:8080/api/v1/pix-keys", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTQxNDA3OTMsInN1YiI6ImZjY2JmZWFiLWIzM2QtNGRkYS1iZjhhLThjMGRjNmE4OTNjMSJ9.eVkFJXOik6efeQZ996sXXbX30O3BZ_EsEVV5M9wQX6C3agjuHubRQrqiVMh_b9xqSIF0wN8j7vFNzWmB5mezlw")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	var dto []PixKeyResponseDTO

	err = json.Unmarshal(body, &dto)
	if err != nil {
		return nil, err
	}

	var domains []model.PixKeyDomainInterface

	for _, pixKeyResponseDTO := range dto {
		domain, err := pixKeyResponseDTO.ToDomain()
		if err != nil {
			return nil, err
		}
		domains = append(domains, domain)
	}

	return domains, nil
}

func (p *PixKeyClient) CreatePixKey(pixKey model.PixKeyDomainInterface) (model.PixKeyDomainInterface, error) {
	client := &http.Client{}

	body := NewPixKeyRequestDTOFromDomain(pixKey)

	jsonBody, err := json.Marshal(body)

	if err != nil {
		return nil, err
	}

	req, _ := http.NewRequest("POST", "http://localhost:8080/api/v1/pix-keys", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTQxNDA3OTMsInN1YiI6ImZjY2JmZWFiLWIzM2QtNGRkYS1iZjhhLThjMGRjNmE4OTNjMSJ9.eVkFJXOik6efeQZ996sXXbX30O3BZ_EsEVV5M9wQX6C3agjuHubRQrqiVMh_b9xqSIF0wN8j7vFNzWmB5mezlw")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)

	if resp.StatusCode != 201 {
		err := errors.New("error creating pix key")
		return nil, err
	}

	var dto PixKeyResponseDTO

	err = json.NewDecoder(resp.Body).Decode(&dto)

	if err != nil {
		return nil, err
	}

	domain, err := dto.ToDomain()
	if err != nil {
		return nil, err
	}

	return domain, nil
}
