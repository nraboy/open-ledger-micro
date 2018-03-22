package main

import (
	"fmt"
	"testing"

	"github.com/btcsuite/btcutil"
	"github.com/stretchr/testify/assert"
)

func TestWalletCipher(t *testing.T) {
	var wallet Wallet
	passphrase := "1234"
	defer wallet.Destroy()
	wallet.Create(passphrase)
	err := wallet.EncryptFile(passphrase)
	assert.Nil(t, err, "The `err` should be `nil`")
	err = wallet.DecryptFile(passphrase)
	assert.Nil(t, err, "The `err` should be `nil`")
	err = wallet.DecryptFile("testing")
	assert.NotNil(t, err.Error(), "The `err` should not be `nil`")
}

func TestWalletImport(t *testing.T) {
	var wallet Wallet
	passphrase := "1234"
	defer wallet.Destroy()
	wallet.Create(passphrase)
	var coin Coin
	err := wallet.Import(coin, passphrase)
	assert.NotNil(t, err.Error(), "The `err` should not be `nil`")
	coin.Generate(network["btc"])
	err = wallet.Import(coin, passphrase)
	assert.Nil(t, err, "The `err` should not be `nil`")
}

func TestWalletAddresses(t *testing.T) {
	var wallet Wallet
	passphrase := "1234"
	defer wallet.Destroy()
	wallet.Create(passphrase)
	var coin Coin
	coin.Generate(network["btc"])
	wallet.Import(coin, passphrase)
	wallet.GetAddresses(passphrase)
	assert.Equal(t, wallet.Coins[0].WIF, "", "The `wif` should not be present")
	wallet.Dump(passphrase)
}

func TestAddress(t *testing.T) {
	decAddr, err := btcutil.DecodeAddress("1KKKK6N21XKo48zWKuQKXdvSsCf95ibHFa", network["btc"].GetNetworkParams())
	if err != nil {
		fmt.Printf("DecodeAddress: %v\n", err)
	}
	pkh, ok := decAddr.(*btcutil.AddressPubKeyHash)
	if !ok {
		fmt.Printf("invalid type: %T\n", pkh)
	}

	if !pkh.IsForNet(network["btc"].GetNetworkParams()) {
		fmt.Printf("address not for %s\n", network["btc"].GetNetworkParams().Name)
	}
}
