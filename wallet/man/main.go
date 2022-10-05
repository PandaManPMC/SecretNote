package main

import (
	"fmt"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/tyler-smith/go-bip39"
	"strings"
)

//  author: laoniqiu
//  since: 2022/10/5
//  desc: man

func main() {
	fmt.Println("软件完全离线使用，有问题可发邮件 pandamancoin@outlook.com")
	fmt.Println("基于 big39 生成助记词和私钥以及0地址和 1地址钱包，导入 metamask 直接使用。")
	entropy, err := bip39.NewEntropy(128)
	//entropy, err := bip39.NewEntropy(256)

	if err != nil {
		fmt.Println(err)
		return
	}
	mnemonic, _ := bip39.NewMnemonic(entropy)
	fmt.Println(fmt.Sprintf("助记词有%d个", len(strings.Split(mnemonic, " "))))
	fmt.Println("mnemonic 助记词:", mnemonic)
	seed := bip39.NewSeed(mnemonic, "")

	wallet, err := hdwallet.NewFromSeed(seed)
	if err != nil {
		fmt.Println(err)
		return
	}
	path := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/0")
	account, err := wallet.Derive(path, false)
	if err != nil {
		fmt.Println(err)
		return
	}

	address := account.Address.Hex()
	privateKey, _ := wallet.PrivateKeyHex(account)
	publicKey, _ := wallet.PublicKeyHex(account)

	fmt.Println("privateKey 私钥:", privateKey)
	fmt.Println("publicKey 公钥:", publicKey)
	fmt.Println("address0 第一个钱包地址:", address)

	path = hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/1") //生成id为1的钱包地址
	account, err = wallet.Derive(path, false)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("address1 第二个钱包地址:", account.Address.Hex())
}
