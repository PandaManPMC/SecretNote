package client

import (
	"testing"
)

func TestGetBalance(t *testing.T){
	//address := "0xC166fe1c611B4d396B131F788C01d011b0187bf6"
	address := "0x4aF0c2A6FAbC3a1E9C2B1f357B316C141Ba5d651"
	client,err := NewEthereumClient("https://5835a8a3d9834e9ab83d2a31842d9664.eth.rpc.rivet.cloud/")
	if nil != err {
		t.Fatal(err)
	}
	bal,err := client.GetBalance(address)
	if nil != err {
		t.Fatal(err)
	}
	eth := WeiToEth(bal)
	t.Logf("余额 %v wei\t %v eth\n",bal,eth)
}



