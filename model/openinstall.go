package model

// 数据渠道
type Openinstall struct {
	BasicField
	ChannelCode string `json:"channelCode" gorm:"column:channel_code;type:varchar(16);default:'';not null;comment:渠道号"`
	Url         string `json:"url" gorm:"column:url;type:varchar(256);default:'';not null;comment:渠道链接"`
	Title       string `json:"title" gorm:"column:title;type:varchar(32);default:'';not null;comment:标题"`
}

func (t Openinstall) TableName() string {
	return "t_openinstall"
}
