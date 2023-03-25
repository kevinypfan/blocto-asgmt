package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type RPCSendData struct {
	Jsonrpc string `json:"jsonrpc"`
	Id      int    `json:"id"`
	Method  string `json:"method"`
	Params  []any  `json:"params"`
}

type RPCResponse[T any] struct {
	Jsonrpc string `json:"jsonrpc"`
	Id      int    `json:"id"`
	Result  T      `json:"result"`
}

type Transaction struct {
	BlockHash        string `json:"blockHash"`
	BlockNumber      string `json:"blockNumber"`
	From             string `json:"from"`
	Gas              string `json:"gas"`
	GasPrice         string `json:"gasPrice"`
	Hash             string `json:"hash"`
	Input            string `json:"input"`
	Nonce            string `json:"nonce"`
	To               string `json:"to"`
	TransactionIndex string `json:"transactionIndex"`
	Value            string `json:"value"`
	TransactionType  string `json:"type"`
	V                string `json:"v"`
	R                string `json:"r"`
	S                string `json:"s"`
}

type Block struct {
	Difficulty       string   `json:"difficulty"`
	ExtraData        string   `json:"extraData"`
	GasLimit         string   `json:"gasLimit"`
	GasUsed          string   `json:"gasUsed"`
	Hash             string   `json:"hash"`
	LogsBloom        string   `json:"logsBloom"`
	Miner            string   `json:"miner"`
	MixHash          string   `json:"mixHash"`
	Nonce            string   `json:"nonce"`
	Number           string   `json:"number"`
	ParentHash       string   `json:"parentHash"`
	ReceiptsRoot     string   `json:"receiptsRoot"`
	Sha3Uncles       string   `json:"sha3Uncles"`
	Size             string   `json:"size"`
	StateRoot        string   `json:"stateRoot"`
	Timestamp        string   `json:"timestamp"`
	TotalDifficulty  string   `json:"totalDifficulty"`
	TransactionsRoot string   `json:"transactionsRoot"`
	Transactions     []string `json:"transactions"`
	Uncles           []string `json:"uncles"`
}

func postJsonRpc(payload string) []byte {
	client := &http.Client{}
	url := "https://data-seed-prebsc-2-s3.binance.org:8545/"
	method := "POST"

	data := strings.NewReader(payload)

	req, err := http.NewRequest(method, url, data)

	if err != nil {
		fmt.Println(err)
		return nil
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return body
}

func GetBlockByNumber(tag string) {
	sendData := RPCSendData{}
	sendData.Jsonrpc = "2.0"
	sendData.Id = 1
	sendData.Params = []any{tag, false}
	sendData.Method = "eth_getBlockByNumber"

	payload, err := json.Marshal(sendData)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(payload))

	resBody := postJsonRpc(string(payload))

	fmt.Println(string(resBody))

	var RpcBlock RPCResponse[Block]

	json.Unmarshal(resBody, &RpcBlock)

	fmt.Println(RpcBlock.Result.Hash)
}
