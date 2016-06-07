package main

import (
	"log"

	"github.com/Beldur/kraken-go-api-client"
	"github.com/JKolios/release/apiUtils"
	"github.com/JKolios/release/types"
	"github.com/JKolios/release/utils"
)

func main() {
	log.Println("Loading configuration...")
	appconfig, err := utils.ParseJSONFile("./conf.json")
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Contacting the Kraken API...")
	api := krakenapi.New(appconfig.APIKey, appconfig.APISecret)

	timeSkew := apiUtils.TimeSkew(*api)
	log.Printf("Time skew (local - remote):%v\n", timeSkew)

	userBalance := types.BalanceFromAPI(*api)
	log.Println(userBalance.String())

	tradeBalance := types.TradeBalanceFromAPI(*api, appconfig.BaseAsset)
	log.Println(tradeBalance.String())

	userLedger := types.LedgerFromAPI(*api, "all")
	log.Println(userLedger.String())

	// assets, err := api.Assets()
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// log.Printf("Assets: %v", assets.ZEUR.Altname)

	// pairs, err := api.AssetPairs()
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// log.Printf("Pairs: %+v", pairs)

	// tickerData, err := api.Ticker("XETHZEUR")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// log.Printf("ETH/EUR: %v", tickerData.XETHZEUR)

}
