package model

// 统计表
type StatisticsActiveUser struct {
	BasicField
	DeviceType int8  `json:"deviceType" gorm:"column:device_type;default:0;not null;comment:类型1ios2安卓3pc4h5"`
	Num        int64 `json:"num" gorm:"column:num;default:0;not null;comment:数量"`
	Date       int64 `json:"date" gorm:"column:date;default:0;not null;comment:时间"`
}

func (t StatisticsActiveUser) TableName() string {
	return "t_statistics_active_user"
}
