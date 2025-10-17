package model

type TBasePlayer struct {
	BasicField
	NameTr       string `json:"nameTr" gorm:"column:name_tr;type:varchar(128);default:'';not null;comment:土耳其名称" nami:"name_tr"`
	NameAr       string `json:"nameAr" gorm:"column:name_ar;type:varchar(128);default:'';not null;comment:埃及名称" nami:"name_ar"`
	NameEn       string `json:"nameEn" gorm:"column:name_en;type:varchar(128);default:'';not null;comment:英文名称" nami:"name_en"`
	ShortNameEn  string `json:"shortNameEn" gorm:"column:short_name_en;type:varchar(128);default:'';not null;comment:英文简称" nami:"short_name_en"`
	Logo         string `json:"logo" gorm:"column:logo;type:varchar(256);default:'';not null;comment:球员logo" nami:"logo"`
	CountryID    int64  `json:"countryID" gorm:"column:country_id;default:0;not null;comment:国家id" nami:"country_id"`
	Nationality  string `json:"nationality" gorm:"column:nationality;type:varchar(256);default:'';not null;comment:国籍" nami:"nationality"`
	NationalLogo string `json:"nationalLogo" gorm:"column:national_logo;type:varchar(256);default:'';not null;comment:球员logo(国家队，可判断球队是国家队时使用)" nami:"national_logo"`
	Birthday     int64  `json:"birthday" gorm:"column:birthday;default:0;not null;comment:生日（0-未知）" nami:"birthday"`
	Age          int64  `json:"age" gorm:"column:age;default:0;not null;comment:年龄" nami:"age"`
	Height       int64  `json:"height" gorm:"column:height;default:0;not null;comment:身高" nami:"height"`
	Weight       int64  `json:"weight" gorm:"column:weight;default:0;not null;comment:体重" nami:"weight"`
	Position     string `json:"position" gorm:"column:position;type:varchar(256);default:'';not null;comment:擅长位置，F-前锋、M-中场、D-后卫、G-守门员、其他为未知" nami:"position"`
}
type FootballPlayer struct {
	MarketValue         int64  `json:"marketValue" gorm:"column:market_value;default:0;not null;comment:市值" nami:"market_value"`
	MarketValueCurrency string `json:"marketValueCurrency" gorm:"column:market_value_currency;type:varchar(256);default:'';not null;comment:市值单位" nami:"market_value_currency"`
	ContractUntil       int64  `json:"contractUntil" gorm:"column:contract_until;default:0;not null;comment:合同截止时间" nami:"contract_until"`
	TBasePlayer
}

func (t FootballPlayer) TableName() string { return "t_football_player" }

type BasketballPlayer struct {
	ContractUntil string `json:"contractUntil" gorm:"column:contract_until;default:'';not null;comment:合同截止时间" nami:"contract_until"`

	TBasePlayer
}

func (t BasketballPlayer) TableName() string { return "t_basketball_player" }
