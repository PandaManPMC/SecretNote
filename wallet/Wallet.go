package wallet

import (
	"encoding/hex"
	"fmt"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/tyler-smith/go-bip39"
)

//	CreateAccount 创建账户
//	return：
//	*string address0 首个钱包地址
//	*string privateKey 私钥
//	*string mnemonic 助记词
//	*string publicKey 公钥
//	error	不为 nil 时 账户创建失败
func CreateAccount() (*string,*string,*string,*string,error){
	entropy, err := bip39.NewEntropy(128)
	if nil != err {
		return nil,nil,nil,nil,err
	}
	mnemonic,err := bip39.NewMnemonic(entropy)
	if nil != err {
		return nil,nil,nil,nil,err
	}
	// password 可以传入指定密码或者空字符串，不同密码生成的助记词不同
	seed := bip39.NewSeed(mnemonic, "")
	wallet, err := hdwallet.NewFromSeed(seed)
	if nil != err {
		return nil,nil,nil,nil,err
	}
	// 最后一位是同一个助记词的地址id，从0开始相同助记词可以生产无限个地址
	path := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/0")
	account, err := wallet.Derive(path, false)
	if nil != err {
		return nil,nil,nil,nil,err
	}

	address := account.Address.Hex()
	privateKey, _ := wallet.PrivateKeyHex(account)

	publicKeyByte, _ := wallet.PublicKeyBytes(account)
	publicKey := hex.EncodeToString(publicKeyByte)
	return &address,&privateKey,&mnemonic,&publicKey,nil
}

//	ImportAccountByMnemonic 根据助记词导入账户
//	mnemonic 助记词，sequence 从 0 开始
//	return sequence 对应的地址、私钥、公钥，如果 error 不为 nil 则错误
func ImportAccountByMnemonic(mnemonic string,sequence uint8) (*string,*string,*string,error){
	seed := bip39.NewSeed(mnemonic, "")
	wallet, err := hdwallet.NewFromSeed(seed)
	if nil != err {
		return nil,nil,nil,err
	}
	path := hdwallet.MustParseDerivationPath(fmt.Sprintf("m/44'/60'/0'/0/%d",sequence))
	account, err := wallet.Derive(path, false)
	if nil != err {
		return nil,nil,nil,err
	}
	address := account.Address.Hex()
	privateKey, _ := wallet.PrivateKeyHex(account)
	publicKeyByte, _ := wallet.PublicKeyBytes(account)
	publicKey := hex.EncodeToString(publicKeyByte)
	return &address,&privateKey,&publicKey,nil
}