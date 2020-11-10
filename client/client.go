package client

import (
	crypto2 "github.com/hz7680/dingtalkSdk/crypto"
	"github.com/hz7680/dingtalkSdk/model"
)

type Client struct {
	corpId      string
	agentId     string
	appKey      string
	appSecret   string
	aesKey      string
	key         string
	token       string
	accessToken *model.AccessToken
	crypto      *crypto2.Crypto
}

func NewClient(corpId, agentId, appKey, appSecret, aesKey, key, token string) *Client {
	return &Client{
		corpId:      corpId,
		agentId:     agentId,
		appKey:      appKey,
		appSecret:   appSecret,
		aesKey:      aesKey,
		key:         key,
		token:       token,
		accessToken: model.NewAccessToken(),
		crypto:      crypto2.NewCrypto(token, aesKey, key),
	}
}
