package client

import (
	"encoding/json"
	"fmt"
	"github.com/hz7680/dingtalkSdk/model"
	"github.com/hz7680/dingtalkSdk/utils"
	"log"
)

func (c *Client) GetUserId(code string) (userIdResp *model.UserIdResp, err error) {
	token, err := c.GetAccessToken()
	if err != nil {
		log.Println("get access token error : " + err.Error())
		return
	}
	url := "https://oapi.dingtalk.com/user/getuserinfo?access_token=" + token + "&code=" + code
	body, err := utils.HttpGet(url)
	if err != nil {
		return
	}
	userIdResp = &model.UserIdResp{}
	err = json.Unmarshal(body, userIdResp)
	return
}

func (c *Client) GetUserDetail(userid string) (user *model.UserDetailResp, err error) {
	token, err := c.GetAccessToken()
	if err != nil {
		log.Println("get access token error : " + err.Error())
		return
	}
	url := "https://oapi.dingtalk.com/user/get?access_token=" + token + "&userid=" + userid
	body, err := utils.HttpGet(url)
	fmt.Println(string(body))
	if err != nil {
		return
	}
	user = &model.UserDetailResp{}
	err = json.Unmarshal(body, user)
	return
}
