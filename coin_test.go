package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCoinGenerate(t *testing.T) {
	var coin Coin
	err := coin.Generate(network["btc"])
	assert.Nil(t, err, "The `err` should not be `nil`")
	err = coin.Generate(network["test"])
	assert.NotNil(t, err, "The `err` should be `nil`")
}

func TestCoinImport(t *testing.T) {
	coin1 := Coin{Symbol: "btc", WIF: "1234"}
	err := coin1.Import(network[coin1.Symbol])
	assert.NotNil(t, err.Error(), "The `err` should not be `nil`")
	coin2 := Coin{Symbol: "test", WIF: "1234"}
	err = coin2.Import(network[coin2.Symbol])
	assert.NotNil(t, err.Error(), "The `err` should not be `nil`")
	coin3 := Coin{Symbol: "btc", WIF: "6vNUWjwJewkaC8TYhNSVa6nZg86x5eQUAgkoPs84YXBFVHemNSk"}
	err = coin3.Import(network[coin3.Symbol])
	assert.NotNil(t, err.Error(), "The `err` should not be `nil`")
	coin4 := Coin{Symbol: "ltc", WIF: "6vNUWjwJewkaC8TYhNSVa6nZg86x5eQUAgkoPs84YXBFVHemNSk"}
	err = coin4.Import(network[coin4.Symbol])
	assert.Nil(t, err, "The `err` should not be `nil`")
}

func TestCoinValidity(t *testing.T) {
	var coin1 Coin
	coin1.Generate(network["ltc"])
	coin2 := Coin{Symbol: "ltc", WIF: coin1.WIF}
	coin2.Import(network[coin2.Symbol])
	assert.Equal(t, coin1, coin2, "The two coins should match")
}
