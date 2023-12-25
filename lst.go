package lst

import (
	"strings"

	"github.com/0xVanfer/chainId"
	"github.com/0xVanfer/ethaddr"
	"github.com/ethereum/go-ethereum/ethclient"
)

var lsts = []*LST{
	// stETH
	{
		Network:    chainId.EthereumChainName,
		Token:      &erc20{Symbol: "stETH", Address: ethaddr.STETHList[chainId.EthereumChainName], Decimals: 18},
		Underlying: &erc20{Symbol: "WETH", Address: ethaddr.WETHList[chainId.EthereumChainName], Decimals: 18},
	},
	// wstETH
	{
		Network:    chainId.EthereumChainName,
		Token:      &erc20{Symbol: "wstETH", Address: ethaddr.WSTETHList[chainId.EthereumChainName], Decimals: 18},
		Underlying: &erc20{Symbol: "WETH", Address: ethaddr.WETHList[chainId.EthereumChainName], Decimals: 18},
	},
	// sAVAX
	{
		Network:    chainId.AvalancheChainName,
		Token:      &erc20{Symbol: "sAVAX", Address: ethaddr.SAVAXList[chainId.AvalancheChainName], Decimals: 18},
		Underlying: &erc20{Symbol: "WAVAX", Address: ethaddr.WAVAXList[chainId.AvalancheChainName], Decimals: 18},
	},
	// MATICX
	{
		Network:    chainId.PolygonChainName,
		Token:      &erc20{Symbol: "MaticX", Address: ethaddr.MaticXList[chainId.PolygonChainName], Decimals: 18},
		Underlying: &erc20{Symbol: "WMATIC", Address: ethaddr.WMATICList[chainId.PolygonChainName], Decimals: 18},
	},
}

func GetLsts() []*LST { return lsts }

func GetLst(symbolOrAddr string) *LST {
	return decideLst(symbolOrAddr)
}

func GetStakeRate(symbolOrAddr string, connector *ethclient.Client, block ...int64) (stake float64, unstake float64, err error) {
	lst := decideLst(symbolOrAddr)
	return lst.StakeRate(connector, block...)
}

func decideLst(symbolOrAddr string) *LST {
	for _, lst := range lsts {
		if strings.EqualFold(lst.Token.Address, symbolOrAddr) ||
			strings.EqualFold(lst.Token.Symbol, symbolOrAddr) {
			return lst
		}
	}
	return nil
}
