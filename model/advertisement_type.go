package model

type AdvertisementType struct {
	BasicField
	AdsType string `json:"adsType" gorm:"column:ads_type;type:varchar(256);default:'';not null;comment:广告类型"`
}

func (t AdvertisementType) TableName() string {
	return "t_advertisement_type"
}

type AdvertisementStatistics struct {
	BasicField
	AdsId      string `json:"adsId" gorm:"column:ads_id;type:varchar(256);default:'';not null;comment:广告位置"`
	DeviceType int8   `json:"deviceType" gorm:"column:device_type;default:0;not null;comment:类型1ios2安卓3pc4h5"`
	DeviceID   string `json:"deviceId" gorm:"column:device_id;type:varchar(64);default:'';not null;comment:设备号"`
	ClientIp   string `json:"clientIp" gorm:"column:client_ip;type:varchar(64);default:'';not null;comment:用户ip"`
	Source     string `json:"source" gorm:"column:source;type:varchar(64);default:'';not null;comment:来源网址"`
}

func (t AdvertisementStatistics) TableName() string {
	return "t_advertisement_statistics"
}

type AdvertisementClick struct {
	BasicField
	AdsId      string `json:"adsId" gorm:"column:ads_id;type:varchar(256);default:'';not null;comment:广告位置"`
	DeviceType int8   `json:"deviceType" gorm:"column:device_type;default:0;not null;comment:类型1ios2安卓3pc4h5"`
	DeviceID   string `json:"deviceId" gorm:"column:device_id;type:varchar(64);default:'';not null;comment:设备号"`
	ClientIp   string `json:"clientIp" gorm:"column:client_ip;type:varchar(64);default:'';not null;comment:用户ip"`
	Source     string `json:"source" gorm:"column:source;type:varchar(64);default:'';not null;comment:来源网址"`
}

func (t AdvertisementClick) TableName() string {
	return "t_advertisement_click"
}
