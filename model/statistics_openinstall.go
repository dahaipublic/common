package model

// 统计表
type StatisticsOpeninstall struct {
	BasicField
	DeviceType int8   `json:"deviceType" gorm:"column:device_type;default:0;not null;comment:类型1ios2安卓"`
	Date       string `json:"date" gorm:"column:date;type:varchar(32);default:'';not null;comment:日期"`
	Visit      int64  `json:"visit" gorm:"column:visit;default:0;not null;comment:落地页的访问数"`
	Click      int64  `json:"click" gorm:"column:click;default:0;not null;comment:落地页的点击数"`
	Install    int64  `json:"install" gorm:"column:install;default:0;not null;comment:安装数"`
	Register   int64  `json:"register" gorm:"column:register;default:0;not null;comment:注册数"`
	SurviveD1  int64  `json:"survive_d1" gorm:"column:survive_d1;default:0;not null;comment:安装1天后的留存数"`
	SurviveD7  int64  `json:"survive_d7" gorm:"column:survive_d7;default:0;not null;comment:安装7天后的留存数"`
	SurviveD30 int64  `json:"survive_d30" gorm:"column:survive_d30;default:0;not null;comment:安装30天后的留存数"`
}

func (t StatisticsOpeninstall) TableName() string {
	return "t_statistics_openinstall"
}
