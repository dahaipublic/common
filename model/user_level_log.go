package model

// 用户等级记录表
type UserLevelLog struct {
	BasicField
	UserID              uint64 `json:"userID" gorm:"column:user_id;default:0;not null;comment:用户id"`
	UserExperienceLogID uint64 `json:"userExperienceLogID" gorm:"column:user_experience_log_id;default:0;not null;comment:来自用户id"`
	BeforeLevel         int64  `json:"beforeLevel" gorm:"column:before_level;default:0;not null;comment:变动前"`
	LastLevel           int64  `json:"lastLevel" gorm:"column:last_level;default:0;not null;comment:变动后"`
}

func (t UserLevelLog) TableName() string {
	return "t_user_level_log"
}
