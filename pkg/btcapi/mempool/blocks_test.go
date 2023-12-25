package mempool

import (
	"testing"

	"github.com/btcsuite/btcd/chaincfg"
)

func TestGetBlockHash(t *testing.T) {
	client := NewClient(&chaincfg.MainNetParams)
	blockHash, err := client.GetBlockHash(822823)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(blockHash)
}

func TestGetTransactions(t *testing.T) {
	client := NewClient(&chaincfg.MainNetParams)
	transactions, err := client.GetTransactions("000000000000000000024839645b0a5a078b450771bd7e8fe3355389bbb13074")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(transactions)
}
