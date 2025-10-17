package model

import (
	"gorm.io/gorm"
)

type Coach struct {
	BasicField
	NameZh        string `json:"nameZh" gorm:"column:name_zh;type:varchar(128);default:'';not null;comment:中文名称" nami:"name_zh"`
	NameZht       string `json:"nameZht" gorm:"column:name_zht;type:varchar(128);default:'';not null;comment:粤语名称" nami:"name_zht"`
	NameEn        string `json:"nameEn" gorm:"column:name_en;type:varchar(128);default:'';not null;comment:英文名称" nami:"name_en"`
	Logo          string `json:"logo" gorm:"column:logo;type:varchar(256);default:'';not null;comment:logo" nami:"logo"`
	TeamID        uint64 `json:"teamID" gorm:"column:team_id;default:0;not null;comment:执教球队id" nami:"team_id"`
	NamiUpdatedAt int64  `json:"namiUpdatedAt" gorm:"column:nami_updated_at;default:0;not null;comment:纳米更新时间" nami:"updated_at"`
}

type CoachInterface interface {
	Update(tx *gorm.DB) error
	Create(tx *gorm.DB) error
	GetID() uint64
}

// 篮球教练表
type BasketballCoach struct {
	Coach
}

func (t BasketballCoach) TableName() string {
	return "t_basketball_coach"
}

func (t *BasketballCoach) Update(tx *gorm.DB) error {
	return tx.Model(&BasketballCoach{}).Where("id = ?", t.ID).Updates(t).Error
}

func (t *BasketballCoach) Create(tx *gorm.DB) error {
	return tx.Model(&BasketballCoach{}).Create(t).Error
}

func (t *BasketballCoach) GetID() uint64 {
	return t.ID
}

// 足球教练表
type FootballCoach struct {
	Coach
	Birthday           int64  `json:"birthday" gorm:"column:birthday;default:0;not null;comment:生日（0-未知）" nami:"birthday"`
	Age                int64  `json:"age" gorm:"column:age;default:0;not null;comment:年龄" nami:"age"`
	PreferredFormation string `json:"preferredFormation" gorm:"column:preferred_formation;type:varchar(64);default:'';not null;comment:习惯的阵型" nami:"preferred_formation"`
	CountryID          int64  `json:"countryID" gorm:"column:country_id;default:0;not null;comment:国家id" nami:"country_id"`
	Nationality        string `json:"nationality" gorm:"column:nationality;type:varchar(64);default:'';not null;comment:国籍" nami:"nationality"`
	Joined             int64  `json:"joined" gorm:"column:joined;default:0;not null;comment:加盟时间" nami:"joined"`
	ContractUntil      uint64 `json:"contractUntil" gorm:"column:contract_until;default:0;not null;comment:合同到期时间" nami:"contract_until"`
}

func (t FootballCoach) TableName() string {
	return "t_football_coach"
}

func (t *FootballCoach) Update(tx *gorm.DB) error {
	return tx.Model(&FootballCoach{}).Where("id = ?", t.ID).Updates(t).Error
}

func (t *FootballCoach) Create(tx *gorm.DB) error {
	return tx.Model(&FootballCoach{}).Create(t).Error
}

func (t *FootballCoach) GetID() uint64 {
	return t.ID
}
