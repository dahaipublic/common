package model

// 用户表
type AdminUser struct {
	BasicField
	Username string `json:"username" gorm:"column:username;type:varchar(50);default:'';not null;comment:账号"`

	Password string `json:"-" gorm:"column:password;type:varchar(128);comment:密码"`
	RoleId   uint8  `json:"roleId" gorm:"column:role_id;type:tinyint;default:0;not null;comment:角色id"`
	RoleName string `json:"roleName" gorm:"-"` // 角色名称，不存储在数据库中
}

func (t AdminUser) TableName() string {
	return "t_admin_user"
}
