package config

import (
	"crypto/sha1"
	"encoding/hex"
	"strings"
)

type ProductConfig struct {
	AuthServer   string `json:"auth_server"`
	LicenseKey   string `json:"license_key"`
	EmailAddress string `json:"email_address"`
	Domain       string `json:"domain"`
	HTTPS        bool   `json:"https"`
	Product      string `json:"product"`
	Registry     string `json:"registry"`
	Repository   string `json:"repository"`
}

func NewProductConfig(authServer string, licenseKey string) *ProductConfig {
	return &ProductConfig{
		AuthServer: authServer,
		LicenseKey: licenseKey,
	}
}

func (p *ProductConfig) ContainerName() string {
	h := sha1.New()
	h.Write([]byte(p.AuthServer))
	h.Write([]byte(p.LicenseKey))
	// h.Write([]byte(p.EmailAddress))
	// h.Write([]byte(p.Domain))
	// h.Write([]byte(p.Product))
	// h.Write([]byte(p.Registry))
	// h.Write([]byte(p.Repository))

	full := h.Sum(nil)
	digest := hex.EncodeToString(full)[:8]

	imageName := p.Repository
	if strings.Contains(p.Repository, "/") {
		parts := strings.Split(p.Repository, "/")
		imageName = parts[len(parts)-1]
	}

	return imageName + "-" + digest
}
