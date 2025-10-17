package model

// 热门球队表
type HotTeam struct {
	BasicField
	TeamID     int64 `json:"teamID" gorm:"column:team_id;default:0;not null;comment:球队id" `
	SportsType int8  `json:"sportsType" gorm:"column:sports_type;default:1;not null;comment:运动类型:1足球2篮球"`
	Sort       int64 `json:"sort" gorm:"column:sort;default:0;not null;comment:排序"`
}

func (t HotTeam) TableName() string {
	return "t_hot_team"
}
