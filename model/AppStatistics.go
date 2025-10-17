package model

type AppStatistics struct {
	BasicField
	IMEI      string `json:"imei" gorm:"column:imei;type:varchar(255);default:'';not null;comment:imei"`
	ClientIp  string `json:"clientIp" gorm:"column:clientIp;type:varchar(255);default:'';not null;comment:clientIp"`
	Timestamp int64  `json:"timestamp" gorm:"column:timestamp;default:0;not null;comment:时间戳"`
	Pkg       string `json:"pkg" gorm:"column:pkg;type:varchar(255);default:'';not null;comment:包名"`
	From      string `json:"from" gorm:"column:from;type:varchar(255);default:'';not null;comment:来源(android/ios)"`
	Type      int64  `json:"type" gorm:"column:type;default:0;not null;comment:类型(1、安装 2、注册 3、登录)"`
}

func (t AppStatistics) TableName() string {
	return "t_app_statistics"
}
