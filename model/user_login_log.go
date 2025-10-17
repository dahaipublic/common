package model

// 用户登录记录表
type UserLoginLog struct {
	BasicField
	UserID uint64 `json:"userID" gorm:"column:user_id;default:0;not null;comment:用户id"`
	IP     string `json:"ip" gorm:"column:ip;type:varchar(64);default:'';not null;comment:ip"`
}

func (t UserLoginLog) TableName() string {
	return "t_user_login_log"
}
