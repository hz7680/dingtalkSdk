package client

import (
	"encoding/json"
	"github.com/hz7680/dingtalkSdk/model"
	"github.com/hz7680/dingtalkSdk/utils"
	"log"
	"strings"
)

const (
	MSG_TYPE_CARD = "action_card"
)

func (c *Client) SendCorpConversation(userIdList []string, deptIdList []string, toAllUser bool, msgType string, msg interface{}) (resp *model.CorpConversationMessageResp, err error) {
	token, err := c.GetAccessToken()
	if err != nil {
		log.Println("get access token error : " + err.Error())
		return
	}
	url := "https://oapi.dingtalk.com/topapi/message/corpconversation/asyncsend_v2?access_token=" + token
	//msgData,_ := json.Marshal(generateMsgData(msgType,msg))
	data := map[string]interface{}{
		"agent_id":    c.agentId,
		"to_all_user": toAllUser,
		"msg":         c.generateMsgData(msgType, msg),
	}
	if userIdList != nil && len(userIdList) > 0 {
		data["userid_list"] = strings.Join(userIdList, ",")
	}
	if deptIdList != nil && len(userIdList) > 0 {
		data["dept_id_list"] = strings.Join(deptIdList, ",")
	}
	body, err := utils.HttpPostJson(url, data)
	resp = &model.CorpConversationMessageResp{}
	err = json.Unmarshal(body, resp)
	return
}

func (c *Client) generateMsgData(msgType string, msg interface{}) map[string]interface{} {
	data := map[string]interface{}{
		"msgtype": msgType,
	}
	switch msgType {
	case MSG_TYPE_CARD:
		data["action_card"] = msg
	}
	return data
}
