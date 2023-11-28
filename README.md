# Broadcast Client Module

## Description
This module is a simple client that used to send broadcast transaction to the network and can monitor the status of the transaction. This module can integrated into other application that developed by Go.

## Installation
```bash
$ go get github.com/mynameismaxz/http_broadcast_client
```

## Usage
```go
package main

import (
    "fmt"
    "github.com/mynameismaxz/http_broadcast_client"
    "time"
)

func main() {
    // Create a new client
    c := client.NewClient("<URL of server>")

    // create a simple transaction
    tx := []byte(`{"price":37037,"symbol":"BTC","timestamp":1701168552}`)

    // broadcast transaction
    txHash, err := c.BroadcastTx(tx)
    if err != nil {
        panic(err)
    }
    fmt.Println("txHash:", txHash)

    // monitor transaction status
    for {
        status, err := c.GetTxStatus(txHash)
        if err != nil {
            panic(err)
        }
        fmt.Println("status:", status)
        if status == "CONFIRMED" {
            break
        }
        time.Sleep(5 * time.Second)
    }
}
```
And so other use cases, you can see more in [example/broadcast]("example/broadcast/main.go") to see more detail.

## License
N/A