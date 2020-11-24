package cli

import (
	"bufio"
	"io/ioutil"
	"path/filepath"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	sdkutils "github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/zongpoljkk/MusicChain/x/MusicChain/types"
)

func GetCmdCreateMusics(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "register-musics [media's local file path] [price] [name]",
		Short: "Register a new musics",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsMediaPath := string(args[0])
			argsPrice, _ := strconv.ParseInt(args[1], 10, 64)
			argsName := string(args[2])

			data, err := ioutil.ReadFile(argsMediaPath)
			if err != nil {
				return err
			}

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(sdkutils.GetTxEncoder(cdc))
			msg, err := types.UploadFileAndCreateMsgCreateMusics(cliCtx.GetFromAddress(), int32(argsPrice), string(argsName), data, filepath.Ext(argsMediaPath))
			if err != nil {
				return err
			}

			err = msg.ValidateBasic()
			if err != nil {
				return err
			}
			return sdkutils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

func GetCmdSetMusics(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "set-musics [id] [price] [name]",
		Short: "Set a new musics",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			id := args[0]
			argsPrice, _ := strconv.ParseInt(args[1], 10, 64)
			argsName := string(args[2])

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(sdkutils.GetTxEncoder(cdc))
			msg := types.NewMsgSetMusics(cliCtx.GetFromAddress(), id, int32(argsPrice), string(argsName))
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return sdkutils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

func GetCmdDeleteMusics(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "delete-musics [id]",
		Short: "Delete a new musics by ID",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(sdkutils.GetTxEncoder(cdc))

			msg := types.NewMsgDeleteMusics(args[0], cliCtx.GetFromAddress())
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return sdkutils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
