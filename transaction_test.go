package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTransaction(t *testing.T) {
	transaction, err := CreateTransaction(network["btc"], "5HusYj2b2x4nroApgfvaSfKYZhRbKFH41bVyPooymbC6KfgSXdD", "1KKKK6N21XKo48zWKuQKXdvSsCf95ibHFa", 91234, "81b4c832d70cb56ff957589752eb4125a4cab78a25a8fc52d6a09e5bd4404d48")
	assert.Nil(t, err, "The `err` should be `nil`")
	assert.Equal(t, transaction.SourceAddress, "1MMMMSUb1piy2ufrSguNUdFmAcvqrQF8M5", "The `SourceAddress` does not match")
	assert.Equal(t, transaction.DestinationAddress, "1KKKK6N21XKo48zWKuQKXdvSsCf95ibHFa", "The `DestinationAddress` does not match")
	assert.Equal(t, transaction.TxId, "4e8378675bcf6a389c8cfe246094aafa44249e48ab88a40e6fda3bf0f44f916a", "The `TxId` does not match")
	assert.Equal(t, transaction.UnsignedTx, "0100000001484d40d45b9ea0d652fca8258ab7caa42541eb52975857f96fb50cd732c8b4810000000000ffffffff0162640100000000001976a914df3bd30160e6c6145baaf2c88a8844c13a00d1d588ac00000000", "The `UnsignedTx` does not match")
	assert.Equal(t, transaction.SignedTx, "01000000016a914ff4f03bda6f0ea488ab489e2444faaa946024fe8c9c386acf5b6778834e000000008b483045022100904dbeddeecccf6391ac92922381ae006bf244c002f42e195daa0a9837a4b5820220461677f9dbb7d9580e268ac486cfeb4b9d87bfdd6d4e7b1be09b8e6f5cc0a70701410414e301b2328f17442c0b8310d787bf3d8a404cfbd0704f135b6ad4b2d3ee751310f981926e53a6e8c39bd7d3fefd576c543cce493cbac06388f2651d1aacbfcdffffffff0162640100000000001976a914c8e90996c7c6080ee06284600c684ed904d14c5c88ac00000000", "The `SignedTx` does not match")
	assert.Equal(t, transaction.Amount, int64(91234), "The `Amount` does not match")
	transaction = Transaction{}
	transaction, err = CreateTransaction(network["btc"], "6KBco75idUx4umGRdrSmZnNKPgW4UuGMNFnvLL6JqaBVA7ka7Y", "1KKKK6N21XKo48zWKuQKXdvSsCf95ibHFa", 91234, "81b4c832d70cb56ff957589752eb4125a4cab78a25a8fc52d6a09e5bd4404d48")
	assert.NotNil(t, err, "The `err` should not be `nil`")
}
