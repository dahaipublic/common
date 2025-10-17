package model

type ChatMessage struct {
	BasicField
	MatchID   uint64 `json:"matchID" gorm:"column:match_id;index:match_id;default:0;not null;comment:比赛id"`
	CurrentId string `json:"currentId" gorm:"column:current_id;type:varchar(255);default:'';not null;comment:当前下标"`
	Message   string `json:"message" gorm:"column:message;type:longtext;comment:消息"`
}
