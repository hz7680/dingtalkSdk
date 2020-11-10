package model

import "time"

type ProcessInstanceDetail struct {
	Errcode         int             `json:"errcode"`
	Errmsg          string          `json:"errmsg"`
	ProcessInstance ProcessInstance `json:"process_instance"`
}

type ProcessInstance struct {
	Title                      string                `json:"title"`
	CreateTime                 *time.Time           `json:"create_time"`
	FinishTime                 *time.Time           `json:"finish_time"`
	OriginatorUserid           string                `json:"originator_userid"`
	OriginatorDeptId           string                `json:"originator_dept_id"`
	Status                     string                `json:"status"`
	CcUserids                  []string              `json:"cc_userids"`
	FormComponentValues        []*FormComponentValue `json:"form_component_values"`
	Result                     string                `json:"result"`
	BusinessId                 string                `json:"business_id"`
	OperationRecords           []*OperationRecord    `json:"operation_records"`
	Tasks                      []*Task               `json:"tasks"`
	OriginatorDeptName         string                `json:"originator_dept_name"`
	BizAction                  string                `json:"biz_action"`
	AttachedProcessInstanceIds []interface{}         `json:"attached_process_instance_ids"`
}

type FormComponentValue struct {
	Name     string `json:"name"`
	Value    string `json:"value"`
	ExtValue string `json:"ext_value"`
}

type OperationRecord struct {
	Userid          string      `json:"userid"`
	Date            *time.Time `json:"date"`
	OperationType   string      `json:"operation_type"`
	OperationResult string      `json:"operation_result"`
	Remark          string      `json:"remark"`
}

type Task struct {
	Userid     string      `json:"userid"`
	TaskStatus string      `json:"task_status"`
	TaskResult string      `json:"task_result"`
	CreateTime *time.Time `json:"create_time"`
	FinishTime *time.Time `json:"finish_time"`
	Taskid     string      `json:"taskid"`
}

type ProcessInstanceListResp struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	Result  struct {
		List       []string `json:"list"`
		NextCursor int      `json:"next_cursor"`
	} `json:"result"`
}
