package client

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

type EthereumClient struct {
	RawUrl string
	Client *ethclient.Client
}

//	NewEthereumClient 创建以太坊客户端实例
//	string rwaUrl 连接的节点 rpc 地址
func NewEthereumClient(rwaUrl string) (*EthereumClient,error){
	client, err := ethclient.Dial("https://api.mycryptoapi.com/eth")
	if nil != err {
		return nil,err
	}
	ec := EthereumClient{RawUrl: rwaUrl,Client: client}
	return &ec,nil
}

//	GetBalance 获取余额
func (instance *EthereumClient) GetBalance(address string) (*big.Int, error){
	account := common.HexToAddress(address)
	balance, err := instance.Client.BalanceAt(context.Background(), account, nil)
	if nil != err{
		return nil,err
	}
	return balance,nil
}


