package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Purchased struct {
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	ID      string         `json:"id" yaml:"id"`
    MusicID string `json:"MusicID" yaml:"MusicID"`
}