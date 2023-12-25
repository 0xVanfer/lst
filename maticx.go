package lst

import (
	"math/big"

	"github.com/0xVanfer/abigen/stader/staderChildPool"
	"github.com/0xVanfer/chainId"
	"github.com/0xVanfer/ethaddr"
	"github.com/0xVanfer/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func getMATICXStakeRates(connector *ethclient.Client, block ...int64) (stake float64, unstake float64, err error) {
	opt := getCallOpts(block...)
	network := chainId.PolygonChainName

	childPool, err := staderChildPool.NewStaderChildPool(types.ToAddress(ethaddr.StaderChildPoolList[network]), connector)
	if err != nil {
		return
	}
	unstakex, _, _, err := childPool.ConvertMaticXToMatic(opt, big.NewInt(1e18))
	if err != nil {
		return
	}
	unstake = types.ToFloat64(unstakex) * 1e-18
	return 1 / unstake, stake, nil
}
