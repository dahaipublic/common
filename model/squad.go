package model

type TBaseSquad struct {
	BasicField
	//	Squad TJsonString `json:"squad" gorm:"column:squad;type:varchar(128);default:'[]';not null;comment:球队阵容列表" nami:"squad"`
	Squad TJsonString `json:"squad" gorm:"column:squad;type:longtext;comment:球队阵容列表"`
}
type FootballSquad struct {
	TBaseSquad
}
type FootBallSquadResults struct {
	PlayerId       int64  `json:"playerId" gorm:"column:player_id;default:0;not null;comment:球员id" nami:"player_id"`
	Position       string `json:"position" gorm:"column:position;type:varchar(128);default:'';not null;comment:球员位置，F-前锋、M-中场、D-后卫、G-守门员、其他为未知" nami:"position"`
	HasShirtNumber int64  `json:"hasShirtNumber" gorm:"column:has_shirt_number;default:0;not null;comment:是否有球衣号，1-是、0-否" nami:"has_shirt_number"`
	ShirtNumber    int64  `json:"shirtNumber" gorm:"column:shirt_number;default:0;not null;comment:球衣号" nami:"shirt_number"`
	IsCaptain      int64  `json:"isCaptain" gorm:"column:is_captain;default:0;not null;comment:是否队长，1-是、0-否" nami:"is_captain"`
}
type BasketballSquad struct {
	TBaseSquad
}
type BasketballSquadResults struct {
	PlayerId    int64  `json:"playerId" gorm:"column:player_id;default:0;not null;comment:球员id" nami:"player_id"`
	Position    string `json:"position" gorm:"column:position;type:varchar(128);default:'';not null;comment:球员位置，C-中锋、SF-小前锋、PF-大前锋、SG-得分后卫、PG-组织后卫、F-前锋、G-后卫，其它都为未知" nami:"position"`
	ShirtNumber string `json:"shirtNumber" gorm:"column:shirt_number;default:0;not null;comment:球衣号" nami:"shirt_number"`
}

func (t FootballSquad) TableName() string   { return "t_football_squad" }
func (t BasketballSquad) TableName() string { return "t_basketball_squad" }
