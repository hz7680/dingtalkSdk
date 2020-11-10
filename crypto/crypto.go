package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"
	"encoding/base64"
	"encoding/binary"
	"errors"
	"fmt"
	"math/rand"
	"sort"
	"time"
)

const (
	AES_ENCODE_KEY_LENGTH = 43
)

type Crypto struct {
	Token  string
	AesKey string
	Key    string
	block  cipher.Block
	bkey   []byte
}

/*
	token	随机串
	aesKey
	suiteKey 一般使用corpID
*/
func NewCrypto(token, aesKey, suiteKey string) (c *Crypto) {
	c = &Crypto{
		Token:  token,
		AesKey: aesKey,
		Key:    suiteKey,
	}
	if len(c.AesKey) != AES_ENCODE_KEY_LENGTH {
		panic("不合法的aeskey")
	}
	var err error
	c.bkey, err = base64.StdEncoding.DecodeString(aesKey + "=")
	if err != nil {
		panic(err.Error())
	}
	c.block, err = aes.NewCipher(c.bkey)
	if err != nil {
		panic(err.Error())
	}
	return c
}

func (c *Crypto) Ddbase64sign(key string) (bt []byte, err error) {
	return base64.StdEncoding.DecodeString(key)
}

func (c *Crypto) sha1Sign(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs)
}

func (c *Crypto) PKCS7UnPadding(plantText []byte) []byte {
	length := len(plantText)
	unpadding := int(plantText[length-1])
	return plantText[:(length - unpadding)]
}

func (c *Crypto) PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// 数据签名
func (c *Crypto) GenerateSignature(token, timeStamp, nonce, secretStr string) string {
	// 先将参数值进行排序
	params := make([]string, 0)
	params = append(params, token)
	params = append(params, secretStr)
	params = append(params, timeStamp)
	params = append(params, nonce)
	sort.Strings(params)
	return c.sha1Sign(params[0] + params[1] + params[2] + params[3])
}

/*
	signature: 签名字符串
	timeStamp: 时间戳
	nonce: 随机字符串
	secretStr: 密文
	返回: 解密后的明文
*/
func (c *Crypto) DecryptMsg(signature, timeStamp, nonce, secretStr string) (string, error) {
	if !c.VerifySignature(c.Token, timeStamp, nonce, secretStr, signature) {
		return "", errors.New("签名不匹配")
	}
	decode, err := base64.StdEncoding.DecodeString(secretStr)
	if err != nil {
		return "", err
	}
	if len(decode) < aes.BlockSize {
		return "", errors.New("密文太短啦")
	}
	blockMode := cipher.NewCBCDecrypter(c.block, c.bkey[:c.block.BlockSize()])
	plantText := make([]byte, len(decode))
	blockMode.CryptBlocks(plantText, decode)
	plantText = c.PKCS7UnPadding(plantText)
	size := binary.BigEndian.Uint32(plantText[16 : 16+4])
	plantText = plantText[16+4:]
	corpid := plantText[size:]

	if string(corpid) != c.Key {
		return "", errors.New("CorpID不正确")
	}
	return string(plantText[:size]), nil
}

/*
	replyMsg: 明文字符串
	timeStamp: 时间戳
	nonce: 随机字符串
	返回: 密文,签名字符串
*/
func (c *Crypto) EncryptMsg(replyMsg, timeStamp, nonce string) (encrypt string, signature string, err error) {
	//原生消息体长度
	size := make([]byte, 4)
	binary.BigEndian.PutUint32(size, uint32(len(replyMsg)))
	replyMsg = c.RandomString(16) + string(size) + replyMsg + c.Key
	plantText := c.PKCS7Padding([]byte(replyMsg), c.block.BlockSize())
	if len(plantText)%aes.BlockSize != 0 {
		return "", "", errors.New("消息体大小不为16的倍数")
	}

	blockMode := cipher.NewCBCEncrypter(c.block, c.bkey[:c.block.BlockSize()])
	ciphertext := make([]byte, len(plantText))
	blockMode.CryptBlocks(ciphertext, plantText)
	encrypt = base64.StdEncoding.EncodeToString(ciphertext)
	signature = c.GenerateSignature(c.Token, timeStamp, nonce, encrypt)
	return
}

// 校验数据签名
func (c *Crypto) VerifySignature(token, timeStamp, nonce, secretStr, sigture string) bool {
	return c.GenerateSignature(token, timeStamp, nonce, secretStr) == sigture
}

func (c *Crypto) RandomString(n int, alphabets ...byte) string {
	const alphanum = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, n)
	var randby bool
	if num, err := rand.Read(bytes); num != n || err != nil {
		rand.Seed(time.Now().UnixNano())
		randby = true
	}
	for i, b := range bytes {
		if len(alphabets) == 0 {
			if randby {
				bytes[i] = alphanum[rand.Intn(len(alphanum))]
			} else {
				bytes[i] = alphanum[b%byte(len(alphanum))]
			}
		} else {
			if randby {
				bytes[i] = alphabets[rand.Intn(len(alphabets))]
			} else {
				bytes[i] = alphabets[b%byte(len(alphabets))]
			}
		}
	}
	return string(bytes)
}