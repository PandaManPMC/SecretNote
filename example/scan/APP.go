package main

import (
	"SecretNote/client"
	"SecretNote/wallet"
	"fmt"
	"sync/atomic"
)

var count uint64 = 0
var stop uint32 = 0
var ethClientList []*client.EthereumClient

func main() {
	version := "v.1.0.5"
	fmt.Printf("红岸神教，战无不胜，红岸教主，文成武德，千秋万载，一统江湖。\n 生成私钥用钱包地址查询余额 %s，找到就停止查询并打印助记词。\n",version)
	chStop := make(chan int,1)
	ethClientList = make([]*client.EthereumClient,0,0)

	rpcURIList := [...]string{
		"https://mainnet-nethermind.blockscout.com/",
		"https://cloudflare-eth.com/",
		"https://api.mycryptoapi.com/eth",
		"https://nodes.mewapi.io/rpc/eth",
		"https://eth-mainnet.token.im",
		"https://web3.1inch.exchange/",
		"https://mainnet.eth.cloud.ava.do/",
		"https://eth-mainnet.gateway.pokt.network/v1/5f3453978e354ab992c4da79",
		"https://5835a8a3d9834e9ab83d2a31842d9664.eth.rpc.rivet.cloud/"}

	for i:=0;i<len(rpcURIList);i++{
		ce,err := client.NewEthereumClient(rpcURIList[i])
		if nil != err {
			fmt.Printf("连接以太坊 RPC:%s 节点失败\n",rpcURIList[i])
			fmt.Println(err)
			continue
		}
		ethClientList = append(ethClientList,ce)
	}

	for i:=0;i<1;i++ {
		go wanWan(chStop)
	}

	<- chStop
	fmt.Println("成功找到一个有余额的账户，所有协程停止。")
	atomic.AddUint32(&stop,1)
}

func wanWan(chStop chan<- int){
	clientLen := uint64(len(ethClientList))
	for {
		if 0 < stop {
			return
		}
		ce := ethClientList[count % clientLen]
		atomic.AddUint64(&count,1)
		fmt.Printf("第%d次查询节点：%s\n",count,ce.RawUrl)
		address0,privateKey,mnemonic,_,err := wallet.CreateAccount()
		if nil != err {
			fmt.Println(err)
			continue
		}
		bal,err := ce.GetBalance(*address0)
		if nil != err {
			fmt.Println("查询余额出现错误")
			fmt.Println(err)
			continue
		}
		fmt.Printf("查询 %s 助记词 %s 余额 %v\n",*address0,*mnemonic,bal)
		if bal.Int64() > 0{
			fmt.Printf("查询到账户%s，私钥%s，助记词%s，有余额：%v wei\n",*address0,*privateKey,*mnemonic,bal)
			chStop <- 1
			return
		}
	}
}
