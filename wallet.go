package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"os"
)

type Wallet struct {
	Coins []Coin `json:"coins"`
}

func (wallet Wallet) CreateHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

func (wallet Wallet) Encrypt(passphrase string) ([]byte, error) {
	block, err := aes.NewCipher([]byte(wallet.CreateHash(passphrase)))
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	nonce := make([]byte, gcm.NonceSize())
	io.ReadFull(rand.Reader, nonce)
	data, err := json.Marshal(wallet)
	if err != nil {
		return nil, err
	}
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext, nil
}

func (wallet *Wallet) Decrypt(data []byte, passphrase string) error {
	block, err := aes.NewCipher([]byte(wallet.CreateHash(passphrase)))
	if err != nil {
		return err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return err
	}
	json.Unmarshal(plaintext, wallet)
	return nil
}

func (wallet Wallet) EncryptFile(passphrase string) error {
	file, err := os.Create("wallet.dat")
	if err != nil {
		return err
	}
	defer file.Close()
	data, err := wallet.Encrypt(passphrase)
	if err != nil {
		return err
	}
	_, err = file.Write(data)
	return err
}

func (wallet *Wallet) DecryptFile(passphrase string) error {
	ciphertext, err := ioutil.ReadFile("wallet.dat")
	if err != nil {
		return err
	}
	err = wallet.Decrypt(ciphertext, passphrase)
	if err != nil {
		return err
	}
	return nil
}

func (wallet Wallet) Create(key string) error {
	err := wallet.EncryptFile(key)
	return err
}

func (wallet Wallet) Destroy() error {
	err := os.Remove("wallet.dat")
	return err
}

func (wallet Wallet) Import(coin Coin, passphrase string) error {
	if coin == (Coin{}) {
		return errors.New("The coin must be valid")
	}
	wallet.Coins = append(wallet.Coins, coin)
	wallet.EncryptFile(passphrase)
	return nil
}

func (wallet *Wallet) Dump(passphrase string) error {
	err := wallet.DecryptFile(passphrase)
	if err != nil {
		return errors.New(`{ "message": "The password is not correct" }`)
	}
	return nil
}

func (wallet *Wallet) GetAddresses(passphrase string) error {
	err := wallet.DecryptFile(passphrase)
	if err != nil {
		return errors.New(`{ "message": "The password is not correct" }`)
	}
	for index, _ := range wallet.Coins {
		wallet.Coins[index].WIF = ""
	}
	return nil
}

func (wallet *Wallet) Authenticate(passphrase string) bool {
	err := wallet.DecryptFile(passphrase)
	if err != nil {
		return false
	}
	return true
}
