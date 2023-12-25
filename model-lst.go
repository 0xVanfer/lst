package lst

import (
	"github.com/0xVanfer/ethaddr"
	"github.com/ethereum/go-ethereum/ethclient"
)

type LST struct {
	Network    string
	Token      *erc20
	Underlying *erc20
}

type erc20 struct {
	Address  string
	Symbol   string
	Decimals float64
}

func (l *LST) StakeRate(connector *ethclient.Client, block ...int64) (stake float64, unstake float64, err error) {
	if l == nil {
		return
	}

	switch l.Token.Address {
	// Aavalanche
	case ethaddr.SAVAXList[l.Network]:
		// sAVAX
		return getSAVAXStakeRates(connector, block...)
	// Polygon
	case ethaddr.MaticXList[l.Network]:
		// MaticX
		return getMATICXStakeRates(connector, block...)
	// Ethereum
	case ethaddr.STETHList[l.Network]:
		// stETH
		return getSTETHStakeRates(connector, block...)
	case ethaddr.WSTETHList[l.Network]:
		// wstETH
		return getWSTETHStakeRates(connector, block...)
	case ethaddr.SDAIList[l.Network]:
		// sDAI
		return getSDAIStakeRates(connector, block...)
	default:
		return
	}
}
