package pvesh

import (
	"crypto/x509"
)

type ProxmoxCertInfo struct {
	Filename      string   `json:"filename,omitempty"`
	Fingerprint   string   `json:"fingerprint,omitempty"`
	Issuer        string   `json:"issuer,omitempty"`
	Notafter      int      `json:"notafter,omitempty"`
	Notbefore     int      `json:"notbefore,omitempty"`
	Pem           string   `json:"pem,omitempty"`
	PublicKeyBits int      `json:"public-key-bits,omitempty"`
	PublicKeyType string   `json:"public-key-type,omitempty"`
	San           []string `json:"san,omitempty"`
	Subject       string   `json:"subject,omitempty"`
}

// Certificate returns x509 certificate instance from proxmox cert
func (crt ProxmoxCertInfo) Certificate() (*x509.Certificate, error) {
	data := []byte(crt.Pem)
	return x509.ParseCertificate(data)
}
