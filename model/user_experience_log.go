package model

const (
	UserExperienceTypeDailyOnline      = 1  //每日在线
	UserExperienceTypeSignIn           = 8  //签到
	UserExperienceTypeWatchLive        = 16 //观看直播
	UserExperienceTypeDailyInteractive = 20 //每日互动
	UserExperienceTypeInviteNewUser    = 24 //邀请新用户
	UserExperienceTypeRegister         = 32 //注册
	UserExperienceTypeUpdateHeaderImg  = 40 //更新头像
	UserExperienceTypeUpdateNickname   = 48 //更新昵称
	UserExperienceTypeUpdateSignature  = 56 //更新签名
)

// 用户经验值记录表
type UserExperienceLog struct {
	BasicField
	UserID              uint64 `json:"userID" gorm:"column:user_id;default:0;not null;comment:用户id"`
	FromUserID          uint64 `json:"fromUserID" gorm:"column:from_user_id;default:0;not null;comment:来自用户id"`
	Type                int8   `json:"type" gorm:"column:type;default:0;comment:性别0:男1:女2:保密"`
	BeforeExperienceNum int64  `json:"beforeExperienceNum" gorm:"column:before_experience_num;default:0;not null;comment:变动前"`
	LastExperienceNum   int64  `json:"lastExperienceNum" gorm:"column:last_experience_num;default:0;not null;comment:变动后"`
	Num                 int64  `json:"num" gorm:"column:num;default:0;not null;comment:数量"`
	ContinuousDays      int8   `json:"continuousDays" gorm:"column:continuous_days;default:0;comment:连续天数"`
}

func (t UserExperienceLog) TableName() string {
	return "t_user_experience_log"
}
