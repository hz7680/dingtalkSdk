package model

type CardMessageSingle struct {
	Title       string `json:"title"`
	Markdown    string `json:"markdown"`
	SingleTitle string `json:"single_title"`
	SingleUrl   string `json:"single_url"`
}

type CardMessageList struct {
	Title          string            `json:"title"`
	Markdown       string            `json:"markdown"`
	BtnOrientation string            `json:"btn_orientation"`
	BtnJsonList    []*CardMessageBtn `json:"btn_json_list"`
}

type CardMessageBtn struct {
	Title     string `json:"title"`
	ActionUrl string `json:"action_url"`
}

type CorpConversationMessageResp struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	TaskId  int64  `json:"task_id"`
}
