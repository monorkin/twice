package config

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

type Product struct {
	AuthServer      string `json:"auth_server"`
	LicenseKey      string `json:"license_key"`
	EmailAddress    string `json:"email_address"`
	Domain          string `json:"domain"`
	HTTPS           bool   `json:"https"`
	Product         string `json:"product"`
	Registry        string `json:"registry"`
	Repository      string `json:"repository"`
	VAPIDPublicKey  string `json:"vapid_public_key"`
	VAPIDPrivateKey string `json:"vapid_private_key"`
	SecretKeyBase   string `json:"secret_key_base"`
}

func NewProduct(authServer string, licenseKey string) *Product {
	product := &Product{
		AuthServer: authServer,
		LicenseKey: licenseKey,
	}

	if err := product.generateSecretKeyBase(); err != nil {
		panic(errors.Wrap(err, "failed to generate secret key base"))
	}

	if err := product.generateVAPIDKeys(); err != nil {
		panic(errors.Wrap(err, "failed to generate VAPID keys"))
	}

	return product
}

func (p *Product) ID() string {
	return p.ContainerName()
}

func (p *Product) ContainerName() string {
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

func (p *Product) Image() string {
	image := fmt.Sprintf("%s/%s", p.Registry, p.Repository)

	if !strings.Contains(image, ":") {
		image = fmt.Sprintf("%s:latest", image)
	}

	return image
}

func (p *Product) generateSecretKeyBase() error {
	key := make([]byte, 64)
	if _, err := rand.Read(key); err != nil {
		return errors.Wrap(err, "failed to generate secret key base")
	}

	p.SecretKeyBase = hex.EncodeToString(key)

	return nil
}

func (p *Product) generateVAPIDKeys() error {
	priv, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return errors.Wrap(err, "failed to generate VAPID keys")
	}

	p.VAPIDPrivateKey = base64.RawURLEncoding.EncodeToString(priv.D.Bytes())

	x := priv.X.Bytes()
	y := priv.Y.Bytes()
	pubBytes := append([]byte{0x04}, append(x, y...)...)

	p.VAPIDPublicKey = base64.RawURLEncoding.EncodeToString(pubBytes)

	return nil
}
