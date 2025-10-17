package model

// 热门赛事表
type XXCompetition struct {
	BasicField
	CompetitionID uint64 `json:"competitionID" gorm:"column:competition_id;default:0;not null;comment:赛事id" `
	SportsType    int8   `json:"sportsType" gorm:"column:sports_type;default:1;not null;comment:运动类型:1足球2篮球"`
	Sort          int64  `json:"sort" gorm:"column:sort;default:0;not null;comment:排序"`
}

type HotCompetition XXCompetition

func (t HotCompetition) TableName() string {
	return "t_hot_competition"
}

//type RecommendCompetition XXCompetition
//
//func (t RecommendCompetition) TableName() string {
//	return "t_recommend_competition"
//}
