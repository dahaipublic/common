package model

// 广告表
type Advertisement struct {
	BasicField
	Content           string `json:"content" gorm:"column:content;type:varchar(256);default:'';not null;comment:内容"`
	Image             string `json:"image" gorm:"column:image;type:text;not null;comment:图片"`
	StartAt           int64  `json:"startAt" gorm:"column:start_at;default:0;not null;comment:开始时间"`
	EndAt             int64  `json:"endAt" gorm:"column:end_at;default:0;not null;comment:结束时间"`
	SportsType        int8   `json:"sportsType" gorm:"column:sports_type;default:1;not null;comment:类型:1足球2篮球"`
	AdvertisementType int8   `json:"advertisementType" gorm:"column:advertisement_type;default:1;not null;comment:类型:1文字2宣传图3二维码"`
	MidfieldDisplay   int8   `json:"midfieldDisplay" gorm:"column:midfield_display;default:0;not null;comment:中场显示:0否1是"`
	Position          int8   `json:"position" gorm:"column:position;default:0;not null;comment:位置:1上2下"`
	Status            int8   `json:"status" gorm:"column:status;default:0;not null;comment:显示:0否1是"`
	EventName         string `json:"eventName" gorm:"column:event_name;type:text;not null;comment:图片"`
	//ChannelId         int8   `json:"channelId" gorm:"column:channel_id;type:varchar(256);default:'';not null;comment:渠道Id" `
	ChannelId int8 `json:"channelId" gorm:"column:channel_id;default:0;not null;comment:渠道Id"`
}

func (t Advertisement) TableName() string {
	return "t_advertisement"
}

// 广告比赛表
type AdvertisementMatch struct {
	BasicField
	AdvertisementID   uint64 `json:"advertisementID" gorm:"column:advertisement_id;default:0;not null;comment:广告id"`
	SportsType        int8   `json:"sportsType" gorm:"column:sports_type;default:0;not null;comment:1足球2篮球"`
	MatchID           uint64 `json:"matchID" gorm:"column:match_id;default:0;not null;comment:比赛id"`
	StartAt           int64  `json:"startAt" gorm:"column:start_at;default:0;not null;comment:开始时间"`
	EndAt             int64  `json:"endAt" gorm:"column:end_at;default:0;not null;comment:结束时间"`
	AdvertisementType int8   `json:"advertisementType" gorm:"column:advertisement_type;default:1;not null;comment:类型:1文字2宣传图3二维码"`
	Status            int8   `json:"status" gorm:"column:status;default:0;not null;comment:显示:0否1是"`
}

func (t AdvertisementMatch) TableName() string {
	return "t_advertisement_match"
}
