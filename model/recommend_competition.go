package model

type RecommendCompetition struct {
	BasicField
	CompetitionID uint64 `json:"competitionID" gorm:"column:competition_id;default:0;not null;comment:赛事id" `
	SportsType    int8   `json:"sportsType" gorm:"column:sports_type;default:1;not null;comment:运动类型:1足球2篮球"`
	Sort          int64  `json:"sort" gorm:"column:sort;default:0;not null;comment:排序"`
}

type RecommendCompetitionAdmin struct {
	RecommendCompetition
	CompetitionNameZh string `json:"competitionNameZh" gorm:"column:competition_name_zh;type:varchar(255);default:'';not null;comment:联赛名称(中文)"`
}

func (t RecommendCompetition) TableName() string {
	return "t_recommend_competition"
}
