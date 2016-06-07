package apiUtils

import (
	"log"
	"time"

	"github.com/Beldur/kraken-go-api-client"
)

// TimeSkew calculates the difference (localTime - serverTime)
func TimeSkew(api krakenapi.KrakenApi) time.Duration {
	log.Println("Calculating time skew...")

	localTime := time.Now()

	serverTimes, err := api.Query("Time", nil)
	if err != nil {
		log.Fatalln(err)
	}
	serverTimestamp, ok := serverTimes.(map[string]interface{})["unixtime"].(float64)
	if !ok {
		log.Fatalln("Server Timestamp conversion failed")
	}

	serverTime := time.Unix(int64(serverTimestamp), 0)
	timeSkew := localTime.Sub(serverTime)

	return timeSkew
}
