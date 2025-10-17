package model

// 意见表
type Opinion struct {
	BasicField
	UserID   uint64 `json:"userID" gorm:"column:user_id;default:0;not null;comment:用户id"`
	Content  string `json:"content" gorm:"column:content;type:varchar(256);default:'';not null;comment:内容"`
	ImageUrl string `json:"imageUrl" gorm:"column:image_url;type:text;comment:图片"`
}

func (t Opinion) TableName() string {
	return "t_opinion"
}
