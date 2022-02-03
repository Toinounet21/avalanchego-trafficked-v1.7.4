// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package state

import (
	"math"

	"github.com/Toinounet21/avalanchego-trafficked-v1.7.4/codec"
	"github.com/Toinounet21/avalanchego-trafficked-v1.7.4/codec/linearcodec"
	"github.com/Toinounet21/avalanchego-trafficked-v1.7.4/codec/reflectcodec"
)

const version = 0

var c codec.Manager

func init() {
	lc := linearcodec.New(reflectcodec.DefaultTagName, math.MaxUint32)
	c = codec.NewManager(math.MaxInt32)

	err := c.RegisterCodec(version, lc)
	if err != nil {
		panic(err)
	}
}
