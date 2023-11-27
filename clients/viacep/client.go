package viacep

import (
	"context"
	"errors"
	"io"
	"net/http"
	"regexp"
)

func FindCEP(ctx context.Context, cep string) (*AddressInformation, error) {
	if !isValidCEP(cep) {
		return nil, errors.New("invalid cep")
	}
	r, err := http.NewRequestWithContext(ctx, http.MethodGet, "", http.NoBody)
	if err != nil {
		return nil, err
	}
	response, err := http.DefaultClient.Do(r)
	if err != nil {
		return nil, err
	}
	defer func(body io.ReadCloser) {
		err := body.Close()
		if err != nil {
			panic(err)
		}
	}(response.Body)
	// continuar
	return nil, nil
}

func isValidCEP(cep string) bool {
	r, _ := regexp.Compile("[0-9]{5}-[0-9]{3}")
	return r.Match([]byte(cep))
}
