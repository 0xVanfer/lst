package lst

import (
	"github.com/0xVanfer/abigen/lido/lidowstETH"
	"github.com/0xVanfer/chainId"
	"github.com/0xVanfer/ethaddr"
	"github.com/0xVanfer/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// 1 stETH === 1 ETH
func getSTETHStakeRates(connector *ethclient.Client, block ...int64) (stake float64, unstake float64, err error) {
	_ = connector
	_ = block
	// opt := getCallOpts(block...)
	return 1, 1, nil
}

// wstETH / stETH = wstETH / ETH
func getWSTETHStakeRates(connector *ethclient.Client, block ...int64) (stake float64, unstake float64, err error) {
	opt := getCallOpts(block...)
	network := chainId.EthereumChainName

	wstETH, err := lidowstETH.NewLidowstETH(types.ToAddress(ethaddr.WSTETHList[network]), connector)
	if err != nil {
		return
	}
	unwrapx, err := wstETH.StEthPerToken(opt)
	if err != nil {
		return
	}
	unwrap := types.ToFloat64(unwrapx) * 1e-18
	return 1 / unwrap, unwrap, nil
}
