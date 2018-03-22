package main

import (
	"errors"

	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/wire"
	"github.com/btcsuite/btcutil"
)

type Coin struct {
	Name                string `json:"name"`
	Symbol              string `json:"symbol"`
	WIF                 string `json:"wif,omitempty"`
	UncompressedAddress string `json:"uncompressed_address"`
	CompressedAddress   string `json:"compressed_address"`
}

type Network struct {
	name        string
	symbol      string
	xpubkey     byte
	xprivatekey byte
	magic       wire.BitcoinNet
}

var network = map[string]Network{
	"rdd":  {name: "reddcoin", symbol: "rdd", xpubkey: 0x3d, xprivatekey: 0xbd, magic: 0xfbc0b6db},
	"dgb":  {name: "digibyte", symbol: "dgb", xpubkey: 0x1e, xprivatekey: 0x80, magic: 0xfac3b6da},
	"btc":  {name: "bitcoin", symbol: "btc", xpubkey: 0x00, xprivatekey: 0x80, magic: 0xf9beb4d9},
	"ltc":  {name: "litecoin", symbol: "ltc", xpubkey: 0x30, xprivatekey: 0xb0, magic: 0xfbc0b6db},
	"dash": {name: "dash", symbol: "dash", xpubkey: 0x4c, xprivatekey: 0xcc, magic: 0xd9b4bef9},
}

func (network Network) GetNetworkParams() *chaincfg.Params {
	networkParams := &chaincfg.MainNetParams
	networkParams.Name = network.name
	networkParams.Net = network.magic
	networkParams.PubKeyHashAddrID = network.xpubkey
	networkParams.PrivateKeyID = network.xprivatekey
	return networkParams
}

func (coin *Coin) Generate(network Network) error {
	if network == (Network{}) {
		return errors.New("Unsupported cryptocurrency symbol provided")
	}
	secret, err := btcec.NewPrivateKey(btcec.S256())
	if err != nil {
		return err
	}
	wif, err := btcutil.NewWIF(secret, network.GetNetworkParams(), false)
	if err != nil {
		return err
	}
	coin.WIF = wif.String()
	uncompressedAddress, err := btcutil.NewAddressPubKey(wif.PrivKey.PubKey().SerializeUncompressed(), network.GetNetworkParams())
	if err != nil {
		return err
	}
	compressedAddress, err := btcutil.NewAddressPubKey(wif.PrivKey.PubKey().SerializeCompressed(), network.GetNetworkParams())
	if err != nil {
		return err
	}
	coin.UncompressedAddress = uncompressedAddress.EncodeAddress()
	coin.CompressedAddress = compressedAddress.EncodeAddress()
	coin.Name = network.name
	coin.Symbol = network.symbol
	return nil
}

func (coin *Coin) Import(network Network) error {
	if network == (Network{}) {
		return errors.New("Unsupported cryptocurrency symbol provided")
	}
	wif, err := btcutil.DecodeWIF(coin.WIF)
	if err != nil {
		return err
	}
	if !wif.IsForNet(network.GetNetworkParams()) {
		return errors.New("The WIF string is not valid for the `" + network.name + "` network")
	}
	uncompressedAddress, err := btcutil.NewAddressPubKey(wif.PrivKey.PubKey().SerializeUncompressed(), network.GetNetworkParams())
	if err != nil {
		return err
	}
	compressedAddress, err := btcutil.NewAddressPubKey(wif.PrivKey.PubKey().SerializeCompressed(), network.GetNetworkParams())
	if err != nil {
		return err
	}
	coin.WIF = wif.String()
	coin.UncompressedAddress = uncompressedAddress.EncodeAddress()
	coin.CompressedAddress = compressedAddress.EncodeAddress()
	coin.Name = network.name
	coin.Symbol = network.symbol
	return nil
}
