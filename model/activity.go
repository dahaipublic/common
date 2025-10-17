package model

// 活动表
type Activity struct {
	BasicField
	ImageUrl         string `json:"imageUrl" gorm:"column:image_url;type:varchar(256);default:'';not null;comment:图片"`
	Title            string `json:"title" gorm:"column:title;type:varchar(64);default:'';not null;comment:标题"`
	Detail           string `json:"detail" gorm:"column:detail;type:varchar(256);default:'';not null;comment:详情"`
	Content          string `json:"content" gorm:"column:content;type:text;comment:更新内容"`
	ActivityCategory uint64 `json:"activityCategory" gorm:"column:activity_category;default:0;not null;comment:类型:1福利2公告3活动"`
	Sort             int64  `json:"sort" gorm:"column:sort;default:0;not null;comment:排序"`
	Status           int8   `json:"status" gorm:"column:status;default:1;not null;comment:状态:1未开始2进行中3已结束"`
}

func (t Activity) TableName() string {
	return "t_activity"
}
