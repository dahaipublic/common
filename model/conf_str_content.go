package model

// 字符串内容表; 用于存字符串内容,目前分享功能在用
type ConfStrContent struct {
	BasicField
	Title      string `json:"title" gorm:"column:title;type:varchar(255);default:'';not null;comment:标题"`
	Content    string `json:"content" gorm:"column:content;type:text;comment:内容1"`
	ContentTwo string `json:"content_two" gorm:"column_two:content;type:text;comment:内容2"`
	Position   string `json:"position" gorm:"column:position;default:'';not null;comment:位置"`
}

func (t ConfStrContent) TableName() string {
	return "t_str_content"
}
