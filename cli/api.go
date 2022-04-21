package cli

import (
	"encoding/json"
	"net/http"
)

type ApiResponse interface {
}

func fetchJSON[T ApiResponse](url string, hc http.Client, data *T) error {
	resp, err := hc.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(data)
	if err != nil {
		return err
	}
	return nil
}
