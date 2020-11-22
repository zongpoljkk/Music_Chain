package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Artist struct {
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	Name    string         `json:"name" yaml:"name"`
}
