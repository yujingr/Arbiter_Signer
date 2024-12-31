// Copyright (c) 2025 The bel2 developers

package events

import (
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	ArbitrationRequested = crypto.Keccak256Hash([]byte("ArbitrationRequested(address,bytes32,bytes,bytes,address)"))

	ArbitrationResultSubmitted = crypto.Keccak256Hash([]byte("ArbitrationResultSubmitted(bytes,bytes32)"))
)
