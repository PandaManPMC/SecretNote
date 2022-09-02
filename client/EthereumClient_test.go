package client

import (
	"SecretNote/tutil"
	"SecretNote/wallet"
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"math"
	"math/big"
	"testing"
	"time"
)

func TestGetBalance(t *testing.T) {
	//address := "0xC166fe1c611B4d396B131F788C01d011b0187bf6"
	//client, err := NewEthereumClient("https://rinkeby.infura.io/v3/16aa425866e742bea51efb11ccdbc900")

	client, err := NewEthereumClient("wss://rinkeby.infura.io/ws/v3/16aa425866e742bea51efb11ccdbc900")
	contractAddress := common.HexToAddress("0xd26C80A81a7a548EAeCAE649C6B174dCFfeC57DE") // rinkeby

	//client, err := NewEthereumClient("https://evmtestnet.confluxrpc.com")
	//client, err := NewEthereumClient("https://evm.confluxrpc.org")

	// cfx espace
	//client, err := NewEthereumClient("https://evm.confluxrpc.com")
	//contractAddress := common.HexToAddress("0x752D95a870C5F7aeddF084746C3E2c6974C5319f")

	ctx := context.Background()
	if nil != err {
		t.Fatal(err)
	}
	chainID, _ := client.Client.ChainID(ctx)
	t.Log(chainID)

	//0x861caF1c18feE1F3e1De005C40f322e9c3eA774E
	address, privateKey, pubkey, err := wallet.ImportAccountByMnemonic("bicycle gate quantum gather forget tattoo grocery bottom remember half animal art", 0)
	t.Log(*privateKey)
	t.Logf("地址0=%s \n 私钥=%s \n 公钥=%s \n", *address, *privateKey, *pubkey)
	if nil != err {
		t.Fatal(err)
	}

	bal, err := client.GetBalance(*address)
	if nil != err {
		t.Fatal(err)
	}
	eth := WeiToEth(bal)
	t.Logf("余额 %v wei\t %v eth\n", bal, eth)

	t.Log(contractAddress)
	contract, err := NewCoin(contractAddress, client.Client)
	if nil != err {
		t.Fatal(err)
	}

	bal2, err2 := contract.BalanceOf(nil, common.HexToAddress(*address))
	t.Log(bal2)
	t.Log(err2)

	// 导入钱包
	pkey, err := crypto.HexToECDSA(*privateKey)
	if nil != err {
		t.Fatal(err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(pkey, chainID)
	if nil != err {
		t.Fatal(err)
	}

	ctxt := context.Background()
	gas, err := client.Client.SuggestGasPrice(ctxt)
	if nil != err {
		log.Fatal(err)
	}
	t.Log(gas)
	nonce, err := client.Client.PendingNonceAt(ctxt, auth.From)
	t.Log(nonce)
	if nil != err {
		log.Fatal(err)
	}

	go func() {
		ctxWatch := context.Background()
		opts := &bind.WatchOpts{}
		opts.Context = ctxWatch
		sink := make(chan *CoinTransfer)
		from := make([]common.Address, 0)
		to := make([]common.Address, 0)
		sub, err := contract.WatchTransfer(opts, sink, from, to)

		//client.Client.SubscribeFilterLogs(ctx)

		if nil != err {
			t.Fatal(tutil.GetInstanceByReflectUtil().ErrToMap(err))
		}
		t.Log(sub)
		for {
			select {
			case v, isOpen := <-sink:
				t.Log("发生了转账事件")
				t.Log(isOpen)
				t.Log(v)
			case v, isOpen := <-sub.Err():
				t.Log(isOpen)
				t.Log(v)
			}
		}
	}()

	t.Log(auth)
	tops := &bind.TransactOpts{}
	tops.From = auth.From
	tops.Signer = auth.Signer
	tops.Value = nil
	//tops.GasLimit = 15000000
	tops.GasPrice = gas
	//tops.Nonce = big.NewInt(int64(nonce + 1))
	t.Log(tops)

	d := int64(math.Pow10(18))
	de := big.NewInt(d)
	val := big.NewInt(200)
	toVal := de.Mul(val, de)
	t.Log(toVal)
	hash, err := contract.Transfer(tops, common.HexToAddress("0x98647263f8e52F755a6ba22FC3325AEF180f1289"), toVal)
	if nil != err {
		t.Log(tutil.GetInstanceByReflectUtil().ErrToMap(err))
		t.Fatal(err)
	}
	t.Log(hash)
	t.Log(hash.Hash())

	t.Log(time.Now().Format("2006.01.02 15:04:05"))
	time.Sleep(time.Second * 15 * 8)
	t.Log(time.Now().Format("2006.01.02 15:04:05"))
	receipt, err := client.Client.TransactionReceipt(ctx, hash.Hash())
	if nil != err {
		t.Fatal(err)
	}
	t.Log(receipt)
	t.Log(receipt.Status)
	t.Log(receipt.Type)

	t.Log("---------------------------- end -----------------------------")
}
