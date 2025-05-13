package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type productsResponse struct {
	NextPageURL *string   `json:"next_page_url"`
	RegistryURL string    `json:"registry_url"`
	Products    []Product `json:"products"`
}

type Product struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	Repository string `json:"repository"`
}

func (c *Client) ListProducts() ([]Product, error) {
	products := []Product{}
	nextPageURL := fmt.Sprintf("%s/install/%s/products", c.BaseURL, c.LicenseKey)

	for {
		request, err := http.NewRequest("GET", nextPageURL, nil)
		if err != nil {
			return products, err
		}

		request.Header.Set("Accept", "application/json")

		response, err := c.HTTPClient.Do(request)
		if err != nil {
			return products, err
		}
		defer response.Body.Close()

		if response.StatusCode != http.StatusOK {
			return products, fmt.Errorf("Unexpected status code: %s", response.Status)
		}

		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return products, err
		}

		fmt.Printf("DEBUG: Response body: %s\n", body)

		var productsResponse productsResponse
		if err := json.Unmarshal(body, &productsResponse); err != nil {
			return products, err
		}

		if len(productsResponse.Products) == 0 {
			break
		}

		products = append(products, productsResponse.Products...)

		if productsResponse.NextPageURL == nil {
			break
		}
	}

	return products, nil
}
