package model

// 主播申请表
type AnchorApply struct {
	BasicField
	UserID  uint64 `json:"userID" gorm:"column:user_id;default:0;not null;comment:用户id"`
	Content string `json:"content" gorm:"column:content;type:varchar(256);default:'';not null;comment:内容"`
}

func (t AnchorApply) TableName() string {
	return "t_anchor_apply"
}
