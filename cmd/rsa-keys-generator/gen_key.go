package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

// Генерация пары ключей RSA
func generateRSAKeys() (*rsa.PrivateKey, *rsa.PublicKey, error) {
    privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
    if err != nil {
        return nil, nil, fmt.Errorf("failed to generate private key: %w", err)
    }
    publicKey := &privateKey.PublicKey
    return privateKey, publicKey, nil
}
// Сохранение приватного ключа в PEM-формате
func savePrivateKeyToFile(privateKey *rsa.PrivateKey, filename string) error {
    privateKeyPEM := &pem.Block{
        Type:  "RSA PRIVATE KEY",
        Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
    }

    file, err := os.Create(filename)
    if err != nil {
        return fmt.Errorf("failed to create private key file: %w", err)
    }
    defer file.Close()

    if err := pem.Encode(file, privateKeyPEM); err != nil {
        return fmt.Errorf("failed to write private key to file: %w", err)
    }
    return nil
}

// Сохранение публичного ключа в PEM-формате
func savePublicKeyToFile(publicKey *rsa.PublicKey, filename string) error {
    publicKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
    if err != nil {
        return fmt.Errorf("failed to marshal public key: %w", err)
    }

    publicKeyPEM := &pem.Block{
        Type:  "PUBLIC KEY",
        Bytes: publicKeyBytes,
    }

    file, err := os.Create(filename)
    if err != nil {
        return fmt.Errorf("failed to create public key file: %w", err)
    }
    defer file.Close()

    if err := pem.Encode(file, publicKeyPEM); err != nil {
        return fmt.Errorf("failed to write public key to file: %w", err)
    }
    return nil
}

// Генерация пары ключей RSA
func main() {
    // Генерация ключей RSA
    privateKey, publicKey, err := generateRSAKeys()
    if err != nil {
        fmt.Println("Error generating RSA keys:", err)
        return
    }

	// Сохранение ключей в файлы
    err = savePrivateKeyToFile(privateKey, "private_key.pem"); 
	if err != nil {
        fmt.Println("Error saving private key:", err)
        return
    }

	err = savePublicKeyToFile(publicKey, "pub_key.pem")

	if err != nil {
        fmt.Println("Error saving private key:", err)
        return
    }
}