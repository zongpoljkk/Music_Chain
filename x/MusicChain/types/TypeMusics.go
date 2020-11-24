package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Musics struct {
	Creator   sdk.AccAddress   `json:"creator" yaml:"creator"`
	ID        string           `json:"id" yaml:"id"`
	MediaLink string           `json:"mediaLink" yaml:"mediaLink"`
	Price     int32            `json:"price" yaml:"price"`
	Name      string           `json:"name" yaml:"name"`
	Artists   []sdk.AccAddress `json:"artists" yaml:"artists"`
}
