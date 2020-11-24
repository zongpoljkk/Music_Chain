package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
)

// RegisterRoutes registers MusicChain-related REST handlers to a router
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
	// this line is used by starport scaffolding # 1
	r.HandleFunc("/MusicChain/musics", createMusicsHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/MusicChain/musics", listMusicsHandler(cliCtx, "MusicChain")).Methods("GET")
	r.HandleFunc("/MusicChain/musics/{key}", getMusicsHandler(cliCtx, "MusicChain")).Methods("GET")
	r.HandleFunc("/MusicChain/musics", setMusicsHandler(cliCtx)).Methods("PUT")
	r.HandleFunc("/MusicChain/musics", deleteMusicsHandler(cliCtx)).Methods("DELETE")

	r.HandleFunc("/MusicChain/artist", createArtistHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/MusicChain/artist", listArtistHandler(cliCtx, "MusicChain")).Methods("GET")
	r.HandleFunc("/MusicChain/artist/{key}", getArtistHandler(cliCtx, "MusicChain")).Methods("GET")

}
