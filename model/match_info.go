package model

type MatchInfo struct {
	BasicField
	SportsType int8   `json:"sportsType" gorm:"column:sports_type;default:1;not null;comment:类型:1足球2篮球"`
	MatchID    uint64 `json:"matchID" gorm:"column:match_id;index:match_id;default:0;not null;comment:比赛id"`
	Message    string `json:"message" gorm:"column:message;type:longtext;comment:消息"`
	Live       string `json:"live" gorm:"column:live;type:longtext;comment:直播"`
	//去掉指数
	//Odds       string `json:"odds" gorm:"column:odds;type:longtext;comment:指数"`
	Analysis string `json:"analysis" gorm:"column:analysis;type:longtext;comment:分析"`
	Lineup   string `json:"lineup" gorm:"column:lineup;type:longtext;comment:阵容"`
}

func (t MatchInfo) TableName() string {
	return "t_match_info"
}
