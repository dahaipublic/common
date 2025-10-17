package model

type GameChannel struct {
	BasicField
	ChannelSite string `json:"channelSite" gorm:"column:channel_site;type:varchar(256);default:'';not null;comment:渠道名称" `
	WebSite     string `json:"WebSite" gorm:"column:web_site;type:varchar(256);default:'';not null;comment:渠道网站" `
}

func (t GameChannel) TableName() string { return "t_game_channel" }
