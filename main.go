package main

import (
  "fmt"
  "log"
  "context"
  "math/big"
  //"encoding/hex"

  "github.com/ethereum/go-ethereum/ethclient"
  "github.com/ethereum/go-ethereum/core/types"
  "github.com/ethereum/go-ethereum/common"
  "github.com/ethereum/go-ethereum/mobile/common"
)

func main() {
  //var ctx context.Context

  client, err := ethclient.Dial("https://mainnet.infura.io")
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println("we have a connection")
  _ = client

  var blockNum string
  //err = client.rpc.CallContext(ctx, &blockNum, "eth_blockNumber")
  //err = client.txpool.content() (map[string]map[string]map[string][]*RPCTransaction)

  if err != nil {
    return
  }

//  header, err := client.HeaderByNumber(context.Background(), nil)
//  if err != nil {
//    log.Fatal(err)
//  }

  blockNumber := big.NewInt(5671744)
  block, err := client.BlockByNumber(context.Background(), blockNumber)
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(len(block.Transactions()))   // 144

  count, err := client.TransactionCount(context.Background(), block.Hash())
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(count) // 144

  var tx *types.Transaction
  var isPending bool
  var findHash common.Hash

  findHash, _ = NewAddressFromHex("0xcb39dcefcd32268053fdcccc22cc678e44b19a25b7c4f8a40eb2bb7db713e6f1")
  //findHash, _ = hex.DecodeString(hashString)

  tx, isPending, err = client.TransactionByHash(context.Background(), findHash)

  fmt.Println("Output: %s", blockNum)
}
