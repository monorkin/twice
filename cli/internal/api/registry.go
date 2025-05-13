package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (c *Client) RegistryURL() (string, error) {
	url := fmt.Sprintf("%s/install/%s/products", c.BaseURL, c.LicenseKey)

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	request.Header.Set("Accept", "application/json")

	response, err := c.HTTPClient.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Unexpected status code: %s", response.Status)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	var productsResponse productsResponse
	if err := json.Unmarshal(body, &productsResponse); err != nil {
		return "", err
	}

	return productsResponse.RegistryURL, nil
}
