package model

import "time"

type ScheduleListResp struct {
	Errmsg  string `json:"errmsg"`
	Errcode int    `json:"errcode"`
	Result  struct {
		Schedules []*Schedule `json:"schedules"`
		HasMore   bool        `json:"has_more"`
	} `json:"result"`
}

type Schedule struct {
	PlanId           int         `json:"plan_id"`
	CheckType        string      `json:"check_type"`
	ApproveId        int         `json:"approve_id"`
	Userid           string      `json:"userid"`
	ClassId          int         `json:"class_id"`
	ClassSettingId   int         `json:"class_setting_id"`
	PlanCheckTime    *time.Time `json:"plan_check_time"`
	GroupId          int         `json:"group_id"`
	ChangedCheckTime *time.Time `json:"changed_check_time"`
}

type AttendanceList struct {
	Errcode      int                 `json:"errcode"`
	Errmsg       string              `json:"errmsg"`
	HasMore      bool                `json:"hasMore"`
	Recordresult []*AttendanceResult `json:"recordresult"`
}

type AttendanceResult struct {
	BaseCheckTime  int    `json:"baseCheckTime"`
	CheckType      string `json:"checkType"`
	GroupId        int    `json:"groupId"`
	Id             int    `json:"id"`
	LocationResult string `json:"locationResult"`
	PlanId         int    `json:"planId"`
	RecordId       int    `json:"recordId"`
	TimeResult     string `json:"timeResult"`
	UserCheckTime  int    `json:"userCheckTime"`
	UserId         string `json:"userId"`
	WorkDate       int    `json:"workDate"`
	ProcInstId     string `json:"procInstId"`
}
