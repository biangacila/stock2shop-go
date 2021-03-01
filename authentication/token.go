package authentication

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"github.com/biangacila/luvungula-go/global"
	"github.com/robbert229/jwt"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"log"
	"stock2shop-go/domain"
	"stock2shop-go/constant"
)

var VerifyKey, SignKey []byte
var TokensList map[string]domain.User

func init() {
	TokensList = make(map[string]domain.User)
	generateKey()
	var err error

	SignKey, err = ioutil.ReadFile(constant.PRIVATE_KEY_PATH)
	if err != nil {
		log.Fatal("Error reading private key")
		return
	}

	VerifyKey, err = ioutil.ReadFile(constant.PUBLIC_KEY_PATH)
	if err != nil {
		log.Fatal("Error reading public key")
		return
	}
}

func IsValidToken(token string) (bool, error) {
	var TOKEN_KEY = string(SignKey)
	algorithm := jwt.HmacSha256(TOKEN_KEY)
	_, err := algorithm.Decode(token)
	if err != nil {
		global.DisplayObject("Error Decode Token", err)
		return false, err
	}
	if algorithm.Validate(token) != nil {
		global.DisplayObject("Error Decode Token Expire", err)
		return false, errors.New("Expire")
	}
	return true, nil
}

func generateKey() {
	savePrivateFileTo := constant.PRIVATE_KEY_PATH
	savePublicFileTo := constant.PUBLIC_KEY_PATH
	bitSize := 4096
	privateKey, err := generatePrivateKey(bitSize)
	if err != nil {
		log.Fatal(err.Error())
	}
	publicKeyBytes, err := generatePublicKey(&privateKey.PublicKey)
	if err != nil {
		log.Fatal(err.Error())
	}
	privateKeyBytes := encodePrivateKeyToPEM(privateKey)
	err = writeKeyToFile(privateKeyBytes, savePrivateFileTo)
	if err != nil {
		log.Fatal(err.Error())
	}
	err = writeKeyToFile([]byte(publicKeyBytes), savePublicFileTo)
	if err != nil {
		log.Fatal(err.Error())
	}
}

// generatePrivateKey creates a RSA Private Key of specified byte size
func generatePrivateKey(bitSize int) (*rsa.PrivateKey, error) {
	// Private Key generation
	privateKey, err := rsa.GenerateKey(rand.Reader, bitSize)
	if err != nil {
		return nil, err
	}
	// Validate Private Key
	err = privateKey.Validate()
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}

// encodePrivateKeyToPEM encodes Private Key from RSA to PEM format
func encodePrivateKeyToPEM(privateKey *rsa.PrivateKey) []byte {
	// Get ASN.1 DER format
	privDER := x509.MarshalPKCS1PrivateKey(privateKey)
	// pem.Block
	privBlock := pem.Block{
		Type:    "RSA PRIVATE KEY",
		Headers: nil,
		Bytes:   privDER,
	}
	// Private key in PEM format
	privatePEM := pem.EncodeToMemory(&privBlock)

	return privatePEM
}

// generatePublicKey take a rsa.PublicKey and return bytes suitable for writing to .pub file
// returns in the format "ssh-rsa ..."
func generatePublicKey(privatekey *rsa.PublicKey) ([]byte, error) {
	publicRsaKey, err := ssh.NewPublicKey(privatekey)
	if err != nil {
		return nil, err
	}
	pubKeyBytes := ssh.MarshalAuthorizedKey(publicRsaKey)
	log.Println("Public key generated")
	return pubKeyBytes, nil
}

// writePemToFile writes keys to a file
func writeKeyToFile(keyBytes []byte, saveFileTo string) error {
	err := ioutil.WriteFile(saveFileTo, keyBytes, 0600)
	if err != nil {
		return err
	}
	return nil
}
