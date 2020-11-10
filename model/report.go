package model

type ReportSimplelist struct {
	Errcode int                     `json:"errcode"`
	Errmsg  string                  `json:"errmsg"`
	Result  *ReportSimplelistResult `json:"result"`
}

type ReportSimplelistResult struct {
	DataList   []*ReportSimpleData `json:"data_list"`
	Size       int                 `json:"size"`
	NextCursor int                 `json:"next_cursor"`
	HasMore    bool                `json:"has_more"`
}

type ReportSimpleData struct {
	Remark       string `json:"remark"`
	TemplateName string `json:"template_name"`
	DeptName     string `json:"dept_name"`
	CreatorName  string `json:"creator_name"`
	CreatorId    string `json:"creator_id"`
	CreateTime   int    `json:"create_time"`
	ReportId     string `json:"report_id"`
}

type ReportList struct {
	Errcode int               `json:"errcode"`
	Errmsg  string            `json:"errmsg"`
	Result  *ReportListResult `json:"result"`
}

type ReportListResult struct {
	DataList   []*ReportData `json:"data_list"`
	Size       int           `json:"size"`
	NextCursor int           `json:"next_cursor"`
	HasMore    bool          `json:"has_more"`
}

type ReportData struct {
	Contents     []*ReportDataContent `json:"contents"`
	Remark       string               `json:"remark"`
	TemplateName string               `json:"template_name"`
	DeptName     string               `json:"dept_name"`
	CreatorName  string               `json:"creator_name"`
	CreatorId    string               `json:"creator_id"`
	CreateTime   int                  `json:"create_time"`
	ReportId     string               `json:"report_id"`
	Images       []string             `json:"images"`
}
type ReportDataContent struct {
	Sort  string `json:"sort"`
	Type  string `json:"type"`
	Value string `json:"value"`
	Key   string `json:"key"`
}
