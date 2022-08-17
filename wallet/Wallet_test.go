package wallet

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/tyler-smith/go-bip39"
	"strings"
	"testing"
)

func TestCreateWallet(t *testing.T) {
	entropy, err := bip39.NewEntropy(128)
	//entropy, err := bip39.NewEntropy(256)

	if err != nil {
		t.Fatal(err)
	}
	mnemonic, _ := bip39.NewMnemonic(entropy)
	fmt.Println("mnemonic:", mnemonic)
	fmt.Println(len(strings.Split(mnemonic, " ")))
	seed := bip39.NewSeed(mnemonic, "")

	wallet, err := hdwallet.NewFromSeed(seed)
	if err != nil {
		t.Fatal(err)
	}
	path := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/0")
	account, err := wallet.Derive(path, false)
	if err != nil {
		t.Fatal(err)
	}

	address := account.Address.Hex()
	privateKey, _ := wallet.PrivateKeyHex(account)
	publicKey, _ := wallet.PublicKeyHex(account)

	fmt.Println("address0:", address)
	fmt.Println("privateKey:", privateKey)
	fmt.Println("publicKey:", publicKey)

	path = hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/1") //生成id为1的钱包地址
	account, err = wallet.Derive(path, false)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("address1:", account.Address.Hex())
}

func TestCreateAccount(t *testing.T) {
	address0, privateKey, mnemonic, publicKey, err := CreateAccount()
	if nil != err {
		t.Fatal(err)
	}
	t.Logf("地址0=%s\n私钥=%s\n助记词=%s\n公钥=%s\n", *address0, *privateKey, *mnemonic, *publicKey)

	pk, err := crypto.HexToECDSA(*privateKey)
	if nil != err {
		t.Fatal(err)
	}
	t.Log(pk)

	b, _ := hexutil.Decode(*publicKey)
	prvK, err := crypto.DecompressPubkey(b)
	if nil != err {
		t.Fatal(err)
	}
	addr := crypto.PubkeyToAddress(*prvK)
	t.Log(addr)
}

func TestImportAccount(t *testing.T) {
	//地址0=0x97Fa4A2d3C3df5b3b05F4B9265F438c3a8B064eb
	//私钥=ec8da26e3e59fa40ac0410fb09df8ff05bd7ebad5ef5fa89ffeea7408ac9739f
	//助记词=sister pyramid polar oyster describe empty unknown night ill youth arrow awkward
	//公钥=8ea0bc55ea26af995f760edde7a07c86faa6404d8f9dc9915b8b9ea30c058f8fbe2172315265a21388185f7ce34db0ddc765401e0a310cd9a609458aa139091c
	mnemonic := "sister pyramid polar oyster describe empty unknown night ill youth arrow awkward"
	address, privateKey, pubkey, err := ImportAccountByMnemonic(mnemonic, 0)
	if nil != err {
		t.Fatal(err)
	}
	t.Logf("地址0=%s\n私钥=%s\n公钥=%s\n助记词=%s\n", *address, *privateKey, *pubkey, mnemonic)
}
