package apicep

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

const url = "https://cdn.apicep.com/file/apicep/%s.json"

func FindCEP(ctx context.Context, cep string) ([]byte, error) {
	r, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf(url, cep), http.NoBody)
	if err != nil {
		return nil, err
	}
	response, err := http.DefaultClient.Do(r)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	decoder := json.NewDecoder(response.Body)
	if response.StatusCode != http.StatusOK {
		errorResponse := new(Error)
		if err = decoder.Decode(errorResponse); err != nil {
			return nil, err
		}
		return nil, errors.New(errorResponse.Message)
	}
	return io.ReadAll(response.Body)
}
