package model

type DailyStatistics struct {
	BasicField

	Date int64 `json:"date" gorm:"column:date;default:0;not null;comment:时间"`

	DailyAddAndroid int64 `json:"dailyAddAndroid" gorm:"column:daily_add_android;default:0;not null;comment:每日新增用户统计Android"`
	DailyAddIos     int64 `json:"dailyAddIos" gorm:"column:daily_add_ios;default:0;not null;comment:每日新增用户统计Ios"`
	DailyAddPC      int64 `json:"dailyAddPC" gorm:"column:daily_add_pc;default:0;not null;comment:每日新增用户统计PC"`
	DailyAddH5      int64 `json:"dailyAddH5" gorm:"column:daily_add_h5;default:0;not null;comment:每日新增用户统计H5"`

	DailyActiveAndroid int64 `json:"dailyActiveAndroid" gorm:"column:daily_active_android;default:0;not null;comment:每日活跃用户统计Android"`
	DailyActiveIos     int64 `json:"dailyActiveIos" gorm:"column:daily_active_ios;default:0;not null;comment:每日活跃用户统计Ios"`
	DailyActivePC      int64 `json:"dailyActivePC" gorm:"column:daily_active_pc;default:0;not null;comment:每日活跃用户统计PC"`
	DailyActiveH5      int64 `json:"dailyActiveH5" gorm:"column:daily_active_h5;default:0;not null;comment:每日活跃用户统计H5"`

	DailyTotalUserAndroid int64 `json:"dailyTotalUserAndroid" gorm:"column:daily_total_user_android;default:0;not null;comment:累计用户统计Android"`
	DailyTotalUserIos     int64 `json:"dailTotalUserIos" gorm:"column:daily_total_user_ios;default:0;not null;comment:累计用户统计Ios"`
	DailyTotalUserPC      int64 `json:"dailTotalUserPc" gorm:"column:daily_total_user_pc;default:0;not null;comment:累计用户统计PC"`
	DailyTotalUserH5      int64 `json:"dailTotalUserH5" gorm:"column:daily_total_user_h5;default:0;not null;comment:累计用户统计H5"`
}

func (t DailyStatistics) TableName() string {
	return "t_daily_statistics"
}
