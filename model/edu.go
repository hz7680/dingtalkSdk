package model

type EduDeptListResp struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	Success bool   `json:"success"`
	Result  struct {
		Details []*EduDeptDetail `json:"details"`
		HasMore bool             `json:"has_more"`
		SuperId int              `json:"super_id"`
	} `json:"result"`
}

type EduDeptResp struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	Success bool   `json:"success"`
	Result  struct {
		Detail *EduDeptDetail `json:"detail"`
	} `json:"result"`
}

type EduDeptDetail struct {
	DeptId      int    `json:"dept_id"`
	Nick        string `json:"nick"`
	Chain       string `json:"chain"`
	Feature     string `json:"feature"`
	Name        string `json:"name"`
	ContactType string `json:"contact_type"`
	DeptType    string `json:"dept_type"`
}

type EduUserListResp struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	Success bool   `json:"success"`
	Result  struct {
		Details []*Detail `json:"details"`
		HasMore bool      `json:"has_more"`
	} `json:"result"`
}

type EduUserResp struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	Success bool   `json:"success"`
	Result  struct {
		Details []*Detail `json:"details"`
	} `json:"result"`
}

type Detail struct {
	Userid  string `json:"userid"`
	ClassId int    `json:"class_id"`
	Role    string `json:"role"`
	Feature string `json:"feature"`
	Name    string `json:"name"`
}

type EduUserRelationListResp struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	Success bool   `json:"success"`
	Result  struct {
		Relations []*Relation `json:"relations"`
		HasMore   bool        `json:"has_more"`
	} `json:"result"`
}

type EduUserRelationResp struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	Success bool   `json:"success"`
	Result  struct {
		Relations []*Relation `json:"relations"`
	} `json:"result"`
}

type Relation struct {
	ClassId      int    `json:"class_id"`
	FromUserid   string `json:"from_userid"`
	RelationCode string `json:"relation_code"`
	RelationName string `json:"relation_name"`
	ToUserid     string `json:"to_userid"`
}
