package model

// import (
// 	"gorm.io/gorm"
// )

type TBaseTeam struct {
	BasicField
	CompetitionID int64  `json:"competitionID" gorm:"column:competition_id;default:0;not null;comment:赛事id" nami:"competition_id"`
	CountryID     int64  `json:"countryID" gorm:"column:country_id;default:0;not null;comment:国家id" nami:"country_id"`
	CoachID       uint64 `json:"coachID" gorm:"column:coach_id;default:0;not null;comment:教练id" nami:"coach_id"`
	NameZh        string `json:"namezh" gorm:"column:name_zh;type:varchar(128);default:'';not null;comment:中文名称" nami:"name_zh"`
	NameZht       string `json:"nameZht" gorm:"column:name_zht;type:varchar(128);default:'';not null;comment:粤语名称" nami:"name_zht"`
	NameEn        string `json:"nameEn" gorm:"column:name_en;type:varchar(128);default:'';not null;comment:英文名称" nami:"name_en"`

	ShortNameZh   string `json:"shortNameZh" gorm:"column:short_name_zh;type:varchar(128);default:'';not null;comment:中文简称" nami:"short_name_zh"`
	ShortNameZht  string `json:"shortNameZht" gorm:"column:short_name_zht;type:varchar(128);default:'';not null;comment:粤语简称" nami:"short_name_zht"`
	ShortNameEn   string `json:"shortNameEn" gorm:"column:short_name_en;type:varchar(128);default:'';not null;comment:英文简称" nami:"short_name_en"`
	Logo          string `json:"logo" gorm:"column:logo;type:varchar(256);default:'';not null;comment:logo" nami:"logo"`
	National      int8   `json:"national" gorm:"column:national;default:0;not null;comment:是否国家队，1-是、0-否" nami:"national"`
	VenueID       int64  `json:"venueID" gorm:"column:venue_id;default:0;not null;comment:场馆id" nami:"venue_id"`
	NamiUpdatedAt int64  `json:"namiUpdatedAt" gorm:"column:nami_updated_at;default:0;not null;comment:纳米更新时间" nami:"updated_at"`

	Sort int64 `json:"sort" gorm:"column:sort;default:0;not null;comment:排序"` // 注意这里不要用nami tag
}

type Team struct {
	TBaseTeam
	NameTr string `json:"nameTr" gorm:"column:name_tr;type:varchar(128);default:'';not null;comment:土耳其名称" nami:"name_tr"`
	NameAr string `json:"nameAr" gorm:"column:name_ar;type:varchar(128);default:'';not null;comment:埃及名称" nami:"name_ar"`
}

// 篮球球队同步实体
type NamiSyncBasketballTeam struct {
	TBaseTeam
	ConferenceID int64 `json:"conferenceID" gorm:"column:conference_id;default:0;not null;comment:赛区id，1-大西洋赛区、2-中部赛区、3-东南赛区、4-太平洋赛区、5-西北赛区、6-西南赛区、7-A组赛区、8-B组赛区、9-C组赛区、10-D组赛区（1~6:NBA 7~10:CBA）、0-无" nami:"conference_id"`
}

func (t NamiSyncBasketballTeam) TableName() string { return "t_basketball_team" }

// 篮球球队数据库实体
type BasketballTeam struct {
	Team
	ConferenceID int64 `json:"conferenceID" gorm:"column:conference_id;default:0;not null;comment:赛区id，1-大西洋赛区、2-中部赛区、3-东南赛区、4-太平洋赛区、5-西北赛区、6-西南赛区、7-A组赛区、8-B组赛区、9-C组赛区、10-D组赛区（1~6:NBA 7~10:CBA）、0-无" nami:"conference_id"`
}

func (t BasketballTeam) TableName() string { return "t_basketball_team" }

// --------------------

type FootballTeamExtend struct {
	SeasonID            int64  `json:"seasonID" gorm:"column:season_id;default:0;not null;comment:赛季id" nami:"season_id"`
	CountryLogo         string `json:"countryLogo" gorm:"column:country_logo;type:varchar(256);default:'';not null;comment:国家队logo（为国家队时存在）" nami:"country_logo"`
	FoundationTime      int64  `json:"foundationTime" gorm:"column:foundation_time;default:0;not null;comment:成立时间" nami:"foundation_time"`
	Website             string `json:"website" gorm:"column:website;type:varchar(256);default:'';not null;comment:球队官网" nami:"website"`
	MarketValue         int64  `json:"marketValue" gorm:"column:market_value;default:0;not null;comment:市值" nami:"market_value"`
	MarketValueCurrency string `json:"marketValueCurrency" gorm:"column:market_value_currency;type:varchar(32);default:'';not null;comment:市值单位" nami:"market_value_currency"`
	TotalPlayers        int64  `json:"totalPlayers" gorm:"column:total_players;default:0;not null;comment:总球员数，-1表示没有该字段数据" nami:"total_players"`
	ForeignPlayers      int64  `json:"foreignPlayers" gorm:"column:foreign_players;default:0;not null;comment:非本土球员数，-1表示没有该字段数据" nami:"foreign_players"`
	NationalPlayers     int64  `json:"nationalPlayers" gorm:"column:national_players;default:0;not null;comment:国家队球员数，-1表示没有该字段数据" nami:"national_players"`
}

// 足球球队数据库实体
type NamiSyncFootballTeam struct {
	TBaseTeam
	FootballTeamExtend
}

func (t NamiSyncFootballTeam) TableName() string { return "t_football_team" }

// 足球球队数据库实体
type FootballTeam struct {
	Team
	FootballTeamExtend
}

func (t FootballTeam) TableName() string { return "t_football_team" }

// func (t *FootballTeam) Update(tx *gorm.DB) error {
// 	return tx.Model(&FootballTeam{}).Where("id = ?", t.ID).Updates(t).Error
// }

// func (t *FootballTeam) Create(tx *gorm.DB) error {
// 	return tx.Model(&FootballTeam{}).Create(t).Error
// }

// func (t *FootballTeam) GetID() uint64 {
// 	return t.ID
// }

// type TeamInterface interface {
// 	Update(tx *gorm.DB) error
// 	Create(tx *gorm.DB) error
// 	GetID() uint64
// }
