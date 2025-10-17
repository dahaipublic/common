package model

type ChatCommunity struct {
	BasicField

	Type          int64 `json:"type" gorm:"column:type;default:0;not null;comment:类型(1.TiktoK。2.Whatsapp)"`
	CommunityLink int64 `json:"communityLink" gorm:"column:community_link;default:0;not null;comment:链接"`
}

func (t ChatCommunity) TableName() string {
	return "t_chat_community"
}
