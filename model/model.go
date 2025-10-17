package model

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

// token参数
type Claims struct {
	UserId     int64  `json:"uid"`
	Username   string `json:"nickname"`
	DeviceUUID string `json:"deviceUUID"`
	RoleID     int64  `json:"roleID"`
	jwt.StandardClaims
}

/* 新基础字段 */
type BasicField struct {
	ID        uint64         `json:"id" gorm:"column:id;PRIMARY_KEY;AUTO_INCREMENT" nami:"id"`
	CreatedAt int64          `json:"createdAt" gorm:"column:created_at;type:int(11);default:0;"`
	UpdatedAt int64          `json:"updatedAt" gorm:"column:updated_at;type:int(11);default:0;"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"column:deleted_at;"`
}

func (base *BasicField) BeforeCreate(tx *gorm.DB) (err error) {
	base.CreatedAt = time.Now().Unix()
	base.UpdatedAt = base.CreatedAt
	return
}

func (base *BasicField) BeforeUpdate(tx *gorm.DB) (err error) {
	base.UpdatedAt = time.Now().Unix()
	return
}

// 迁移或创建数据表
func InitTable(db *gorm.DB) {

	db.AutoMigrate(&Activity{})         //活动表
	db.AutoMigrate(&ActivityCategory{}) //活动分类表
	db.AutoMigrate(&AdminUser{})        //管理用户
	db.AutoMigrate(&Advertisement{})    //广告表
	db.AutoMigrate(&Anchor{})           //主播表

	db.AutoMigrate(&AdminRole{})   //后台管理权限
	db.AutoMigrate(&AdminMenu{})   // 菜单栏管理
	db.AutoMigrate(&AnchorApply{}) //主播申请表
	//广告比赛表
	db.AutoMigrate(&AdvertisementMatch{})
	//广告表
	db.AutoMigrate(&AppAdvertisement{})
	//app功能控制
	db.AutoMigrate(&AppFunctionControl{})
	//广告类型表
	db.AutoMigrate(&AdvertisementType{})

	//广告展示统计表
	db.AutoMigrate(&AdvertisementStatistics{})
	//广告统计表
	db.AutoMigrate(&AdvertisementClick{})
	//app版本表
	db.AutoMigrate(&AppVersion{})

	// 字符串内容表; 目前分享功能在用
	db.AutoMigrate(&ConfStrContent{})

	db.AutoMigrate(&BasketballCoach{})       //篮球教练表
	db.AutoMigrate(&BasketballCompetition{}) //篮球赛事表
	db.AutoMigrate(&BasketballMatchAdmin{})  //篮球比赛表
	db.AutoMigrate(&BasketballTeam{})        //篮球比赛表

	db.AutoMigrate(&FootballPlayer{})   //足球球员列表
	db.AutoMigrate(&BasketballPlayer{}) //篮球球员列表

	db.AutoMigrate(&Config{})      //配置表
	db.AutoMigrate(&DownloadLog{}) //下载记录表
	//表情包表
	db.AutoMigrate(&Emoticons{})

	db.AutoMigrate(&GameChannel{})

	db.AutoMigrate(&BallEGCompetition{})
	db.AutoMigrate(&BallEGWhiteList{})
	db.AutoMigrate(&FootballCoach{})        //足球教练表
	db.AutoMigrate(&FootballCompetition{})  //足球赛事表
	db.AutoMigrate(&FootballMatchAdmin{})   //足球比赛表
	db.AutoMigrate(&FootballTeam{})         //足球球队表
	db.AutoMigrate(&HotCompetition{})       //热门赛事表
	db.AutoMigrate(&RecommendCompetition{}) //推荐赛事表
	db.AutoMigrate(&HotMatchAdmin{})        //热门比赛表
	db.AutoMigrate(&HotTeam{})              //热门球队表
	db.AutoMigrate(&FootballSquad{})        //足球球队阵容表
	db.AutoMigrate(&BasketballSquad{})      //篮球球队阵容表

	db.AutoMigrate(&FootballInjury{})   //足球球队伤停列表
	db.AutoMigrate(&BasketballInjury{}) //篮球球队伤停列表
	db.AutoMigrate(&FootballSeason{})   //足球赛季列表
	db.AutoMigrate(&BasketballSeason{}) //篮球赛季列表

	db.AutoMigrate(&SmsCodeRecord{})
	//比赛公告表
	db.AutoMigrate(&MatchAnnouncement{})
	//比赛详情表
	db.AutoMigrate(&MatchInfo{})
	//比赛直播表
	db.AutoMigrate(&MatchLive{})
	//比赛视频表
	db.AutoMigrate(&MatchVideo{})
	//数据渠道表
	db.AutoMigrate(&Openinstall{})
	//意见反馈表
	db.AutoMigrate(&Opinion{})
	//回放表
	db.AutoMigrate(&Playback{})
	//推荐比赛表
	db.AutoMigrate(&RecommendMatch{})
	//敏感词表
	db.AutoMigrate(&SensitiveWord{})
	//统计表
	db.AutoMigrate(&StatisticsActiveUser{})
	//每日统计
	db.AutoMigrate(&DailyStatistics{})
	//Openinstall统计
	//db.AutoMigrate(&StatisticsOpeninstall{})
	//用户表
	db.AutoMigrate(&User{})
	//用户关注表
	db.AutoMigrate(&UserAttention{})
	//用户经验值记录表
	db.AutoMigrate(&UserExperienceLog{})
	//用户等级表
	db.AutoMigrate(&UserLevel{})
	//用户等级记录表
	db.AutoMigrate(&UserLevelLog{})
	//用户登录记录表
	db.AutoMigrate(&UserLoginLog{})
	//用户角色表
	db.AutoMigrate(&UserRole{})
	//用户角色记录表
	db.AutoMigrate(&UserRoleLog{})

	//历史记录
	db.AutoMigrate(&WatchHistory{})
	//用户统计
	db.AutoMigrate(&AppStatistics{})
	//聊天记录
	db.AutoMigrate(&ChatMessage{})
	//

}
