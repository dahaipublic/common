package model

// 广告表
type AppAdvertisement struct {
	BasicField
	AdvertisementType int8   `json:"advertisementType" gorm:"column:advertisement_type;default:1;not null;comment:类型:1启动图2首页banner3首页滚动广告"`
	ImageType         int8   `json:"imageType" gorm:"column:image_type;default:0;not null;comment:类型:1图片2视频3git"`
	ImageUrl          string `json:"imageUrl" gorm:"column:image_url;type:text;comment:图片"`
	GifUrl            string `json:"gifUrl" gorm:"column:gif_url;type:text;comment:gif地址"`
	PlayTime          int64  `json:"playTime" gorm:"column:play_time;default:0;not null;comment:播放时间"`
	JumpType          int8   `json:"jumpType" gorm:"column:jump_type;default:0;not null;comment:跳转类型:1比赛2外链"`
	Title             string `json:"title" gorm:"column:title;type:varchar(255);default:'';not null;comment:标题"`
	Content           string `json:"content" gorm:"column:content;type:text;comment:内容"`
	SportsType        int8   `json:"sportsType" gorm:"column:sports_type;default:0;not null;comment:类型:1足球2篮球"`
	MatchID           uint64 `json:"matchID" gorm:"column:match_id;default:0;not null;comment:比赛id"`
	Sort              int64  `json:"sort" gorm:"column:sort;default:0;not null;comment:排序号"`
	Url               string `json:"url" gorm:"column:url;type:text;comment:外链"`
	BackgroundImage   string `json:"backgroundImage" gorm:"column:background_image;type:text;comment:背景图片"`
	ColorValue        string `json:"colorValue" gorm:"column:color_value;type:varchar(16);default:'';not null;comment:色值"`
	OnOff             int8   `json:"onOff" gorm:"column:on_off;default:1;not null;comment:开关:0关1开"`
	//ChannelId         int8   `json:"channelId" gorm:"column:channel_id;type:varchar(256);default:'';not null;comment:渠道Id" `
	ChannelId int8 `json:"channelId" gorm:"column:channel_id;default:1;not null;comment:渠道Id"`
}

func (t AppAdvertisement) TableName() string {
	return "t_app_advertisement"
}
