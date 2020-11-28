package rest

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/types/rest"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/gorilla/mux"
	"github.com/zongpoljkk/MusicChain/x/MusicChain/types"
	"github.com/zongpoljkk/MusicChain/x/MusicChain/utils"
)

func listPurchasedHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/list-purchased", storeName), nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}
		rest.PostProcessResponse(w, cliCtx, res)
	}
}

func getPurchasedHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		key := vars["key"]

		res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/get-purchased/%s", storeName, key), nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}
		rest.PostProcessResponse(w, cliCtx, res)
	}
}

func playPurchasedHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var customTx authtypes.StdTx
		var musicInfo types.Musics

		body, _ := ioutil.ReadAll(r.Body)
		err := cliCtx.Codec.UnmarshalJSON(body, &customTx)

		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		PurMusicTempoLinkReq := customTx.Msgs[0].(types.MsgRequestPurchasedMusicTemporarilyLink)
		err = PurMusicTempoLinkReq.ValidateBasic()

		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		addr := PurMusicTempoLinkReq.GetSigners()[0]
		accGetter := authtypes.NewAccountRetriever(cliCtx)
		account, err := accGetter.GetAccount(addr)

		if err := accGetter.EnsureExists(addr); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "Signer address not found.")
			return
		}

		pubKey := account.GetPubKey()
		if pubKey == nil {
			rest.WriteErrorResponse(w, http.StatusUnauthorized, "pubkey on account is not set")
			return
		}

		signBytes := authtypes.StdSignBytes(
			cliCtx.ChainID,
			account.GetAccountNumber(),
			0,
			customTx.Fee,
			customTx.Msgs,
			customTx.Memo,
		)

		if !pubKey.VerifyBytes(signBytes, customTx.Signatures[0].Signature) {
			rest.WriteErrorResponse(w, http.StatusUnauthorized, "signature verification failed.")
			return
		}

		_, _, err = cliCtx.QueryWithData(
			fmt.Sprintf("custom/%s/get-purchased/%s-%s",
				storeName,
				PurMusicTempoLinkReq.Requestor,
				PurMusicTempoLinkReq.MusicID,
			),
			nil)

		if err != nil {
			rest.WriteErrorResponse(w, http.StatusPaymentRequired, "error while querying purchased info, did u purchased this music?")
			return
		}

		musicInfoBytes, _, err := cliCtx.QueryWithData(
			fmt.Sprintf("custom/%s/get-musics/%s",
				storeName,
				PurMusicTempoLinkReq.MusicID,
			),
			nil)
		cliCtx.Codec.UnmarshalJSON(musicInfoBytes, &musicInfo)

		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, "error while querying music info, "+err.Error())
			return
		}

		tempLink, err := utils.GetObjectSignedURL(musicInfo.MediaLink, PurMusicTempoLinkReq.Timestamp, 10*time.Minute)

		w.Write([]byte(tempLink))
		return
	}
}
