package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// RegisterCodec registers concrete types on codec
func RegisterCodec(cdc *codec.Codec) {
	// this line is used by starport scaffolding # 1
	cdc.RegisterConcrete(MsgCreatePurchased{}, "MusicChain/CreatePurchased", nil)
	cdc.RegisterConcrete(MsgCreateMusics{}, "MusicChain/CreateMusics", nil)
	cdc.RegisterConcrete(MsgSetMusics{}, "MusicChain/SetMusics", nil)
	cdc.RegisterConcrete(MsgDeleteMusics{}, "MusicChain/DeleteMusics", nil)
	cdc.RegisterConcrete(MsgCreateArtist{}, "MusicChain/CreateArtist", nil)
	cdc.RegisterConcrete(MsgRequestPurchasedMusicTemporarilyLink{}, "MusicChain/ReqPMusicTempL", nil)
}

// ModuleCdc defines the module codec
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}
