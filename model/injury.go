package model

type TBaseInjury struct {
	BasicField
	Injury TJsonString `json:"injury" gorm:"column:injury;type:longtext;comment:球队伤停列表"`
}

type FootballInjury struct {
	TBaseInjury
}
type BasketballInjury struct {
	TBaseInjury
}

type InjuryInfo struct {
	PlayerId      uint64 `json:"player_id"`
	CptID         uint64 `json:"competition_id"`
	Type          int64  `json:"type"`           //球员logo
	Reason        string `json:"reason"`         //类型，1-受伤、2-停赛、0-未知
	StartTime     int64  `json:"start_time"`     // 开始时间
	EndTime       int64  `json:"end_time"`       // 结束时间
	MissedMatches int64  `json:"missed_matches"` // 投篮次数

}

func (t FootballInjury) TableName() string   { return "t_football_injury" }
func (t BasketballInjury) TableName() string { return "t_basketball_injury" }
