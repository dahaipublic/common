package model

// 推荐比赛表
type RecommendMatch struct {
	BasicField
	MatchID           uint64 `json:"matchID" gorm:"column:match_id;default:0;not null;comment:比赛id" `
	SportsType        int8   `json:"sportsType" gorm:"column:sports_type;default:1;not null;comment:运动类型:1足球2篮球"`
	Sort              int64  `json:"sort" gorm:"column:sort;default:0;not null;comment:排序"`
	ImageUrl          string `json:"imageUrl" gorm:"column:image_url;type:varchar(255);default:'';not null;comment:图片地址"`
	QueryTime         int64  `json:"queryTime" gorm:"column:query_time;index:index_query_time;default:0;not null;comment:查询时间"`
	CompetitionNameZh string `json:"competitionNameZh" gorm:"column:competition_name_zh;type:varchar(255);default:'';not null;comment:联赛名称(中文)"`
}

type RecommendMatchAdmin struct {
	RecommendMatch
}

func (t RecommendMatch) TableName() string {
	return "t_recommend_match"
}
