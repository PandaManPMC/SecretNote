package encrypt

import (
	"SecretNote/tutil"
	"encoding/base64"
	"fmt"
	"testing"
	"time"
)

func TestAesEncrypt(t *testing.T){
	data := "0xc35c0c2ae2f4b6ba21dc2d0db73cbefdf4089b1216fb1e4504ba0b5b55776df2"
	key := tutil.RandCharacterString(32)
	result,err := AesEncrypt([]byte(data),[]byte(key))
	t.Log(err)
	t.Log(result)
	t.Log(string(result))

	d,err := AesDecrypt([]byte(string(result)),[]byte(key))
	t.Log(err)
	t.Log(d)
	t.Log(string(d))
	if data == string(d){
		t.Log("success")
	}

	br := base64.URLEncoding.EncodeToString(result)
	t.Log(br)
	dr,err := base64.URLEncoding.DecodeString(br)
	t.Log(err)
	t.Log(dr)
	d,err = AesDecrypt(dr,[]byte(key))
	t.Log(err)
	t.Log(d)
	if data == string(d){
		t.Log("success   22222")
	}
}

func TestAesEncryptToBase64(t *testing.T){
	data := "0xc35c0c2ae2f4b6ba21dc2d0db73cbefdf4089b1216fb1e4504ba0b5b55776df2"
	key := tutil.RandCharacterString(32)
	t.Log(fmt.Sprintf("data=%s  ***  key=%s",data,key))

	ciphertext,err := AesEncryptToBase64([]byte(data),key)
	t.Log(err)
	t.Log(*ciphertext)

	plaintext,err := AesDecryptByBase64(*ciphertext,key)
	t.Log(err)
	t.Log(*plaintext)

	if data == *plaintext {
		t.Log("success")
	}

}

type Model struct {
	SareId int64 `json:"sareId"`	// 管理员角色编号
	SanId int64 `json:"sanId"`	// 管理员编号
	SanPosition string `json:"sanPosition"`	// 职位
	CreateDate time.Time `json:"createDate"`	// 创建时间
}


