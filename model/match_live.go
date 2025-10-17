package model

type MatchLive struct {
	BasicField
	MatchID       uint64 `json:"matchID" gorm:"column:match_id;default:0;not null;comment:比赛id"`
	SportsType    int8   `json:"sportsType" gorm:"column:sports_type;default:1;not null;comment:运动类型:1足球2篮球"`
	CompetitionID uint64 `json:"competitionID" gorm:"column:competition_id;default:0;not null;comment:赛事id"`
	LiveName      string `json:"liveName" gorm:"column:live_name;type:varchar(255);default:'';not null;comment:主播房间名"`
	NickName      string `json:"nickName" gorm:"column:nick_name;type:varchar(255);default:'';not null;comment:主播昵称"`
	LiveStatus    int8   `json:"liveStatus" gorm:"column:live_status;default:1;not null;comment:直播状态:1未开播2直播中3直播结束"`
	StartTime     int64  `json:"startTime" gorm:"column:start_time;default:0;not null;comment:开始直播时间"`
	EndTime       int64  `json:"endTime" gorm:"column:end_time;default:0;not null;comment:结束直播时间"`
	PushUrl       string `json:"pushUrl" gorm:"column:push_url;type:varchar(255);default:'';not null;comment:推流地址"`
	PlayUrl       string `json:"playUrl" gorm:"column:play_url;type:varchar(255);default:'';not null;comment:播放地址"`
	StreamName    string `json:"streamName" gorm:"column:stream_name;type:varchar(32);default:'';not null;comment:stream_name"`
	AnchorUserID  uint64 `json:"anchorUserID" gorm:"column:anchor_user_id;default:0;not null;comment:主播用户id"`
	ImageUrl      string `json:"imageUrl" gorm:"column:image_url;type:varchar(255);default:'';not null;comment:图片地址"`
}

func (t MatchLive) TableName() string {
	return "t_match_live"
}
