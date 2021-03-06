package cli

import (
	"fmt"
	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"

	// "github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/zongpoljkk/MusicChain/x/MusicChain/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string, cdc *codec.Codec) *cobra.Command {
	// Group MusicChain queries under a subcommand
	MusicChainQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	MusicChainQueryCmd.AddCommand(
		flags.GetCommands(
			// this line is used by starport scaffolding # 1
			GetCmdListPurchased(queryRoute, cdc),
			GetCmdGetPurchased(queryRoute, cdc),
			GetCmdListMusics(queryRoute, cdc),
			GetCmdGetMusics(queryRoute, cdc),
			GetCmdListArtist(queryRoute, cdc),
			GetCmdGetArtist(queryRoute, cdc),
		)...,
	)

	return MusicChainQueryCmd
}
