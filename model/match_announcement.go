package model

// 比赛公告表
type MatchAnnouncement struct {
	BasicField
	UserID   uint64 `json:"userID" gorm:"column:user_id;default:0;not null;comment:用户id"`
	Name     string `json:"name" gorm:"column:name;type:varchar(32);default:'';not null;comment:用户名"`
	RoleID   int8   `json:"roleID" gorm:"column:role_id;default:0;not null;comment:角色id"`
	MatchID  uint64 `json:"matchID" gorm:"column:match_id;default:0;not null;comment:比赛id"`
	ImageUrl string `json:"imageUrl" gorm:"column:image_url;type:text;not null;comment:图片"`
	Content  string `json:"content" gorm:"column:content;type:text;not null;comment:内容"`
}

func (t MatchAnnouncement) TableName() string {
	return "t_match_announcement"
}
