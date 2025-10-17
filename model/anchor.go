package model

// 主播表
type Anchor struct {
	BasicField
	RoomID         uint64 `json:"roomID" gorm:"column:room_id;default:0;not null;comment:直播间id"`
	AnchorType     int8   `json:"anchorType" gorm:"column:anchor_type;default:0;not null;comment:类型:1系统主播2真人主播3优质主播"`
	UserID         uint64 `json:"userID" gorm:"column:user_id;default:0;not null;comment:用户id"`
	FansCount      uint64 `json:"fansCount" gorm:"column:fans_count;default:0;not null;comment:粉丝数量"`
	BasisFansCount uint64 `json:"basisFansCount" gorm:"column:basis_fans_count;default:0;not null;comment:基础粉丝数量"`
	IsLive         int8   `json:"isLive" gorm:"column:is_live;default:1;not null;comment:是否直播1否2是"`
	ImageUrl       string `json:"imageUrl" gorm:"column:image_url;type:varchar(255);default:'';not null;comment:图片地址"`
}

func (t Anchor) TableName() string {
	return "t_anchor"
}
