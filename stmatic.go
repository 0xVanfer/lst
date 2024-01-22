package lst

import (
	"math"
	"math/big"

	"github.com/0xVanfer/abigen/lido/stMATIC"
	"github.com/0xVanfer/chainId"
	"github.com/0xVanfer/ethaddr"
	"github.com/0xVanfer/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func getStMaticStakeRates(connector *ethclient.Client, block ...int64) (stake float64, unstake float64, err error) {
	opt := getCallOpts(block...)

	network := chainId.EthereumChainName
	stmatic, err := stMATIC.NewStMATIC(types.ToAddress(ethaddr.STMATICList[network]), connector)
	if err != nil {
		return
	}
	unstakex, err := stmatic.ConvertStMaticToMatic(opt, big.NewInt(1e18))
	if err != nil {
		return
	}

	unstake = types.ToFloat64(unstakex.AmountInMatic) * math.Pow10(-18)
	return 1 / unstake, unstake, nil
}
