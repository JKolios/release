package types

import (
	"fmt"
	"log"
	"strconv"

	"github.com/Beldur/kraken-go-api-client"
)

// TradeBalance reflects the value of all assets owned by the user
type TradeBalance struct {
	// eb = equivalent balance (combined balance of all currencies)
	// tb = trade balance (combined balance of all equity currencies)
	// m = margin amount of open positions
	// n = unrealized net profit/loss of open positions
	// c = cost basis of open positions
	// v = current floating valuation of open positions
	// e = equity = trade balance + unrealized net profit/loss
	// mf = free margin = equity - initial margin (maximum margin available to open new positions)
	EquivalentBalance, TradeBalance,
	MarginOpen, ProfitOpen, CostOpen, ValueOpen,
	Equity, FreeMargin float64
}

// TradeBalanceFromAPI constructs a TradeBalance struct by Querying the Kraken API
func TradeBalanceFromAPI(api krakenapi.KrakenApi, baseAsset string) *TradeBalance {
	apiOptions := make(map[string]string)
	apiOptions["asset"] = baseAsset
	tradeBalance, err := api.Query("TradeBalance", apiOptions)
	if err != nil {
		log.Fatalln(err)
	}
	typedBalance := tradeBalance.(map[string]interface{})

	EquivalentBalance, _ := strconv.ParseFloat(typedBalance["eb"].(string), 64)
	EquityTradeBalance, _ := strconv.ParseFloat(typedBalance["tb"].(string), 64)
	MarginOpen, _ := strconv.ParseFloat(typedBalance["m"].(string), 64)
	ProfitOpen, _ := strconv.ParseFloat(typedBalance["n"].(string), 64)
	CostOpen, _ := strconv.ParseFloat(typedBalance["c"].(string), 64)
	ValueOpen, _ := strconv.ParseFloat(typedBalance["v"].(string), 64)
	Equity, _ := strconv.ParseFloat(typedBalance["e"].(string), 64)
	FreeMargin, _ := strconv.ParseFloat(typedBalance["mf"].(string), 64)

	return &TradeBalance{EquivalentBalance, EquityTradeBalance, MarginOpen, ProfitOpen, CostOpen, ValueOpen, Equity, FreeMargin}
}

func (current TradeBalance) String() string {
	return fmt.Sprintf(`Equivalent balance:%v
    Trade balance:%v
    Margin amount of open positions:%v
    Unrealized Profit of open positions:%v
    Cost Basis of open positions:%v
    Floating valuation of open positions:%v
    Equity:%v
    Free Margin:%v`,
		current.EquivalentBalance,
		current.TradeBalance,
		current.MarginOpen,
		current.ProfitOpen,
		current.CostOpen,
		current.ValueOpen,
		current.Equity,
		current.FreeMargin,
	)
}
