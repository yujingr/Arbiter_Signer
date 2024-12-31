// Copyright (c) 2025 The bel2 developers

package mempool

type Config struct {
	Network string

	ApiBaseUrl string
	Proxy      string
}

const Mainnet_ApiBaseUrl = "https://mempool.space/api/tx/"
const Testnet_ApiBaseUrl = "https://mempool.space/testnet/api/tx/"

var DefaultConfig = Config{
	ApiBaseUrl: Testnet_ApiBaseUrl,
}
