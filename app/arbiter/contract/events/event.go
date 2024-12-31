// Copyright (c) 2025 The bel2 developers

package events

import (
	"github.com/ethereum/go-ethereum/common"
)

type State uint8

const (
	Request  State = iota
	Response State = iota
	Confirm  State = iota
)

type ContractLogEvent struct {
	EventData []byte
	TxHash    common.Hash
	Topics    []common.Hash
	Block     uint64
	TxIndex   uint
}
