// Copyright (c) 2025 The bel2 developers

package mempool

import (
	"fmt"
	"testing"
)

func TestGetFee(t *testing.T) {
	t.Skip()
	api := NewAPI(DefaultConfig)
	response, err := api.GetRawTransaction("ad0d072993e4f482fe97097d3cdec2ba24068e07028f892ee39678f9e20c9f36")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(response)
}
