package viacep

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"regexp"
)

const url = "https://viacep.com.br/ws/%s/json/"

func FindCEP(ctx context.Context, cep string) ([]byte, error) {
	if !isValidCEP(cep) {
		return nil, errors.New("invalid cep")
	}
	r, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf(url, cep), http.NoBody)
	if err != nil {
		return nil, err
	}
	response, err := http.DefaultClient.Do(r)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode == http.StatusBadRequest {
		return nil, errors.New("invalid cep")
	}
	result := new(AddressInformation)
	if err = json.NewDecoder(response.Body).Decode(result); err != nil {
		return nil, err
	}
	if result.Error != nil && *result.Error {
		return nil, errors.New("non-existent cep")
	}
	return json.Marshal(result)
}

func isValidCEP(cep string) bool {
	r, _ := regexp.Compile("[0-9]{5}-[0-9]{3}")
	return r.Match([]byte(cep))
}
