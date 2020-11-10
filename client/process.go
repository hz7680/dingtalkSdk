package client

import (
	"encoding/json"
	"github.com/hz7680/dingtalkSdk/model"
	"github.com/hz7680/dingtalkSdk/utils"
	"log"
)

func (c *Client) GetProcessInstance(instanceId string) (instance *model.ProcessInstanceDetail, err error) {
	token, err := c.GetAccessToken()
	if err != nil {
		log.Println("get accessToken error : " + err.Error())
		return
	}
	url := "https://oapi.dingtalk.com/topapi/processinstance/get?access_token=" + token
	body, err := utils.HttpPostJson(url, map[string]interface{}{
		"process_instance_id": instanceId,
	})
	if err != nil {
		log.Println("get process instance error : " + err.Error())
		return
	}
	instance = &model.ProcessInstanceDetail{}
	err = json.Unmarshal(body, instance)
	if err != nil {
		log.Println("process instance parse json error : " + err.Error())
		return
	}
	return
}

func (c *Client) GetProcessIntanceList(processCode string, startTime int64, endTime int) (resp *model.ProcessInstanceListResp, err error) {
	token, err := c.GetAccessToken()
	if err != nil {
		log.Println("get accessToken error : " + err.Error())
		return
	}
	url := "https://oapi.dingtalk.com/topapi/processinstance/listids?access_token=" + token
	body, err := utils.HttpPostJson(url, map[string]interface{}{
		"process_code": processCode,
		"start_time":   startTime,
		"end_time":     endTime,
	})
	if err != nil {
		log.Println("get process instance list error : " + err.Error())
		return
	}
	resp = &model.ProcessInstanceListResp{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		log.Println("process instance list parse json error : " + err.Error())
		return
	}
	return
}
