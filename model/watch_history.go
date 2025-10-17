package model

// 观看历史
type WatchHistory struct {
	BasicField
	MatchID    uint64 `json:"matchID" gorm:"column:match_id;default:0;not null;comment:比赛id" `
	SportsType int8   `json:"sportsType" gorm:"column:sports_type;default:1;not null;comment:运动类型:1足球2篮球"`
	UserID     uint64 `json:"userID" gorm:"column:user_id;default:0;not null;comment:用户id" `
}

func (t WatchHistory) TableName() string {
	return "t_watch_history"
}
