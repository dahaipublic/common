package model

// app功能控制表
type AppFunctionControl struct {
	BasicField
	PID     uint64 `json:"pid" gorm:"column:pid;default:0;not null;comment:父级id"`
	Name    string `json:"name" gorm:"column:name;type:varchar(32);not null;comment:名称"`
	Keyword string `json:"keyword" gorm:"column:keyword;type:varchar(32);not null;comment:关键字"`
	OnOff   uint64 `json:"onOff" gorm:"column:on_off;default:0;not null;comment:开关:0关1开"`
}

func (t AppFunctionControl) TableName() string {
	return "t_app_function_control"
}
