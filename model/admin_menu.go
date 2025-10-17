package model

type AdminMenu struct {
	BasicField
	Title    string `json:"title" gorm:"column:title;type:varchar(32);not null;comment:'菜单名称'"`
	Path     string `json:"path" gorm:"column:path;type:varchar(128);not null;comment:'页面路径'"`
	Sort     int32  `json:"sort" gorm:"column:sort;type:int;default:0;not null;comment:'排序'"`
	ParentId int64  `json:"parentId" gorm:"column:parent_id;type:bigint;default:0;not null;comment:'父级id'"`
}

func (t AdminMenu) TableName() string {
	return "t_admin_menu"
}
