package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type TJsonString string

func (this *TJsonString) UnmarshalJSON(b []byte) error {
	*this = TJsonString(string(b))
	if *this == "" {
		return nil
	}

	strings.Trim(string(*this), "\"")
	return nil
}
func (t TJsonString) GetLineup() (int, error) {

	var obj map[string]interface{}
	err := json.Unmarshal([]byte(t), &obj)
	if err != nil {
		return 0, err
	}
	v, ok := obj["lineup"]
	if !ok {
		return 0, errors.New("lineup 字段不存在")
	}
	// 注意 JSON 数字默认解析为 float64
	if num, ok := v.(float64); ok {
		return int(num), nil
	}
	return 0, errors.New("lineup 字段格式错误")
}
func (this TJsonString) ToInt64Slice2() (intArr []int64) {

	err := json.Unmarshal([]byte(this), &intArr)
	if err != nil {
		fmt.Println("解析失败:", err)
		return this.ToInt64Slice()
	}
	return
}

func (this TJsonString) ToInt64Slice3() (intArr []int64) {
	var tempStr string
	err := json.Unmarshal([]byte(this), &tempStr)
	if err != nil {

		fmt.Println("第一步解析失败:", err)
		return this.ToInt64Slice2()
	}
	err2 := json.Unmarshal([]byte(tempStr), &intArr)
	if err2 != nil {
		fmt.Println("解析失败:", err2)
		return
	}
	return
}
func (this TJsonString) ToInt64Slice() (intArr []int64) {
	if this == "" {
		return
	}

	// 使用逗号分隔字符串
	input := ""
	if (this)[0] == '[' {
		input = (string)(this[1 : len(this)-1])
	} else {
		input = (string)(this)
		input = strings.Trim(input, "\x00")
		input = strings.Trim(input, "\"")
	}

	strArr := strings.Split(input, ",")
	// 创建 int64 数组
	intArr = make([]int64, len(strArr), len(strArr))
	// 将字符串转换为 int64 并存储到数组中
	for i, str := range strArr {
		num, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			panic(err)
		}
		intArr[i] = num
	}
	return
}

type Match struct {
	BasicField
	SeasonID      uint64 `json:"seasonID" gorm:"column:season_id;default:0;not null;comment:赛季id" nami:"season_id"`
	CompetitionID uint64 `json:"competitionID" gorm:"column:competition_id;index:index_competition_id;default:0;not null;comment:赛事id" nami:"competition_id"`
	HomeTeamID    uint64 `json:"homeTeamID" gorm:"column:home_team_id;default:0;not null;comment:主队id" nami:"home_team_id"`
	AwayTeamID    uint64 `json:"awayTeamID" gorm:"column:away_team_id;default:0;not null;comment:客队id" nami:"away_team_id"`
	StatusID      int8   `json:"statusID" gorm:"column:status_id;index:index_status_id;default:0;not null;comment:比赛状态，详见状态码->比赛状态" nami:"status_id"`
	MatchTime     int64  `json:"matchTime" gorm:"column:match_time;index:index_match_time;default:0;not null;comment:比赛时间" nami:"match_time"`
	MatchEndTime  int64  `json:"matchEndTime" gorm:"column:match_end_time;default:0;not null;comment:比赛结束时间(大概的不是准确的)" nami:"match_end_time"`
	Neutral       int64  `json:"neutral" gorm:"column:neutral;default:0;not null;comment:是否中立场，1-是、0-否" nami:"neutral"`

	HomeScores TJsonString `json:"homeScores" gorm:"column:home_scores;type:varchar(128);default:'';not null;comment:比分字段说明" nami:"home_scores"`
	AwayScores TJsonString `json:"awayScores" gorm:"column:away_scores;type:varchar(128);default:'';not null;comment:比分字段说明" nami:"away_scores"`

	HomePosition string      `json:"homePosition" gorm:"column:home_position;type:varchar(128);default:'';not null;comment:主队排名" nami:"home_position"`
	AwayPosition string      `json:"awayPosition" gorm:"column:away_position;type:varchar(128);default:'';not null;comment:客队排名" nami:"away_position"`
	Coverage     TJsonString `json:"coverage" gorm:"column:coverage;type:varchar(128);default:'';not null;comment:动画、情报、阵容 字段" nami:"coverage"`

	VenueID uint64 `json:"venueID" gorm:"column:venue_id;default:0;not null;comment:场馆id" nami:"venue_id"`

	Round TJsonString `json:"round" gorm:"column:round;type:varchar(128);default:'';not null;comment:关联信息" nami:"round"`

	NamiUpdatedAt int64 `json:"namiUpdatedAt" gorm:"column:nami_updated_at;default:0;not null;comment:纳米更新时间" nami:"updated_at"`

	Tbd   int8  `json:"tbd" gorm:"column:tbd;comment:比赛时间是否待定 1.是(有数据时返回)" nami:"tbd"`
	Ended int64 `json:"ended" gorm:"column:ended;comment:结束时间(有数据时返回)" nami:"ended"`

	//去掉指数
	//Asia             string `json:"asia" gorm:"column:asia;type:varchar(64);default:'';not null;comment:亚盘指数"`
	//Bs               string `json:"bs" gorm:"column:bs;type:varchar(64);default:'';not null;comment:大小球指数"`
	//Eu               string `json:"eu" gorm:"column:eu;type:varchar(64);default:'';not null;comment:欧赔指数"`
	PlaybackUrl string `json:"playbackUrl" gorm:"column:playback_url;type:text;comment:回放地址"`

	Lineup      int    `json:"coverage.lineup" gorm:"column:lineup;default:0;not null;comment:阵容信息抽取" nami:"coverage.lineup"`
	VideoSource string `json:"videoSource" gorm:"column:video_source;type:varchar(256);default:'';not null;comment:视频源"`
	IsLineup    int8   `json:"isLineup" gorm:"column:is_lineup;default:0;not null;comment:是否有阵容:1无2有"`
	Snapshot    string `json:"snapshot" gorm:"column:snapshot;type:text;comment:比赛截图"`
}

type MatchAdmin struct {
	BasicField
	SeasonID      uint64 `json:"seasonID" gorm:"column:season_id;default:0;not null;comment:赛季id" nami:"season_id"`
	CompetitionID uint64 `json:"competitionID" gorm:"column:competition_id;index:index_competition_id;default:0;not null;comment:赛事id" nami:"competition_id"`
	HomeTeamID    uint64 `json:"homeTeamID" gorm:"column:home_team_id;default:0;not null;comment:主队id" nami:"home_team_id"`
	AwayTeamID    uint64 `json:"awayTeamID" gorm:"column:away_team_id;default:0;not null;comment:客队id" nami:"away_team_id"`
	StatusID      int8   `json:"statusID" gorm:"column:status_id;index:index_status_id;default:0;not null;comment:比赛状态，详见状态码->比赛状态" nami:"status_id"`
	MatchTime     int64  `json:"matchTime" gorm:"column:match_time;index:index_match_time;default:0;not null;comment:比赛时间" nami:"match_time"`
	MatchEndTime  int64  `json:"matchEndTime" gorm:"column:match_end_time;default:0;not null;comment:比赛结束时间(大概的不是准确的)" nami:"match_end_time"`
	Neutral       int64  `json:"neutral" gorm:"column:neutral;default:0;not null;comment:是否中立场，1-是、0-否" nami:"neutral"`

	HomeScores TJsonString `json:"homeScores" gorm:"column:home_scores;type:varchar(128);default:'';not null;comment:比分字段说明" nami:"home_scores"`
	AwayScores TJsonString `json:"awayScores" gorm:"column:away_scores;type:varchar(128);default:'';not null;comment:比分字段说明" nami:"away_scores"`

	HomePosition string      `json:"homePosition" gorm:"column:home_position;type:varchar(128);default:'';not null;comment:主队排名" nami:"home_position"`
	AwayPosition string      `json:"awayPosition" gorm:"column:away_position;type:varchar(128);default:'';not null;comment:客队排名" nami:"away_position"`
	Coverage     TJsonString `json:"coverage" gorm:"column:coverage;type:varchar(128);default:'';not null;comment:动画、情报、阵容 字段" nami:"coverage"`

	VenueID uint64 `json:"venueID" gorm:"column:venue_id;default:0;not null;comment:场馆id" nami:"venue_id"`

	Round TJsonString `json:"round" gorm:"column:round;type:varchar(128);default:'';not null;comment:关联信息" nami:"round"`

	NamiUpdatedAt int64 `json:"namiUpdatedAt" gorm:"column:nami_updated_at;default:0;not null;comment:纳米更新时间" nami:"updated_at"`

	Tbd   int8  `json:"tbd" gorm:"column:tbd;comment:比赛时间是否待定 1.是(有数据时返回)" nami:"tbd"`
	Ended int64 `json:"ended" gorm:"column:ended;comment:结束时间(有数据时返回)" nami:"ended"`

	//去掉指数
	//Asia             string `json:"asia" gorm:"column:asia;type:varchar(64);default:'';not null;comment:亚盘指数"`
	//Bs               string `json:"bs" gorm:"column:bs;type:varchar(64);default:'';not null;comment:大小球指数"`
	//Eu               string `json:"eu" gorm:"column:eu;type:varchar(64);default:'';not null;comment:欧赔指数"`
	PlaybackUrl      string `json:"playbackUrl" gorm:"column:playback_url;type:text;comment:回放地址"`
	IsClose          int8   `json:"isClose" gorm:"column:is_close;default:0;not null;comment:是否关闭:0否1是"`
	TopStatus        int8   `json:"topStatus" gorm:"column:top_status;default:0;not null;comment:置顶状态:0不置顶1置顶"`
	Lineup           int    `json:"coverage.lineup" gorm:"column:lineup;default:0;not null;comment:阵容信息抽取" nami:"coverage.lineup"`
	VideoSource      string `json:"videoSource" gorm:"column:video_source;type:varchar(256);default:'';not null;comment:视频源"`
	IsLineup         int8   `json:"isLineup" gorm:"column:is_lineup;default:0;not null;comment:是否有阵容:1无2有"`
	ViewCount        int64  `json:"viewCount" gorm:"column:view_count;default:0;not null;comment:观看数"`
	Snapshot         string `json:"snapshot" gorm:"column:snapshot;type:text;comment:比赛截图"`
	HighestOnlineNum int64  `json:"highestOnlineNum" gorm:"column:highestOnlineNum;default:0;not null;comment:最高在线人数"`
}

// type MatchInterface interface {
// 	Update(tx *gorm.DB) error
// 	Create(tx *gorm.DB) error
// 	GetMatch() Match
// 	SetMatch(match Match)
// }

// 篮球比赛表
type BasketballMatch struct {
	Match
	Kind        int8 `json:"kind" gorm:"column:kind;default:0;not null;comment:类型id，1-常规赛、2-季后赛、3-季前赛、4-全明星、5-杯赛、6-附加赛、0-无" nami:"kind"`
	PeriodCount int8 `json:"periodCount" gorm:"column:period_count;default:0;not null;comment:比赛总节数(不包含加时)" nami:"period_count"`

	OverTimeScores TJsonString `json:"overTimeScores" gorm:"column:over_time_scores;type:varchar(128);default:'';not null;comment:加时赛比分字段说明（大于1个加时才有该字段，每一位为1节加时比分）" nami:"over_time_scores"`
}
type BasketballMatchAdmin struct {
	MatchAdmin
	Kind        int8 `json:"kind" gorm:"column:kind;default:0;not null;comment:类型id，1-常规赛、2-季后赛、3-季前赛、4-全明星、5-杯赛、6-附加赛、0-无" nami:"kind"`
	PeriodCount int8 `json:"periodCount" gorm:"column:period_count;default:0;not null;comment:比赛总节数(不包含加时)" nami:"period_count"`

	OverTimeScores TJsonString `json:"overTimeScores" gorm:"column:over_time_scores;type:varchar(128);default:'';not null;comment:加时赛比分字段说明（大于1个加时才有该字段，每一位为1节加时比分）" nami:"over_time_scores"`
}

// func (t *BasketballMatch) Update(tx *gorm.DB) error {
// 	return tx.Model(&BasketballMatch{}).Where("id = ?", t.ID).Updates(t).Error
// }

// func (t *BasketballMatch) Create(tx *gorm.DB) error {
// 	return tx.Model(&BasketballMatch{}).Create(t).Error
// }

// func (t *BasketballMatch) GetMatch() Match {
// 	return t.Match
// }

// func (t *BasketballMatch) SetMatch(match Match) {
// 	t.Match = match
// }

func (t BasketballMatch) TableName() string {
	return "t_basketball_match"
}
func (t BasketballMatchAdmin) TableName() string {
	return "t_basketball_match"
}

// 足球比赛表
type FootballMatch struct {
	Match
	Note      string `json:"note" gorm:"column:note;type:varchar(128);default:'';not null;comment:备注" nami:"note"`
	RefereeID int64  `json:"refereeID" gorm:"column:referee_id;default:0;not null;comment:裁判id" nami:"referee_id"`
	RelatedID int64  `json:"relatedID" gorm:"column:related_id;default:0;not null;comment:双回合中另一回合比赛id" nami:"related_id"`

	AggScore    TJsonString `json:"aggScore" gorm:"column:agg_score;type:varchar(128);default:'';not null;comment:双回合常规时间(包括加时时间)总比分 字段说明" nami:"agg_score"`
	Environment TJsonString `json:"environment" gorm:"column:environment;type:varchar(128);default:'';not null;comment:比赛环境,有数据才有此字段" nami:"environment"`

	HasOt       int8 `json:"hasOt gorm:"column:has_ot;comment:是否有加时 1.是(有数据时返回)" nami:"has_ot""`
	TeamReverse int8 `json:"teamReverse" gorm:"column:team_reverse;comment:是否主客相反 1.是(有数据时返回 eg：相反-官网A vs B、纳米B vs A)" nami:"team_reverse"`
}
type FootballMatchAdmin struct {
	MatchAdmin

	Note      string `json:"note" gorm:"column:note;type:varchar(128);default:'';not null;comment:备注" nami:"note"`
	RefereeID int64  `json:"refereeID" gorm:"column:referee_id;default:0;not null;comment:裁判id" nami:"referee_id"`
	RelatedID int64  `json:"relatedID" gorm:"column:related_id;default:0;not null;comment:双回合中另一回合比赛id" nami:"related_id"`

	AggScore    TJsonString `json:"aggScore" gorm:"column:agg_score;type:varchar(128);default:'';not null;comment:双回合常规时间(包括加时时间)总比分 字段说明" nami:"agg_score"`
	Environment TJsonString `json:"environment" gorm:"column:environment;type:varchar(128);default:'';not null;comment:比赛环境,有数据才有此字段" nami:"environment"`

	HasOt       int8 `json:"hasOt gorm:"column:has_ot;comment:是否有加时 1.是(有数据时返回)" nami:"has_ot""`
	TeamReverse int8 `json:"teamReverse" gorm:"column:team_reverse;comment:是否主客相反 1.是(有数据时返回 eg：相反-官网A vs B、纳米B vs A)" nami:"team_reverse"`
}

// func (t *FootballMatch) Update(tx *gorm.DB) error {
// 	return tx.Model(&FootballMatch{}).Where("id = ?", t.ID).Updates(t).Error
// }

// func (t *FootballMatch) Create(tx *gorm.DB) error {
// 	return tx.Model(&FootballMatch{}).Create(t).Error
// }

// func (t *FootballMatch) GetMatch() Match {
// 	return t.Match
// }

// func (t *FootballMatch) SetMatch(match Match) {
// 	t.Match = match
// }

func (t FootballMatch) TableName() string {
	return "t_football_match"
}
func (t FootballMatchAdmin) TableName() string {
	return "t_football_match"
}
