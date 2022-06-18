package client

import (
	"math"
	"math/big"
)

//	WeiToEth wei 转换为 eth
func WeiToEth(amount *big.Int) *big.Float{
	f := new(big.Float)
	f.SetString(amount.String())
	ethValue := new(big.Float).Quo(f, big.NewFloat(math.Pow10(18)))
	return ethValue
}





