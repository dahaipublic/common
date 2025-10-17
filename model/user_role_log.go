package model

// 用户等级记录表
type UserRoleLog struct {
	BasicField
	UserID         uint64 `json:"userID" gorm:"column:user_id;default:0;not null;comment:用户id"`
	UserLevelLogID uint64 `json:"userLevelLogID" gorm:"column:user_level_log_id;default:0;not null;comment:管理id"`
	BeforeRoleID   uint64 `json:"beforeRoleID" gorm:"column:before_role_id;default:0;not null;comment:变动前"`
	LastRoleID     uint64 `json:"lastRoleID" gorm:"column:last_role_id;default:0;not null;comment:变动后"`
}

func (t UserRoleLog) TableName() string {
	return "t_user_role_log"
}
