package encrypt

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

//	Sign 签名
//	return *string sigHex 签名，如果 error 不为 nil 则失败
func Sign(data ,privateKey string) (*string,error){
	dataByte := crypto.Keccak256([]byte(data))
	pkey,err := crypto.HexToECDSA(privateKey)
	if nil != err {
		return nil,err
	}
	signature, err := crypto.Sign(dataByte, pkey)
	if nil != err {
		return nil,err
	}
	sigHex := hexutil.Encode(signature)
	return &sigHex,err
}

//	VerifySignature 验证签名
//	return *string address 钱包地址, *bool 验签成功 true,error 不为 nil 则出现异常
func VerifySignature(data,signature string) (*string,*bool,error){
	dataHash := crypto.Keccak256([]byte(data))
	sigByte,err := hexutil.Decode(signature)
	if nil != err {
		return nil,nil,err
	}

	recoveredPub, err := crypto.Ecrecover(dataHash, sigByte)
	if nil != err {
		return nil,nil,err
	}
	recoveredPubKey, _ := crypto.UnmarshalPubkey(recoveredPub)
	address := crypto.PubkeyToAddress(*recoveredPubKey).String()

	publicKeyBytes := crypto.FromECDSAPub(recoveredPubKey)
	ok := crypto.VerifySignature(publicKeyBytes,dataHash[:],sigByte[:len(sigByte)-1])
	return &address,&ok,nil
}



