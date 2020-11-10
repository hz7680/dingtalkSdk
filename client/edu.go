package client

import (
	"encoding/json"
	"fmt"
	"github.com/hz7680/dingtalkSdk/model"
	"github.com/hz7680/dingtalkSdk/utils"
	"log"
)

const (
	ROLE_TEACHER  = "teacher"
	ROLE_STUDENT  = "student"
	ROLE_GUARDIAN = "guardian"

	EDU_DEPT_TYPE_CLASS = "class"
)

func (c *Client) GetEduDeptList(pageNo, pageSize, superId int) (resp *model.EduDeptListResp, err error) {
	token, err := c.GetAccessToken()
	if err != nil {
		log.Println("get access token error : " + err.Error())
		return
	}
	url := "https://oapi.dingtalk.com/topapi/edu/dept/list?access_token=" + token
	data := map[string]interface{}{
		"page_no":   pageNo,
		"page_size": pageSize,
	}
	if superId != 0 {
		data["super_id"] = superId
	}
	body, err := utils.HttpPostJson(url, data)
	if err != nil {
		return
	}
	resp = &model.EduDeptListResp{}
	err = json.Unmarshal(body, resp)
	return
}

func (c *Client) GetEduDept(deptId int) (resp *model.EduDeptResp, err error) {
	token, err := c.GetAccessToken()
	if err != nil {
		log.Println("get access token error : " + err.Error())
		return
	}
	url := "https://oapi.dingtalk.com/topapi/edu/dept/get?access_token=" + token
	body, err := utils.HttpPostJson(url, map[string]interface{}{
		"dept_id": deptId,
	})
	if err != nil {
		return
	}
	resp = &model.EduDeptResp{}
	err = json.Unmarshal(body, resp)
	return
}

func (c *Client) GetEduUserList(classId, pageNo, pageSize int, role string) (resp *model.EduUserListResp, err error) {
	token, err := c.GetAccessToken()
	if err != nil {
		log.Println("get access token error : " + err.Error())
		return
	}
	url := "https://oapi.dingtalk.com/topapi/edu/user/list?access_token=" + token
	body, err := utils.HttpPostJson(url, map[string]interface{}{
		"class_id":  classId,
		"role":      role,
		"page_no":   pageNo,
		"page_size": pageSize,
	})
	resp = &model.EduUserListResp{}
	err = json.Unmarshal(body, resp)
	return
}

func (c *Client) GetEduUser(classid int, userid, role string) (user *model.EduUserResp, err error) {
	token, err := c.GetAccessToken()
	if err != nil {
		log.Println("get access token error : " + err.Error())
		return
	}
	url := "https://oapi.dingtalk.com/topapi/edu/user/get?access_token=" + token
	body, err := utils.HttpPostJson(url, map[string]interface{}{
		"class_id": classid,
		"userid":   userid,
		"role":     role,
	})
	if err != nil {
		return
	}
	fmt.Println(string(body))
	user = &model.EduUserResp{}
	err = json.Unmarshal(body, user)
	return
}

func (c *Client) GetEduUserRelationList(classId, pageNo, pageSize int) (resp *model.EduUserRelationListResp, err error) {
	token, err := c.GetAccessToken()
	if err != nil {
		log.Println("get access token error : " + err.Error())
		return
	}
	url := "https://oapi.dingtalk.com/topapi/edu/user/relation/list?access_token=" + token
	body, err := utils.HttpPostJson(url, map[string]interface{}{
		"class_id":  classId,
		"page_no":   pageNo,
		"page_size": pageSize,
	})
	if err != nil {
		return
	}
	resp = &model.EduUserRelationListResp{}
	err = json.Unmarshal(body, resp)
	return
}

func (c *Client) GetEduUserRelation(classId int, fromUserId string) (relation *model.EduUserRelationResp, err error) {
	token, err := c.GetAccessToken()
	if err != nil {
		log.Println("get access token error : " + err.Error())
		return
	}
	url := "https://oapi.dingtalk.com/topapi/edu/user/relation/get?access_token=" + token
	body, err := utils.HttpPostJson(url, map[string]interface{}{
		"class_id":    classId,
		"from_userid": fromUserId,
	})
	if err != nil {
		return
	}
	relation = &model.EduUserRelationResp{}
	err = json.Unmarshal(body, relation)
	return
}
