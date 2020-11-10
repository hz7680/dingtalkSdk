package client

import (
	"encoding/json"
	"github.com/hz7680/dingtalkSdk/model"
	"github.com/hz7680/dingtalkSdk/utils"
	"log"
)

func (c *Client) GetBlackboardList(userid string) (list *model.BlackboardListResp, err error) {
	token, err := c.GetAccessToken()
	if err != nil {
		log.Println("get access token error : " + err.Error())
		return
	}
	url := "https://oapi.dingtalk.com/topapi/blackboard/listtopten?access_token=" + token
	body, err := utils.HttpPostJson(url, map[string]interface{}{
		"userid": userid,
	})
	if err != nil{
		return
	}
	list = &model.BlackboardListResp{}
	err = json.Unmarshal(body, list)
	return
}
