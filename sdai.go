package lst

import (
	"github.com/0xVanfer/abigen/erc4626"
	"github.com/0xVanfer/chainId"
	"github.com/0xVanfer/ethaddr"
	"github.com/0xVanfer/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func getSDAIStakeRates(connector *ethclient.Client, block ...int64) (stake float64, unstake float64, err error) {
	opt := getCallOpts(block...)

	network := chainId.EthereumChainName
	sdai, err := erc4626.NewErc4626(types.ToAddress(ethaddr.SDAIList[network]), connector)
	if err != nil {
		return
	}
	unstakx, err := sdai.ConvertToAssets(opt, types.ToBigInt(1e18))
	if err != nil {
		return
	}

	unstake = types.ToFloat64(unstakx) * 1e-18
	return 1 / unstake, unstake, nil
}
