package utils

import (
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/params"

	"github.com/consolelabs/mochi-pay-api/internal/model"
)

func ConvertBigIntString(amount string, token *model.Token) *big.Float {
	decimals := token.Decimal
	if token.Decimal == 0 {
		decimals = 18
	}
	wei := new(big.Int)
	wei.SetString(amount, 10)
	return WeiToEther(wei, int(decimals))
}

func WeiToEther(wei *big.Int, decimals ...int) *big.Float {
	f := new(big.Float)
	f.SetPrec(236) //  IEEE 754 octuple-precision binary floating-point format: binary256
	f.SetMode(big.ToNearestEven)
	fWei := new(big.Float)
	fWei.SetPrec(236) //  IEEE 754 octuple-precision binary floating-point format: binary256
	fWei.SetMode(big.ToNearestEven)

	var e *big.Float
	if len(decimals) == 0 {
		e = big.NewFloat(params.Ether)
	} else {
		e = big.NewFloat(math.Pow(10, float64(decimals[0])))
	}
	return f.Quo(fWei.SetInt(wei), e)
}
