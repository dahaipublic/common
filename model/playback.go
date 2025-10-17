package model

// 回放
type Playback struct {
	BasicField
	SportsType   int8   `json:"sportsType" gorm:"column:sports_type;default:1;not null;comment:运动类型:1足球2篮球"`
	MatchID      uint64 `json:"matchID" gorm:"column:match_id;default:0;not null;comment:比赛id"`
	HomeTeamName string `json:"homeTeamName" gorm:"column:home_team_name;type:varchar(128);default:'';not null;comment:主队名称"`
	AwayTeamName string `json:"awayTeamName" gorm:"column:away_team_name;type:varchar(128);default:'';not null;comment:客队名称"`
	Status       int8   `json:"status" gorm:"column:status;default:1;not null;comment:录制状态:1进行中2成功3失败"`
}

func (t Playback) TableName() string {
	return "t_playback"
}
