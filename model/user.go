package model

type UserIdResp struct {
	Errcode  int    `json:"errcode"`
	Errmsg   string `json:"errmsg"`
	Userid   string `json:"userid"`
	IsSys    bool   `json:"is_sys"`
	SysLevel int    `json:"sys_level"`
}

type UserDetailResp struct {
	Errcode int       `json:"errcode"`
	Errmsg  string    `json:"errmsg"`
	Avatar  string    `json:"avatar"`
	Userid  string    `json:"userid"`
	Name    string    `json:"name"`
	Tags    *UserTags `json:"tags"`
}
type UserTags struct {
	Guardian []string `json:"guardian"`
	Student  []string `json:"student"`
}
