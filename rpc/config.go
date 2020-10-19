package rpc

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client"

	"github.com/cosmos/ethermint/server"

	"github.com/ethereum/go-ethereum/rpc"
)

// RegisterRoutes creates a new ethereum JSON-RPC server and recreates a CLI command to start Cosmos REST server with web3 RPC API and
// Cosmos rest-server endpoints
func RegisterRoutes(clientCtx client.Context, r *mux.Router, apiConfig server.APIConfig) {
	server := rpc.NewServer()
	r.HandleFunc("/", server.ServeHTTP).Methods("POST", "OPTIONS")

	apis := GetRPCAPIs(clientCtx)

	// Register all the APIs exposed by the namespace services
	// TODO: handle allowlist and private APIs
	for _, api := range apis {
		if err := server.RegisterName(api.Namespace, api.Service); err != nil {
			panic(err)
		}
	}
}

// StartEthereumWebsocket starts the Filter api websocket
func StartEthereumWebsocket(clientCtx client.Context, apiConfig server.APIConfig) {
	ws := newWebsocketsServer(clientCtx, apiConfig.WebsocketAddress)
	ws.start()
}
