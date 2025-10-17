package model

type AdminRole struct {
	BasicField
	RoleName string      `json:"roleName" gorm:"column:role_name;type:varchar(32);default:'';not null;comment:角色名称"`
	MenuIds  TJsonString `json:"menuIds" gorm:"column:menu_ids;type:varchar(128);default:'';not null;comment:角色权限id" `
}

func (t AdminRole) TableName() string {
	return "t_admin_role"
}
