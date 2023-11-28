package client_test

import (
	"os"
	"testing"

	client "github.com/mynameismaxz/http_broadcast_client"
)

var c *client.Client

func TestMain(m *testing.M) {
	setup()
	runtimeCode := m.Run()

	os.Exit(runtimeCode)
}

func setup() {
	c = client.NewClient("https://mock-node-wgqbnxruha-as.a.run.app")
}

func TestBroadcastTx(t *testing.T) {
	tx := []byte(`{"symbol":"BTC","price":10000,"timestamp":1623720000}`)
	txHash, err := c.BroadcastTx(tx)
	if err != nil {
		t.Error(err)
	}
	t.Log(txHash)
}

func TestGetTxStatus(t *testing.T) {
	txHash := "aecc734572b5a5886f2e9807d9bcdbd7e3c9c0c9398126c9cdfd6acb1bb2cc6a"
	txStatus, err := c.GetTxStatus(txHash)
	if err != nil {
		t.Error(err)
	}
	t.Log(txStatus)
}
