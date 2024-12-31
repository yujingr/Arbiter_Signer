// Copyright (c) 2025 The bel2 developers

package mempool

type GetRawTransactionResponse struct {
	Txid string `json:"txid"`
	Vin  []struct {
		Txid    string `json:"txid"`
		Vout    int    `json:"vout"`
		Prevout struct {
			Scriptpubkey        string `json:"scriptpubkey"`
			ScriptpubkeyAsm     string `json:"scriptpubkey_asm"`
			ScriptpubkeyType    string `json:"scriptpubkey_type"`
			ScriptpubkeyAddress string `json:"scriptpubkey_address"`
			Value               int64  `json:"value"`
		} `json:"prevout"`
		Scriptsig             string   `json:"scriptsig"`
		ScriptsigAsm          string   `json:"scriptsig_asm"`
		Witness               []string `json:"witness"`
		IsCoinbase            bool     `json:"is_coinbase"`
		Sequence              uint64   `json:"sequence"`
		InnerWitnessscriptAsm string   `json:"inner_witnessscript_asm"`
	} `json:"vin"`
	Vout []struct {
		Scriptpubkey        string `json:"scriptpubkey"`
		ScriptpubkeyAsm     string `json:"scriptpubkey_asm"`
		ScriptpubkeyType    string `json:"scriptpubkey_type"`
		ScriptpubkeyAddress string `json:"scriptpubkey_address"`
		Value               int64  `json:"value"`
	} `json:"vout"`
	Size   int64 `json:"size"`
	Weight int64 `json:"weight"`
	Sigops int64 `json:"sigops"`
	Fee    int64 `json:"fee"`
	Status struct {
		Confirmed   bool   `json:"confirmed"`
		BlockHeight int64  `json:"block_height"`
		BlockHash   string `json:"block_hash"`
		BlockTime   int64  `json:"block_time"`
	} `json:"status"`
}
