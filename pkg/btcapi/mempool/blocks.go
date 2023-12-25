package mempool

import (
	"fmt"
	"net/http"
)

func (c *MempoolClient) GetBlockHash(blockHeight int) (string, error) {
	res, err := c.request(http.MethodGet, fmt.Sprintf("/block-height/%d", blockHeight), nil)
	if err != nil {
		return "", err
	}

	return string(res), nil
}

func (c *MempoolClient) GetTransactions(blockHash string) (string, error) {
	res, err := c.request(http.MethodGet, fmt.Sprintf("/block/%s/txs", blockHash), nil)
	if err != nil {
		return "", err
	}

	return string(res), nil
}
