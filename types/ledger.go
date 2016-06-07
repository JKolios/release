package types

// Balance reflects the user's current  balance
import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/Beldur/kraken-go-api-client"
)

//TODO:SORT LEDGER ENTRIES DATE ASCENDING/DESCENDING

// A Ledger contains the details of all transactions of a particular asset type made by a user
type Ledger struct {
	Transactions map[string]Transaction
}

// LedgerFromAPI constructs a Ledger struct by querying the Kraken API
func LedgerFromAPI(api krakenapi.KrakenApi, assetType string) *Ledger {
	input := make(map[string]string)
	input["aclass"] = assetType
	ledgerResponse, err := api.Query("Ledgers", input)

	if err != nil {
		log.Fatalln(err)
	}
	ledger := &Ledger{Transactions: make(map[string]Transaction)}
	typedResponse := ledgerResponse.(map[string]interface{})["ledger"].(map[string]interface{})

	for transactionID, transaction := range typedResponse {
		transaction := transaction.(map[string]interface{})

		parsedAmount, err := strconv.ParseFloat(transaction["amount"].(string), 64)
		if err != nil {
			log.Fatalln(err)
		}

		parsedFee, err := strconv.ParseFloat(transaction["fee"].(string), 64)
		if err != nil {
			log.Fatalln(err)
		}

		parsedBalance, err := strconv.ParseFloat(transaction["balance"].(string), 64)
		if err != nil {
			log.Fatalln(err)
		}
		transactionRecord := Transaction{
			ReferenceID:     transaction["refid"].(string),
			TransactionType: transaction["type"].(string),
			AssetClass:      transaction["aclass"].(string),
			Asset:           transaction["asset"].(string),
			Time:            transaction["time"].(float64),
			Amount:          parsedAmount,
			Fee:             parsedFee,
			Balance:         parsedBalance,
		}

		ledger.Transactions[transactionID] = transactionRecord
	}

	return ledger
}

//String() converts a ledger to a pretty-printed string
func (ledger Ledger) String() string {
	var ledgerString string
	for transactionID, transaction := range ledger.Transactions {
		ledgerString += fmt.Sprintf("Transaction %v:\n %v\n", transactionID, transaction.String())
	}
	return ledgerString
}

// A Transaction contains all details of a particular transaction
type Transaction struct {
	ReferenceID, TransactionType, AssetClass, Asset string
	Time, Amount, Fee, Balance                      float64
}

//String() converts a Transaction to a pretty-printed string
func (transaction Transaction) String() string {
	return fmt.Sprintf("\tReferenceID: %v\n\tTransactionType: %v\n\tAssetClass: %v\n\tAsset: %v\n\tTime: %v\n\tAmount: %v\n\tFee: %v\n\tBalance: %v\n",
		transaction.ReferenceID,
		transaction.TransactionType,
		transaction.AssetClass,
		transaction.Asset,
		time.Unix(int64(transaction.Time), 0),
		transaction.Amount,
		transaction.Fee,
		transaction.Balance)
}
