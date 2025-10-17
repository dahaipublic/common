package model

// 敏感词表
type SensitiveWord struct {
	BasicField
	SensitiveWord       string `json:"sensitiveWord" gorm:"column:sensitive_word;type:varchar(32);default:'';not null;comment:敏感词"`
	SensitiveWordPinyin string `json:"sensitiveWordPinyin" gorm:"column:sensitive_word_pinyin;type:varchar(32);default:'';not null;comment:敏感词拼音"`
	PinyinStatus        int8   `json:"pinyinStatus" gorm:"column:pinyin_status;default:0;not null;comment:0否1是"`
	AdminUserID         int8   `json:"adminUserID" gorm:"column:admin_user_id;default:0;not null;comment:管理员id"`
}

func (t SensitiveWord) TableName() string {
	return "t_sensitive_word"
}
