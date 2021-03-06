// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package genesis

import (
	"github.com/Toinounet21/avalanchego-trafficked-v1.7.4/ids"
	"github.com/Toinounet21/avalanchego-trafficked-v1.7.4/utils/constants"
	"github.com/Toinounet21/avalanchego-trafficked-v1.7.4/vms/avm"
	"github.com/Toinounet21/avalanchego-trafficked-v1.7.4/vms/evm"
	"github.com/Toinounet21/avalanchego-trafficked-v1.7.4/vms/nftfx"
	"github.com/Toinounet21/avalanchego-trafficked-v1.7.4/vms/platformvm"
	"github.com/Toinounet21/avalanchego-trafficked-v1.7.4/vms/propertyfx"
	"github.com/Toinounet21/avalanchego-trafficked-v1.7.4/vms/secp256k1fx"
)

// Aliases returns the default aliases based on the network ID
func Aliases(genesisBytes []byte) (map[string][]string, map[ids.ID][]string, error) {
	apiAliases := map[string][]string{
		constants.ChainAliasPrefix + constants.PlatformChainID.String(): {
			"P",
			"platform",
			constants.ChainAliasPrefix + "P",
			constants.ChainAliasPrefix + "platform",
		},
	}
	chainAliases := map[ids.ID][]string{
		constants.PlatformChainID: {"P", "platform"},
	}
	genesis := &platformvm.Genesis{} // TODO let's not re-create genesis to do aliasing
	if _, err := platformvm.GenesisCodec.Unmarshal(genesisBytes, genesis); err != nil {
		return nil, nil, err
	}
	if err := genesis.Initialize(); err != nil {
		return nil, nil, err
	}

	for _, chain := range genesis.Chains {
		uChain := chain.UnsignedTx.(*platformvm.UnsignedCreateChainTx)
		switch uChain.VMID {
		case avm.ID:
			apiAliases[constants.ChainAliasPrefix+chain.ID().String()] = []string{"X", "avm", constants.ChainAliasPrefix + "X", constants.ChainAliasPrefix + "avm"}
			chainAliases[chain.ID()] = GetXChainAliases()
		case evm.ID:
			apiAliases[constants.ChainAliasPrefix+chain.ID().String()] = []string{"C", "evm", constants.ChainAliasPrefix + "C", constants.ChainAliasPrefix + "evm"}
			chainAliases[chain.ID()] = GetCChainAliases()
		}
	}
	return apiAliases, chainAliases, nil
}

func GetCChainAliases() []string {
	return []string{"C", "evm"}
}

func GetXChainAliases() []string {
	return []string{"X", "avm"}
}

func GetVMAliases() map[ids.ID][]string {
	return map[ids.ID][]string{
		platformvm.ID:  {"platform"},
		avm.ID:         {"avm"},
		evm.ID:         {"evm"},
		secp256k1fx.ID: {"secp256k1fx"},
		nftfx.ID:       {"nftfx"},
		propertyfx.ID:  {"propertyfx"},
	}
}
