package model

// 系统配置
type Config struct {
	Key   string `json:"key" gorm:"column:key;PRIMARY_KEY;type:varchar(128);default:'';not null;comment:参数键名"`
	Value string `json:"value" gorm:"column:value;type:varchar(32);default:'';not null;comment:参数键值"`
}

func (t Config) TableName() string {
	return "t_config"
}
