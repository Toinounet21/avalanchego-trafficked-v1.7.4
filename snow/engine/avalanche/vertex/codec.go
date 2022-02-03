// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package vertex

import (
	"github.com/Toinounet21/avalanchego-trafficked-v1.7.4/codec"
	"github.com/Toinounet21/avalanchego-trafficked-v1.7.4/codec/linearcodec"
	"github.com/Toinounet21/avalanchego-trafficked-v1.7.4/codec/reflectcodec"
	"github.com/Toinounet21/avalanchego-trafficked-v1.7.4/utils/units"
)

const (
	// maxSize is the maximum allowed vertex size. It is necessary to deter DoS
	maxSize = units.MiB

	codecVersion            uint16 = 0
	codecVersionWithStopVtx uint16 = 1
)

var c codec.Manager

func init() {
	lc := linearcodec.New(reflectcodec.DefaultTagName+"V0", maxSize)
	lc2 := linearcodec.New(reflectcodec.DefaultTagName+"V1", maxSize)

	c = codec.NewManager(maxSize)
	// for backward compatibility, still register the initial codec version
	if err := c.RegisterCodec(codecVersion, lc); err != nil {
		panic(err)
	}
	if err := c.RegisterCodec(codecVersionWithStopVtx, lc2); err != nil {
		panic(err)
	}
}