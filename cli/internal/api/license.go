package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type License struct {
	key     string  `json:"key"`
	owner   Owner   `json:"owner"`
	product Product `json:"product"`
}

type Product struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	Registry   string `json:"repository"`
	Repository string `json:"repository"`
}

type Owner struct {
	emailAddress string `json:"email_address"`
}

func (c *Client) InspectLicense(licenseKey) (*License, error) {
	url := fmt.Sprintf("%s/install/%s", c.BaseURL, c.LicenseKey)

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Accept", "application/json")

	response, err := c.HTTPClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	switch response.StatusCode {
	case http.StatusNotFound:
		return nil, nil
	default:
		return nil, fmt.Errorf("Unexpected status code: %s", response.Status)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var license License
	if err := json.Unmarshal(body, &license); err != nil {
		return nil, err
	}

	return license, nil
}
