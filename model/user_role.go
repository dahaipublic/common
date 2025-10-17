package model

// 用户角色表
type UserRole struct {
	BasicField
	Level                           int64  `json:"level" gorm:"column:level;default:0;not null;comment:需要达到的等级"`
	Name                            string `json:"name" gorm:"column:name;type:varchar(32);default:'';not null;comment:名称"`
	ImageUrl                        string `json:"imageUrl" gorm:"column:image_url;type:varchar(256);default:'';not null;comment:等级图片链接"`
	BackgroundImageUrl              string `json:"backgroundImageUrl" gorm:"column:background_image_url;type:varchar(256);default:'';not null;comment:背景图片链接"`
	EnterChatRoomBackgroundImageUrl string `json:"enterChatRoomBackgroundImageUrl" gorm:"column:enter_chat_room_background_image_url;type:varchar(256);default:'';not null;comment:进入聊天室背景图片链接"`
	RoleImageUrl                    string `json:"roleImageUrl" gorm:"column:role_image_url;type:varchar(256);default:'';not null;comment:角色图标链接"`
}

func (t UserRole) TableName() string {
	return "t_user_role"
}
