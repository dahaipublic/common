package model

// 表情包表
type Emoticons struct {
	BasicField
	ImageUrl      string `json:"imageUrl" gorm:"column:image_url;type:text;not null;comment:图片"`
	Sort          int64  `json:"sort" gorm:"column:sort;default:0;not null;comment:排序"`
	EmoticonsType int8   `json:"emoticonsType" gorm:"column:emoticons_type;default:1;not null;comment:类型:1基础表情2gif表情"`
	EmoticonsCode string `json:"emoticonsCode" gorm:"column:emoticons_code;type:varchar(32);default:'';not null;comment:表情代号"`
}

func (t Emoticons) TableName() string {
	return "t_emoticons"
}
