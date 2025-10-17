package model

// 活动类别表
type ActivityCategory struct {
	BasicField
	Name     string `json:"name" gorm:"column:name;type:varchar(32);default:'';not null;comment:名称"`
	ImageUrl string `json:"imageUrl" gorm:"column:image_url;type:varchar(256);default:'';not null;comment:图片"`
}

func (t ActivityCategory) TableName() string {
	return "t_activity_category"
}
