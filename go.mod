module github.com/zongpoljkk/MusicChain

go 1.15

require (
	cloud.google.com/go v0.72.0 // indirect
	cloud.google.com/go/storage v1.12.0
	github.com/cosmos/cosmos-sdk v0.39.1
	github.com/google/uuid v1.1.2
	github.com/gorilla/mux v1.7.4
	github.com/onsi/ginkgo v1.8.0 // indirect
	github.com/onsi/gomega v1.5.0 // indirect
	github.com/spf13/afero v1.2.2 // indirect
	github.com/spf13/cobra v1.0.0
	github.com/spf13/viper v1.7.0
	github.com/tendermint/go-amino v0.15.1
	github.com/tendermint/tendermint v0.33.7
	github.com/tendermint/tm-db v0.5.1
	golang.org/x/net v0.0.0-20201110031124-69a78807bb2b // indirect
	golang.org/x/oauth2 v0.0.0-20201109201403-9fd604954f58
	golang.org/x/sys v0.0.0-20201119102817-f84b799fce68 // indirect
	golang.org/x/tools v0.0.0-20201124005743-911501bfb504 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto v0.0.0-20201119123407-9b1e624d6bc4 // indirect
)

replace github.com/cosmos/cosmos-sdk v0.39.1 => github.com/intaniger/cosmos-sdk v0.39.1-groot
