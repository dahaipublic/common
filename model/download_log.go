package model

// 下载记录表
type DownloadLog struct {
	BasicField
	DeviceType int8   `json:"deviceType" gorm:"column:device_type;default:1;not null;comment:设备类型：1ios2Android"`
	IP         string `json:"ip" gorm:"column:ip;type:varchar(64);default:'';not null;comment:ip"`
}

func (t DownloadLog) TableName() string {
	return "t_download_log"
}
