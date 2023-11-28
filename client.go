package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type Client struct {
	Hostname   string
	httpClient *http.Client
}

// type Transaction that will be used to send to server.
type Transaction struct {
	Symbol    string `json:"symbol:"`
	Price     uint64 `json:"price"`
	TimeStamp uint64 `json:"timestamp"`
}

// type BroadcastResponse that will be used to receive response from server.
type BroadcastResponse struct {
	TxHash string `json:"tx_hash"`
}

// TransactionResponse is a response from server that show status of transaction by txHash.
type TransactionResponse struct {
	TxStatus string `json:"tx_status"`
}

func NewClient(hostname string) *Client {
	chttp := &http.Client{
		Timeout: 10 * time.Second,
	}
	return &Client{Hostname: hostname, httpClient: chttp}
}

// BroadcastTx is a broadcast transaction function. It will send transaction that is in JSON format to server.
func (c *Client) BroadcastTx(tx []byte) (string, error) {
	// validate tx
	if err := json.Unmarshal(tx, &Transaction{}); err != nil {
		return "", err
	}
	// convert to io.Reader
	reader := bytes.NewBuffer(tx)

	// broadcast tx to server
	resp, err := c.httpClient.Post(c.Hostname+"/broadcast", "application/json", reader)
	if err != nil {
		return "", err
	}

	// check status code of response
	if resp.StatusCode != http.StatusOK {
		return "", errors.New("broadcast tx failed, status code: " + resp.Status)
	}
	defer resp.Body.Close()

	// read response body
	var result BroadcastResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	return result.TxHash, nil
}

// GetTxStatus is a get transaction status function. That's receive txHash and return status of transaction from server are `CONFIRMED`, `FAILED`, `PENDING` or `DNE`.
func (c *Client) GetTxStatus(txHash string) ([]byte, error) {
	// get tx status from server
	resp, err := c.httpClient.Get(c.Hostname + "/check/" + txHash)
	if err != nil {
		return nil, err
	}

	// check status code of response
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("get tx status failed, status code: " + resp.Status)
	}
	defer resp.Body.Close()

	// read response body
	var result TransactionResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return []byte(result.TxStatus), nil
}
