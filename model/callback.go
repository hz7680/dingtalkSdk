package model

type CallBackReq struct {
	Signature string `form:"signature"`
	TimeStamp string `form:"timestamp"`
	Nonce     string `form:"nonce"`
	Encrypt   string `json:"encrypt"`
}

type RegisterCallbackResp struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

type DeleteCallbackResp struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

type CallbackTag string

type CallBackResp struct {
	MsgSignature string `json:"msg_signature"`
	TimeStamp    string `json:"timeStamp"`
	Nonce        string `json:"nonce"`
	Encrypt      string `json:"encrypt"`
}

type CheckUrlEvent struct {
	EventType string `json:"EventType"`
}

type BpmsInstanceChangeEvent struct {
	EventType         string `c:"EventType"`
	ProcessInstanceId string `c:"processInstanceId"`
	FinishTime        int64  `c:"finishTime"`
	CorpId            string `c:"corpId"`
	Title             string `c:"title"`
	Type              string `c:"type"`
	Url               string `c:"url"`
	Result            string `c:"result"`
	CreateTime        int64  `c:"createTime"`
	StaffId           string `c:"staffId"`
	ProcessCode       string `c:"processCode"`
}

type UserExtValue struct {
	Nick         string `json:"nick"`
	OrgUserName  string `json:"orgUserName"`
	EmplId       string `json:"emplId"`
	Name         string `json:"name"`
	Avatar       string `json:"avatar"`
	SelectDeptId int    `json:"selectDeptId"`
}
type DeptExtValue struct {
	Number int    `json:"number"`
	Name   string `json:"name"`
	Id     int    `json:"id"`
}

type AttendanceCheckRecordEvent struct {
	EventType string `c:"EventType"`
	DataList  []struct {
		Address        string  `c:"address"`
		CheckTime      int     `c:"checkTime"`
		CorpId         string  `c:"corpId"`
		GroupId        string  `c:"groupId"`
		Latitude       float64 `c:"latitude"`
		BizId          string  `c:"bizId"`
		LocationMethod string  `c:"locationMethod"`
		UserId         string  `c:"userId"`
		Longitude      float64 `c:"longitude"`
	} `c:"DataList"`
}

/*
家校相关
*/
type EduDeptInsertEvent struct {
	EventType string `c:"EventType"`
	CorpId    string `c:"corpId"`
	DeptId    int    `c:"deptId"`
}
type EduDeptUpdateEvent struct {
	EventType string `c:"EventType"`
	CorpId    string `c:"corpId"`
	DeptId    int    `c:"deptId"`
}
type EduDeptDeleteEvent struct {
	EventType  string `c:"EventType"`
	CorpId     string `c:"corpId"`
	DeptId     int    `c:"deptId"`
	IsGraduate int    `c:"isGraduate"`
}
type EduUserInsertEvent struct {
	EventType string `c:"EventType"`
	ClassId   int    `c:"classId"`
	CorpId    string `c:"corpId"`
	Feature   string `c:"feature"`
	Name      string `c:"name"`
	Role      string `c:"role"`
	UserId    string `c:"userId"`
}
type EduUserUpdateEvent struct {
	EventType string `c:"EventType"`
	ClassId   int    `c:"classId"`
	CorpId    string `c:"corpId"`
	Feature   string `c:"feature"`
	Name      string `c:"name"`
	Role      string `c:"role"`
	UserId    string `c:"userId"`
}
type EduUserDeleteEvent struct {
	EventType  string `c:"EventType"`
	ClassId    int    `c:"classId"`
	CorpId     string `c:"corpId"`
	Feature    string `c:"feature"`
	Role       string `c:"role"`
	UserId     string `c:"userId"`
	IsGraduate int    `c:"isGraduate"`
}
type EduUserRelationInsertEvent struct {
	EventType    string `c:"EventType"`
	ClassId      int    `c:"classId"`
	CorpId       string `c:"corpId"`
	FromUserId   string `c:"fromUserId"`
	RelationCode string `c:"relationCode"`
	RelationName string `c:"relationName"`
	ToUserId     string `c:"toUserId"`
}
type EduUserRelationUpdateEvent struct {
	EventType    string `c:"EventType"`
	ClassId      int    `c:"classId"`
	CorpId       string `c:"corpId"`
	FromUserId   string `c:"fromUserId"`
	RelationCode string `c:"relationCode"`
	RelationName string `c:"relationName"`
	ToUserId     string `c:"toUserId"`
}
type EduUserRelationDeleteEvent struct {
	EventType    string `c:"EventType"`
	ClassId      int    `c:"classId"`
	CorpId       string `c:"corpId"`
	FromUserId   string `c:"fromUserId"`
	RelationCode string `c:"relationCode"`
	RelationName string `c:"relationName"`
	ToUserId     string `c:"toUserId"`
}
