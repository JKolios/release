package types

// Balance reflects the user's current  balance
import (
	"fmt"
	"log"
	"strconv"

	"github.com/Beldur/kraken-go-api-client"
)

// Balance reflects the amount the current user owns per asset
type Balance struct {
	ZEUR, XXBT, XETH float64
}

// BalanceFromAPI constructs a Balance struct by Querying the Kraken API
func BalanceFromAPI(api krakenapi.KrakenApi) *Balance {
	balance, err := api.Query("Balance", nil)
	if err != nil {
		log.Fatalln(err)
	}
	typedBalance := balance.(map[string]interface{})

	eur, err := strconv.ParseFloat(typedBalance["ZEUR"].(string), 64)
	if err != nil {
		log.Fatalln("EUR Balance conversion failed")
	}
	xbt, err := strconv.ParseFloat(typedBalance["XXBT"].(string), 64)
	if err != nil {
		log.Fatalln("XBT Balance conversion failed")
	}
	eth, err := strconv.ParseFloat(typedBalance["XETH"].(string), 64)
	if err != nil {
		log.Fatalln("ETH Balance conversion failed")
	}

	return &Balance{ZEUR: eur, XXBT: xbt, XETH: eth}
}

func (current Balance) String() string {
	return fmt.Sprintf("EUR:%v  XBT:%v  ETH:%v", current.ZEUR, current.XXBT, current.XETH)
}
