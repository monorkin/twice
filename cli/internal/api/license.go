package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type License struct {
	Key     string  `json:"key"`
	Owner   Owner   `json:"owner"`
	Product Product `json:"product"`
}

type Product struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	Registry   string `json:"registry"`
	Repository string `json:"repository"`
}

type Owner struct {
	EmailAddress string `json:"email_address"`
}

func (c *Client) InspectLicense(licenseKey string) (*License, error) {
	url := fmt.Sprintf("%s/install/%s", c.BaseURL, licenseKey)

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
	case http.StatusOK:
	case http.StatusNotFound:
		return nil, nil
	default:
		return nil, fmt.Errorf("License inspection returned unexpected status code: %s", response.Status)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, errors.Join(errors.New("Couldn't to read license response"), err)
	}

	license := &License{}
	if err := json.Unmarshal(body, license); err != nil {
		return nil, errors.Join(errors.New("Couldn't to unmarshal license response"), err)
	}

	return license, nil
}
