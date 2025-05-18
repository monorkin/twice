package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

type Config struct {
	Products []ProductConfig `json:"products"`
}

func LoadOrCreateConfig() (*Config, error) {
	config, err := LoadConfig()
	if err != nil {
		config = NewConfig()

		_, err := config.Save()
		if err != nil {
			return nil, errors.Wrap(err, "failed to save new config")
		}
	}

	return config, nil
}

func NewConfig() *Config {
	return &Config{}
}

func LoadConfig() (*Config, error) {
	path, err := ConfigFilePath()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get config file path")
	}

	return LoadConfigFrom(path)
}

func LoadConfigFrom(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't to open config file")
	}
	defer file.Close()

	var config Config
	if err := json.NewDecoder(file).Decode(&config); err != nil {
		return nil, errors.Wrap(err, "failed to decode config file")
	}

	return &config, nil
}

func (c *Config) Save() (string, error) {
	path, err := ConfigFilePath()
	if err != nil {
		return "", errors.Wrap(err, "failed to get config file path")
	}

	return c.SaveTo(path)
}

func (c *Config) SaveTo(path string) (string, error) {
	dirPath := filepath.Dir(path)
	err := os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		return "", errors.Wrap(err, "failed to create config directory")
	}

	file, err := os.Create(path)
	if err != nil {
		return "", errors.Wrap(err, "failed to create config file")
	}
	defer file.Close()

	encoder := json.NewEncoder(file)

	// Make the file human-readable
	encoder.SetIndent("", "  ")

	if err := encoder.Encode(c); err != nil {
		return "", errors.Wrap(err, "failed to encode config to JSON")
	}

	return path, nil
}

func (c *Config) AddProduct(newProduct *ProductConfig) error {
	for _, product := range c.Products {
		if product.AuthServer == newProduct.AuthServer && product.LicenseKey == newProduct.LicenseKey {
			return errors.New("product with this license key already exists")
		}
	}

	c.Products = append(c.Products, *newProduct)

	return nil
}

func (c *Config) RemoveProduct(product *ProductConfig) error {
	for i, p := range c.Products {
		if product.AuthServer == p.AuthServer && product.LicenseKey == p.LicenseKey {
			c.Products = append(c.Products[:i], c.Products[i+1:]...)
			return nil
		}
	}

	return errors.New("product with this license key not found")
}

func (c *Config) UpdateProduct(product *ProductConfig) error {
	fmt.Printf("Updated Product: %+v\n", product)
	for i, p := range c.Products {
		if product.AuthServer == p.AuthServer && product.LicenseKey == p.LicenseKey {
			c.Products[i] = *product
			fmt.Printf("New Products: %+v\n", c.Products)
			return nil
		}
	}

	return errors.New("product with this license key not found")
}
