package model

type DepartmentListResp struct {
	Errcode    int                    `json:"errcode"`
	Errmsg     string                 `json:"errmsg"`
	Department []*DepartmentThumbnail `json:"department"`
}
type DepartmentThumbnail struct {
	Id              int    `json:"id"`
	Name            string `json:"name"`
	Parentid        int    `json:"parentid"`
	CreateDeptGroup bool   `json:"createDeptGroup"`
	AutoAddUser     bool   `json:"autoAddUser"`
	Ext             string `json:"ext"`
}
