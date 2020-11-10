package client

import (
	"encoding/json"
	"fmt"
	"github.com/hz7680/dingtalkSdk/model"
	"github.com/hz7680/dingtalkSdk/utils"
	"log"
)

func (c *Client) GetDepartmentList(fetchChild bool,parentId string) (list *model.DepartmentListResp,err error) {
	token, err := c.GetAccessToken()
	if err != nil{
		log.Println("get access token error : "+err.Error())
		return
	}
	url:=fmt.Sprintf("https://oapi.dingtalk.com/department/list?access_token=%s&fet_child=%t&id=%s",token,fetchChild,parentId)
	body, err := utils.HttpGet(url)
	if err != nil{
		log.Println(err)
		return
	}
	list = &model.DepartmentListResp{}
	err = json.Unmarshal(body, list)
	if err != nil{
		return
	}
	return
}