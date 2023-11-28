package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	client "github.com/mynameismaxz/http_broadcast_client"
)

const (
	MOCK_SERVER = "https://mock-node-wgqbnxruha-as.a.run.app"
)

func main() {
	// initialize client
	client := client.NewClient(MOCK_SERVER)

	// example, i'll fetch BTC price from public server and broadcast it.
	// you can use your own data.
	// ex: https://min-api.cryptocompare.com/data/generateAvg?fsym=BTC&tsym=USD&e=coinbase
	httpClient := &http.Client{}
	resp, err := httpClient.Get("https://min-api.cryptocompare.com/data/price?fsym=BTC&tsyms=USD")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		panic(err)
	}
	// convert result to Transaction struct
	tx := map[string]interface{}{
		"symbol":    "BTC",
		"price":     uint64(result["USD"].(float64)),
		"timestamp": time.Now().Unix(),
	}
	// convert to []byte JSON
	txBytes, err := json.Marshal(tx)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(txBytes))

	// broadcast TX
	txHash, err := client.BroadcastTx(txBytes)
	if err != nil {
		panic(err)
	}
	fmt.Println(txHash)

	// Get TX status
	txStatus, err := client.GetTxStatus(txHash)
	if err != nil {
		panic(err)
	}
	// txStatus is a string that can be "CONFIRMED", "FAILED", "PENDING" and "DNE"
	fmt.Println(string(txStatus))
}
