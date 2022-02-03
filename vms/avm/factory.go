// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package avm

import (
	"github.com/Toinounet21/avalanchego-trafficked-v1.7.4/ids"
	"github.com/Toinounet21/avalanchego-trafficked-v1.7.4/snow"
)

// ID that this VM uses when labeled
var (
	ID = ids.ID{'a', 'v', 'm'}
)

type Factory struct {
	TxFee            uint64
	CreateAssetTxFee uint64
}

func (f *Factory) New(*snow.Context) (interface{}, error) {
	return &VM{Factory: *f}, nil
}
