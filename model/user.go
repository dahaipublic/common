package model

import (
	"time"
)

const (
	UserRoleIDTourist   = 1
	UserRoleIDUser      = 8
	UserRoleIDAnchor    = 12
	UserRoleIDAdminUser = 16
)

// 用户表
type User struct {
	BasicField
	DeviceID        string     `json:"deviceID" gorm:"column:device_id;type:varchar(64);default:'';not null;comment:设备号"`
	Nickname        string     `json:"nickname" gorm:"column:nickname;type:varchar(200);comment:昵称"`
	Phone           string     `json:"phone" gorm:"column:phone;type:varchar(20);comment:手机"`
	AreaNum         string     `json:"areaNum" gorm:"column:area_num;type:varchar(10);not null;comment:国家区号"`
	Password        string     `json:"password" gorm:"column:password;type:varchar(128);comment:密码"`
	Sex             int8       `json:"sex" gorm:"column:sex;default:0;comment:性别0:男1:女2:保密"`
	UserStatus      int8       `json:"userStatus" gorm:"column:user_status;comment:用户状态(0:正常,1:禁言)"`
	HeaderImg       string     `json:"headerImg" gorm:"column:header_img;type:varchar(200);default:'';not null;comment:图像"`
	RoleID          int8       `json:"roleID" gorm:"column:role_id;default:1;not null;comment:角色:1游客8普通用户16超管用户"`
	SecondaryRoleID uint64     `json:"secondaryRoleID" gorm:"column:secondary_role_id;default:1;not null;comment:二级角色id:1球迷"`
	Level           int64      `json:"level" gorm:"column:level;default:1;not null;comment:用户等级"`
	Experience      int64      `json:"experience" gorm:"column:experience;default:0;not null;comment:经验值"`
	LastLoginTime   *time.Time `json:"lastLoginTime" gorm:"column:last_login_time;comment:上次登录时间"`
	CreateTime      *time.Time `json:"createTime" gorm:"column:create_time;comment:注册时间"`
	Signature       string     `json:"signature" gorm:"column:signature;type:varchar(256);default:'';not null;comment:个性签名"`
	InviteCode      uint64     `json:"inviteCode" gorm:"column:invite_code;default:0;not null;comment:邀请码"`
}

func (t User) TableName() string {
	return "t_user"
}
