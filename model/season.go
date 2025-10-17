package model

type TBaseSeason struct {
	BasicField

	CompetitionId int64  `json:"competitionId" gorm:"column:competition_id;default:0;not null;comment:赛事id" nami:"competition_id"`
	Year          string `json:"year" gorm:"column:year;type:varchar(256);default:'';not null;comment:赛季年份" nami:"year"`

	IsCurrent      int64 `json:"isCurrent" gorm:"column:is_current;default:0;not null;comment:是否最新赛季，1-是、0-否" nami:"is_current"`
	HasPlayerStats int64 `json:"hasPlayerStats" gorm:"column:has_player_stats;default:0;not null;comment:是否有球员统计，1-是、0-否" nami:"has_player_stats"`
	HasTeamStats   int64 `json:"hasTeamStats" gorm:"column:has_team_stats;default:0;not null;comment:是否有球队统计，1-是、0-否" nami:"has_team_stats"`
	HasTable       int64 `json:"HasTable" gorm:"column:has_table;default:0;not null;comment:是否有积分榜，1-是、0-否" nami:"has_table"`
}

type FootballSeason struct {
	StartTime int64 `json:"startTime" gorm:"column:start_time;default:0;not null;comment:开始时间" nami:"start_time"`
	EndTime   int64 `json:"endTime" gorm:"column:end_time;default:0;not null;comment:结束时间" nami:"end_time"`
	TBaseSeason
}
type BasketballSeason struct {
	TBaseSeason
}

func (t FootballSeason) TableName() string   { return "t_football_season" }
func (t BasketballSeason) TableName() string { return "t_basketball_season" }
