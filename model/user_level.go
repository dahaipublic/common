package model

// 用户等级表
type UserLevel struct {
	BasicField
	Level      int64 `json:"level" gorm:"column:level;default:0;not null;comment:等级"`
	Experience int64 `json:"experience" gorm:"column:experience;default:0;not null;comment:经验值"`
}

func (t UserLevel) TableName() string {
	return "t_user_level"
}
