package codec

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/crypto/keys"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"

	emintcrypto "github.com/cosmos/ethermint/crypto"
	ethermint "github.com/cosmos/ethermint/types"
)

// MakeCodec registers the necessary types and interfaces for an sdk.App. This
// codec is provided to all the modules the application depends on.
//
// NOTE: This codec will be deprecated in favor of AppCodec once all modules are
// migrated to protobuf.
func MakeCodec(bm module.BasicManager) *codec.Codec {
	cdc := codec.NewLegacyAminoLegacyAmino()

	bm.RegisterLegacyAminoCodec(cdc)
	vesting.RegisterLegacyAminoCodec(cdc)
	sdk.RegisterLegacyAminoCodec(cdc)
	emintcrypto.RegisterLegacyAminoCodec(cdc)
	codec.RegisterCrypto(cdc)
	ethermint.RegisterLegacyAminoCodec(cdc)
	keys.RegisterLegacyAminoCodec(cdc) // temporary. Used to register keyring.Info

	return cdc
}
