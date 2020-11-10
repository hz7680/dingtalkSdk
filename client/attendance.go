package client

import (
	"encoding/json"
	"github.com/hz7680/dingtalkSdk/model"
	"github.com/hz7680/dingtalkSdk/utils"
	"log"
	"time"
)

func (c *Client) GetScheduleList(workDate *time.Time, offset int) (resp *model.ScheduleListResp, err error) {
	token, err := c.GetAccessToken()
	if err != nil {
		log.Println("get access token error : " + err.Error())
		return
	}
	url := "https://oapi.dingtalk.com/topapi/attendance/listschedule?access_token=" + token
	body, err := utils.HttpPostJson(url, map[string]interface{}{
		"workDate": workDate,
		"offset":   offset,
	})
	if err != nil {
		return
	}

	resp = &model.ScheduleListResp{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		return
	}
	return
}

func (c *Client) GetAttendanceList(workDateFrom, workDateTo string, userIdList []string, offset, limit int) (result *model.AttendanceList, err error) {
	token, err := c.GetAccessToken()
	if err != nil {
		log.Println("get access token error : " + err.Error())
		return
	}
	url := "https://oapi.dingtalk.com/attendance/list?access_token=" + token
	body, err := utils.HttpPostJson(url, map[string]interface{}{
		"workDateFrom": workDateFrom,
		"workDateTo":   workDateTo,
		"userIdList":   userIdList,
		"offset":       offset,
		"limit":        limit,
	})
	if err != nil {
		return
	}
	result = &model.AttendanceList{}
	err = json.Unmarshal(body, result)
	return
}
