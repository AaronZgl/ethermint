package ethsecp256k1

import (
	cryptoamino "github.com/tendermint/tendermint/crypto/encoding/amino"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/crypto/keys"
)

// CryptoCodec is the default amino codec used by ethermint
var CryptoCodec = codec.NewLegacyAminoLegacyAmino()

func init() {
	// replace the keyring codec with the ethermint crypto codec to prevent
	// amino panics because of unregistered Priv/PubKey
	keys.CryptoCdc = CryptoCodec
	keys.RegisterLegacyAminoCodec(CryptoCodec)
	cryptoamino.RegisterAmino(CryptoCodec)
	RegisterLegacyAminoCodec(CryptoCodec)
}

// RegisterLegacyAminoCodec registers all the necessary types with amino for the given
// codec.
func RegisterLegacyAminoCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(PubKey{}, PubKeyName, nil)
	cdc.RegisterConcrete(PrivKey{}, PrivKeyName, nil)
}
