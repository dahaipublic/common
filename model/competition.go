package model

//import "gorm.io/gorm"

type TBaseCompetition struct {
	BasicField
	CategoryID int64  `json:"categoryID" gorm:"column:category_id;default:0;not null;comment:分类id" nami:"category_id"`
	CountryID  int64  `json:"countryID" gorm:"column:country_id;default:0;not null;comment:国家id" nami:"country_id"`
	NameZh     string `json:"nameZh" gorm:"column:name_zh;type:varchar(128);default:'';not null;comment:中文名称" nami:"name_zh"`
	NameZht    string `json:"nameZht" gorm:"column:name_zht;type:varchar(128);default:'';not null;comment:粤语名称" nami:"name_zht"`
	NameEn     string `json:"nameEn" gorm:"column:name_en;type:varchar(128);default:'';not null;comment:英文名称" nami:"name_en"`

	ShortNameZh   string `json:"shortNameZh" gorm:"column:short_name_zh;type:varchar(128);default:'';not null;comment:中文简称" nami:"short_name_zh"`
	ShortNameZht  string `json:"shortNameZht" gorm:"column:short_name_zht;type:varchar(128);default:'';not null;comment:粤语简称" nami:"short_name_zht"`
	ShortNameEn   string `json:"shortNameEn" gorm:"column:short_name_en;type:varchar(128);default:'';not null;comment:英文简称" nami:"short_name_en"`
	Logo          string `json:"logo" gorm:"column:logo;type:varchar(256);default:'';not null;comment:赛事logo" nami:"logo"`
	Type          int8   `json:"type" gorm:"column:type;default:0;not null;comment:赛事类型:篮球0-未知、1-联赛、2-杯赛,足球0-未知、1-联赛、2-杯赛、3-友谊赛" nami:"type"`
	NamiUpdatedAt int64  `json:"namiUpdatedAt" gorm:"column:nami_updated_at;default:0;not null;comment:纳米更新时间" nami:"updated_at"`
}

type Competition struct {
	TBaseCompetition
	NameTr  string `json:"nameTr" gorm:"column:name_tr;type:varchar(128);default:'';not null;comment:土耳其名称" nami:"name_tr"`
	NameAr  string `json:"nameAr" gorm:"column:name_ar;type:varchar(128);default:'';not null;comment:埃及名称" nami:"name_ar"`
	IsClose int8   `json:"isClose" gorm:"column:is_close;default:0;not null;comment:是否关闭:0否1是"`
}

// 篮球赛事同步实体
type NamiSyncBasketballCompetition struct{ TBaseCompetition }

func (t NamiSyncBasketballCompetition) TableName() string { return "t_basketball_competition" }

// type CompetitionInterface interface {
// 	Update(tx *gorm.DB) error
// 	Create(tx *gorm.DB) error
// 	GetCompetition() Competition
// }

// 篮球赛事全量实体
type BasketballCompetition struct {
	Competition
}

// func (t *BasketballCompetition) Update(tx *gorm.DB) error {
// 	return tx.Model(&BasketballCompetition{}).Where("id = ?", t.ID).Updates(t).Error
// }

// func (t *BasketballCompetition) Create(tx *gorm.DB) error {
// 	return tx.Model(&BasketballCompetition{}).Create(t).Error
// }

// func (t *BasketballCompetition) GetCompetition() Competition {
// 	return t.Competition
// }

func (t BasketballCompetition) TableName() string { return "t_basketball_competition" }

// ----------------------
// 下面字段是老代码未从纳米同步的字段
type FootballCompetitionExtend struct {
	TitleHolder    TJsonString `json:"titleHolder" gorm:"column:title_holder;type:varchar(128);default:'[]';not null;comment:卫冕冠军" nami:"title_holder"`
	MostTitles     TJsonString `json:"mostTitles" gorm:"column:most_titles;type:varchar(128);default:'[]';not null;comment:夺冠最多球队" nami:"most_titles"`
	Newcomers      TJsonString `json:"newcomers" gorm:"column:newcomers;type:varchar(256);default:'[]';not null;comment:晋级淘汰球队" nami:"newcomers"`
	Divisions      TJsonString `json:"divisions" gorm:"column:divisions;type:varchar(128);default:'[]';not null;comment:赛事层级" nami:"divisions"`
	Host           TJsonString `json:"host" gorm:"column:host;type:varchar(64);default:'[]';not null;comment:东道主" nami:"host"`
	PrimaryColor   string      `json:"primaryColor" gorm:"column:primary_color;type:varchar(32);default:'';not null;comment:主颜色" nami:"primary_color"`
	SecondaryColor string      `json:"secondaryColor" gorm:"column:secondary_color;type:varchar(32);default:'';not null;comment:次颜色" nami:"secondary_color"`
}

// 篮球赛事同步实体
type NamiSyncFootballCompetition struct {
	TBaseCompetition
	FootballCompetitionExtend
}

func (t NamiSyncFootballCompetition) TableName() string { return "t_football_competition" }

// 足球赛事表
type FootballCompetition struct {
	Competition
	FootballCompetitionExtend
}

// func (t *FootballCompetition) Update(tx *gorm.DB) error {
// 	return tx.Model(&FootballCompetition{}).Where("id = ?", t.ID).Updates(t).Error
// }

// func (t *FootballCompetition) Create(tx *gorm.DB) error {
// 	return tx.Model(&FootballCompetition{}).Create(t).Error
// }

// func (t *FootballCompetition) GetCompetition() Competition {
// 	return t.Competition
// }

func (t FootballCompetition) TableName() string { return "t_football_competition" }
