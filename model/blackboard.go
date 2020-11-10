package model

import "time"

type BlackboardListResp struct {
	Errcode        int           `json:"errcode"`
	Errmsg         string        `json:"errmsg"`
	BlackboardList []*Blackboard `json:"blackboard_list"`
}

type Blackboard struct {
	GmtCreate *time.Time `json:"gmt_create"`
	Title     string      `json:"title"`
	Url       string      `json:"url"`
}
