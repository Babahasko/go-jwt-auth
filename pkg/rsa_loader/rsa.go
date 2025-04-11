package rsa_loader

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

type RSALoader struct {
    PrivateKey *rsa.PrivateKey
    PublicKey *rsa.PublicKey
}

func NewRSA(privateKeyFile, publicKeyFile string) (*RSALoader, error) {
    loader := &RSALoader{}
    privateKey, err := loader.loadPrivateKeyFromFile(privateKeyFile)
    if err != nil {
        return nil, fmt.Errorf("failed to load private key: %w", err)
    }
    loader.PrivateKey = privateKey
    publicKey, err := loader.loadPublicKeyFromFile(publicKeyFile)
    if err != nil {
        return nil, fmt.Errorf("failed to load public key: %w", err)
    }
    loader.PublicKey = publicKey

    return loader, nil
}
// Загрузка приватного ключа из PEM-файла
func (loader *RSALoader) loadPrivateKeyFromFile(filename string) (*rsa.PrivateKey, error) {
    data, err := os.ReadFile(filename)
    if err != nil {
        return nil, fmt.Errorf("failed to read private key file: %w", err)
    }

    block, _ := pem.Decode(data)
    if block == nil || block.Type != "RSA PRIVATE KEY" {
        return nil, fmt.Errorf("failed to decode PEM block containing private key")
    }

    privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
    if err != nil {
        return nil, fmt.Errorf("failed to parse private key: %w", err)
    }
    return privateKey, nil
}

// Загрузка публичного ключа из PEM-файла
func (loader *RSALoader)loadPublicKeyFromFile(filename string) (*rsa.PublicKey, error) {
    data, err := os.ReadFile(filename)
    if err != nil {
        return nil, fmt.Errorf("failed to read public key file: %w", err)
    }

    block, _ := pem.Decode(data)
    if block == nil || block.Type != "PUBLIC KEY" {
        return nil, fmt.Errorf("failed to decode PEM block containing public key")
    }

    pub, err := x509.ParsePKIXPublicKey(block.Bytes)
    if err != nil {
        return nil, fmt.Errorf("failed to parse public key: %w", err)
    }

    publicKey, ok := pub.(*rsa.PublicKey)
    if !ok {
        return nil, fmt.Errorf("parsed public key is not an RSA key")
    }
    return publicKey, nil
}