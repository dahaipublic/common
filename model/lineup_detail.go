package model

type TFootBallLineupDetail struct {
	BasicField
	Confirmed     int64               `json:"confirmed" gorm:"column:confirmed;default:0;not null;comment:正式阵容 1-是、0-不是" nami:"confirmed"`
	HomeFormation string              `json:"homeFormation" gorm:"column:home_formation;type:varchar(128);default:'';not null;comment:主队阵型" nami:"home_formation"`
	AwayFormation string              `json:"awayFormation" gorm:"column:away_formation;type:varchar(128);default:'';not null;comment:客队阵型" nami:"away_formation"`
	HomeCoachId   int64               `json:"homeCoachId" gorm:"column:home_coach_id;default:0;not null;comment:主队带队教练" nami:"home_coach_id"`
	AwayCoachId   int64               `json:"awayCoachId" gorm:"column:home_coach_id;default:0;not null;comment:客队带队教练" nami:"away_coach_id"`
	HomeColor     int64               `json:"homeColor" gorm:"column:home_color;default:0;not null;comment:主队球衣颜色" nami:"home_color"`
	AwayColor     int64               `json:"awayColor" gorm:"column:away_color;default:0;not null;comment:客队球衣颜色" nami:"away_color"`
	Home          TTeamLineupDetailV2 `json:"home" gorm:"column:home;default:0;not null;comment:主队阵型球员列表" nami:"home"`
	Away          TTeamLineupDetailV2 `json:"Away" gorm:"column:away;default:0;not null;comment:客队阵型球员列表" nami:"Away"`
}
type TTeamLineupDetailV2 struct {
	Id           int8                   `json:"id"`            //球员Id
	TeamId       int8                   `json:"team_id"`       //球队id
	First        int8                   `json:"first"`         //是否首发，1-是、0-否
	Captain      int8                   `json:"captain"`       //是否队长，1-是、0-否
	Name         string                 `json:"name"`          //球员名称
	Logo         string                 `json:"logo"`          //球员logo
	NationalLogo string                 `json:"national_logo"` //球员logo(国家队)
	ShirtNumber  int8                   `json:"shirt_number"`  //球衣号
	Position     string                 `json:"position"`      //球员位置，F前锋、M中场、D后卫、G守门员、其他为未知
	X            int8                   `json:"x"`             //阵容x坐标，总共100
	Y            int8                   `json:"y"`             //阵容y坐标，总共100
	Rating       string                 `json:"rating"`        //评分，10为满分
	Incidents    []TTeamLineupIncidents `json:"incidents"`     //评分，10为满分

}
type TTeamLineupIncidents struct {
	Type       int8                   `json:"type"`        //事件类型，详见状态码->技术统计
	Time       string                 `json:"time"`        //事件发生时间（含加时时间，'A+B':A-比赛时间,B-加时时间）
	Belong     int8                   `json:"belong"`      //发生方，0-中立、1-主队、2-客队
	HomeScore  int8                   `json:"home_score"`  //主队比分
	AwayScore  int8                   `json:"away_score"`  //客队比分
	ReasonType int8                   `json:"reason_type"` //红黄牌、换人事件原因，详见状态码->事件原因（红黄牌、换人事件存在）
	Player     map[string]interface{} `json:"player"`      // 球员信息
}

func (t TFootBallLineupDetail) TableName() string { return "t_football_lineup" }

/*=============================================*/
// 结构体，存储解析后单个球员数据
type BasketballLineupTeamPlayerDetail struct {
	ID           uint64 `json:"id"`
	Name         string `json:"name"`
	Logo         string `json:"logo"`         //球员logo
	Num          string `json:"num"`          //球衣号
	PlayTime     string `json:"play_time"`    // 在场时间
	FGMade       int    `json:"fg_made"`      // 命中次数
	FGAttempt    int    `json:"fg_attempt"`   // 投篮次数
	ThreePMade   int    `json:"three_p_made"` // 三分命中
	ThreePAttemp int    `json:"three_p_att"`  // 三分投篮次数
	FTMade       int    `json:"ft_made"`      // 罚球命中
	FTAttempt    int    `json:"ft_att"`       // 罚球投篮次数
	ORB          int    `json:"orb"`          // 进攻篮板
	DRB          int    `json:"drb"`          // 防守篮板
	TRB          int    `json:"trb"`          // 总篮板
	AST          int    `json:"ast"`          // 助攻
	STL          int    `json:"stl"`          // 抢断
	BLK          int    `json:"blk"`          // 盖帽
	TO           int    `json:"to"`           // 失误
	PF           int    `json:"pf"`           // 个人犯规
	PlusMinus    int    `json:"plus_minus"`   // +/-
	PTS          int    `json:"pts"`          // 得分
	Played       int    `json:"played"`       // 是否出场(1-出场，0-没出场)
	OnCourt      int    `json:"on_court"`     // 是否在场上(0-在场，1-不在场)
	Substitute   int    `json:"substitute"`   // 是否是替补(1-替补，0-首发)
	Position     string `json:"position"`     //位置
}

// 主客队结构
type BasketballLineupPlayer struct {
	MatchId      uint64                             `json:"matchId"` //球员Id
	HomeTeamName string                             `json:"home_team_name"  mapstructure:"home_team_name" msgpack:"home_team_name"`
	HomeTeamLogo string                             `json:"home_team_logo"  mapstructure:"home_team_logo" msgpack:"home_team_logo"`
	AwayTeamName string                             `json:"away_team_name"  mapstructure:"away_team_name" msgpack:"away_team_name"`
	AwayTeamLogo string                             `json:"away_team_logo"  mapstructure:"away_team_logo" msgpack:"away_team_logo"`
	Home         []BasketballLineupTeamPlayerDetail `json:"home"`
	Away         []BasketballLineupTeamPlayerDetail `json:"away"`
	HomeData     BasketballLineupTeamStats          `json:"homeData"`
	AwayData     BasketballLineupTeamStats          `json:"awayData"`
	HomeInjury   []Injury                           `json:"home_injury"  mapstructure:"home_injury" msgpack:"home_injury"`
	AwayInjury   []Injury                           `json:"away_injury"  mapstructure:"away_injury" msgpack:"away_injury"`
	HomeHighest  BasketballLineupHighestStatsDetail `json:"homeHighest"`
	AwayHighest  BasketballLineupHighestStatsDetail `json:"awayHighest"`
}

type Injury struct {
	PlayerName    string `json:"playerName"`
	Logo          string `json:"logo"`
	Position      string `json:"position"`
	Type          int64  `json:"type"`          //类型，1-受伤、2-停赛、0-未知
	MissedMatches int64  `json:"missedMatches"` // 缺失比赛场次
}
type BasketballLineupTeamStats struct {
	FGMade       int `json:"fg_made"`       // 命中次数
	FGAttempt    int `json:"fg_attempt"`    // 投篮次数
	ThreeMade    int `json:"three_made"`    // 三分命中次数
	ThreeAttempt int `json:"three-attempt"` // 三分投篮次数
	FTMade       int `json:"ft_made"`       // 罚球命中次数
	FTAttempt    int `json:"ft_attempt"`    // 罚球投篮次数
	OffRebound   int `json:"off_rebound"`   // 进攻篮板
	DefRebound   int `json:"def_rebound"`   // 防守篮板
	TotalRebound int `json:"total_rebound"` // 总篮板
	Assists      int `json:"assists"`       // 助攻数
	Steals       int `json:"steals"`        // 抢断数
	Blocks       int `json:"blocks"`        // 盖帽数
	Turnovers    int `json:"turnovers"`     // 失误次数
	Fouls        int `json:"fouls"`         // 犯规次数
	Points       int `json:"points"`        // 得分
}

type BasketballLineupHighestStatsDetail struct {
	PTS BasketballLineupHighestStatsPlayerDetail `json:"pts"`
	AST BasketballLineupHighestStatsPlayerDetail `json:"ast"`
	TRB BasketballLineupHighestStatsPlayerDetail `json:"trb"`
}

type BasketballLineupHighestStatsPlayerDetail struct {
	Logo  string `json:"logo"` //球员logo
	Num   string `json:"num"`  //球衣号
	Name  string `json:"name"`
	Score int    `json:"score"`
}
