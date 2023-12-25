package lst

import (
	"github.com/0xVanfer/abigen/benqi/benqisavax"
	"github.com/0xVanfer/chainId"
	"github.com/0xVanfer/ethaddr"
	"github.com/0xVanfer/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func getSAVAXStakeRates(connector *ethclient.Client, block ...int64) (stake float64, unstake float64, err error) {
	opt := getCallOpts(block...)

	network := chainId.AvalancheChainName
	savax, err := benqisavax.NewBenqisavax(types.ToAddress(ethaddr.SAVAXList[network]), connector)
	if err != nil {
		return
	}
	totalPooledAvax, err := savax.TotalPooledAvax(opt)
	if err != nil {
		return
	}
	totalSupply, err := savax.TotalSupply(opt)
	if err != nil {
		return
	}
	unstake = types.ToFloat64(totalPooledAvax) / types.ToFloat64(totalSupply)
	return 1 / unstake, unstake, nil
}
