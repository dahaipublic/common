package model

type MatchVideo struct {
	BasicField
	MatchID    uint64 `json:"matchID" gorm:"column:match_id;default:0; null;comment:比赛id"`
	SportsType int8   `json:"sportsType" gorm:"column:sports_type;default:1; null;comment:运动类型:1足球2篮球"`
	PushUrl1   string `json:"pushUrl1" gorm:"column:push_url1;default:''; null;comment:<UNK>url"`
	PushUrl3   string `json:"pushUrl3" gorm:"column:push_url3;default:''; null;comment:<UNK>url"`
}

func (t MatchVideo) TableName() string {
	return "t_match_video"
}
