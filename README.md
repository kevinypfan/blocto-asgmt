# blocto-asgmt

[![Go Reference](https://pkg.go.dev/badge/golang.org/x/example.svg)](https://pkg.go.dev/golang.org/x/example)

## 架構說明

此專案經過不少的修改，從剛開始拿到 Block 後，拿 Transaction 再拿 logs 為了讓這能平行運行，嘗試使用了 goruntine 的各種方法，在`web3/crawl.go` 我也嘗試使用多個 worker 同時消耗 channel 上的目標，但是結果都不盡理想。所以我最後一次修改，將 Transaction 送至 Kafka 使用 Event Sourcing 的方式，在準備一隻 Comsumer `kafka/transaction/main.go` 使這個可以很好的做水平擴展，很好的加速 blocks 的爬取。

`api` 此為 RESTful API Router 主要路徑。

`db` Database 操作與 sqlc 套件相關路徑。

`kafka` Transaction 因為數量太多，後來將 Transaction 送至 Kafka 做 event sourcing，`kafka/transaction/main.go` 此為微服務。

`util` 環境參數與剛開始實驗 jsonrpc code。

## Run Project

```
$ docker compose up -d --scale service=6
```
