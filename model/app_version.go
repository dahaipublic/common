package model

// app版本
type AppVersion struct {
	BasicField
	VersionCode uint64 `json:"versionCode" gorm:"column:version_code;index:version_code;default:0;not null;comment:版本code"`
	VersionName string `json:"versionName" gorm:"column:version_name;type:varchar(32);default:'';not null;comment:版本号"`
	AppUrl      string `json:"appUrl" gorm:"column:app_url;type:varchar(256);default:'';not null;comment:app链接"`
	Content     string `json:"content" gorm:"column:content;type:text;comment:更新内容"`
	ForceUpdate int8   `json:"forceUpdate" gorm:"column:force_update;default:0;not null;comment:是否强制更新0否1是"`
	DeviceType  int8   `json:"deviceType" gorm:"column:device_type;default:0;not null;comment:设备类型:1android2ios"`
}

func (t AppVersion) TableName() string {
	return "t_app_version"
}
