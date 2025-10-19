package catapi

import (
	"net/http"
	"encoding/json"
	"io"
)

type factAPIResponse struct {
	Fact string `json:"fact"`
}


func GetCatFacts(url string) (factAPIResponse, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return factAPIResponse{}, err
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return factAPIResponse{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return factAPIResponse{}, err
	}

	var factapires factAPIResponse
	if err := json.Unmarshal(body, &factapires); err != nil {
		return factAPIResponse{}, err
	}

	return factapires, nil
}