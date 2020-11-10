package client

import (
	"encoding/json"
	"errors"
	"github.com/hz7680/dingtalkSdk/model"
	"github.com/hz7680/dingtalkSdk/utils"
)

func (c *Client) DeleteCallback() (resp *model.DeleteCallbackResp, err error) {
	token, err := c.GetAccessToken()
	if err != nil {
		return
	}
	url := "https://oapi.dingtalk.com/call_back/delete_call_back?access_token=" + token
	body, err := utils.HttpGet(url)
	if err != nil {
		return
	}
	resp = &model.DeleteCallbackResp{}
	err = json.Unmarshal(body, resp)
	return
}

func (c *Client) RegisterCallback(callbackUrl string, tags []model.CallbackTag) (resp *model.RegisterCallbackResp, err error) {
	token, err := c.GetAccessToken()
	if err != nil {
		return
	}
	url := "https://oapi.dingtalk.com/call_back/register_call_back?access_token=" + token

	body, err := utils.HttpPostJson(url, map[string]interface{}{
		"call_back_tag": tags,
		"token":         c.token,
		"aes_key":       c.aesKey,
		"url":           callbackUrl,
	})
	if err != nil {
		return
	}
	resp = &model.RegisterCallbackResp{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		return
	}
	if resp.Errcode != 0 {
		err = errors.New(resp.Errmsg)
		return
	}
	return
}

func (c *Client) Decrypt(callBack *model.CallBackReq) (decryptMsg map[string]interface{}, err error) {
	decryptMsgJson, err := c.crypto.DecryptMsg(callBack.Signature, callBack.TimeStamp, callBack.Nonce, callBack.Encrypt)
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(decryptMsgJson), &decryptMsg)
	if err != nil {
		return
	}
	return
}

func (c *Client) Encrypt(reply string, back *model.CallBackReq) (encrypt string, signature string, err error) {
	return c.crypto.EncryptMsg(reply, back.TimeStamp, back.Nonce)
}

func (c *Client) GenerateSuccessResp(callback *model.CallBackReq) (resp *model.CallBackResp, err error) {
	encrypt, signature, err := c.Encrypt("success", callback)
	if err != nil {
		return
	}
	resp = &model.CallBackResp{
		MsgSignature: signature,
		TimeStamp:    callback.TimeStamp,
		Nonce:        callback.Nonce,
		Encrypt:      encrypt,
	}
	return
}