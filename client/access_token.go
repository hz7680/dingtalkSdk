package client

import (
	"encoding/json"
	"errors"
	"github.com/hz7680/dingtalkSdk/model"
	"github.com/hz7680/dingtalkSdk/utils"
)

func (c *Client) GetAccessToken() (token string, err error) {
	if c.accessToken.IsExpired() {
		err = c.RefreshAccessToken()
		if err != nil {
			return
		}
	}
	c.accessToken.RLock()
	defer c.accessToken.RUnlock()

	return c.accessToken.GetAccessToken(), nil
}

func (c *Client) RefreshAccessToken() error {
	c.accessToken.Lock()
	defer c.accessToken.UnLock()

	if !c.accessToken.IsExpired() {
		return nil
	}

	url := "https://oapi.dingtalk.com/gettoken?appkey=" + c.appKey + "&appsecret=" + c.appSecret
	body, err := utils.HttpGet(url)
	if err != nil {
		return err
	}
	resp := &model.AccessTokenResp{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		return err
	}
	if resp.Errcode != 0 {
		err = errors.New(resp.Errmsg)
		return err
	}
	c.accessToken.SetAccessToken(resp.AccessToken, resp.ExpiresIn)
	return nil
}
