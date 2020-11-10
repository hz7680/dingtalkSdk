package client

import (
	"encoding/json"
	"fmt"
	"github.com/hz7680/dingtalkSdk/model"
	"github.com/hz7680/dingtalkSdk/utils"
	"log"
)

func (c *Client) GetReportTemplateListByUserid(userid string, offset, size int) {
	token, err := c.GetAccessToken()
	if err != nil {
		log.Println("get access token error : " + err.Error())
		return
	}
	url := "https://oapi.dingtalk.com/topapi/report/template/listbyuserid?access_token=" + token
	json, err := utils.HttpPostJson(url, map[string]interface{}{
		"userid": userid,
		"offset": offset,
		"size":   size,
	})
	fmt.Println(string(json))
}

func (c *Client) GetReportSimplelist(startTime, endTime int64, templateName, userid string, cursor, size int) (list *model.ReportSimplelist, err error) {
	token, err := c.GetAccessToken()
	if err != nil {
		log.Println("get access token error : " + err.Error())
		return
	}
	url := "https://oapi.dingtalk.com/topapi/report/simplelist?access_token=" + token
	body, err := utils.HttpPostJson(url, map[string]interface{}{
		"start_time":    startTime,
		"end_time":      endTime,
		"template_name": templateName,
		"userid":        userid,
		"cursor":        cursor,
		"size":          size,
	})
	if err != nil {
		return
	}

	list = &model.ReportSimplelist{}
	err = json.Unmarshal(body, list)
	return
}

func (c *Client) GetReportlist(startTime, endTime int64, templateName, userid string, cursor, size int) (list *model.ReportList, err error) {
	token, err := c.GetAccessToken()
	if err != nil {
		log.Println("get access token error : " + err.Error())
		return
	}
	url := "https://oapi.dingtalk.com/topapi/report/list?access_token=" + token
	body, err := utils.HttpPostJson(url, map[string]interface{}{
		"start_time":    startTime,
		"end_time":      endTime,
		"template_name": templateName,
		"userid":        userid,
		"cursor":        cursor,
		"size":          size,
	})
	if err != nil {
		return
	}

	list = &model.ReportList{}
	err = json.Unmarshal(body, list)
	return
}
