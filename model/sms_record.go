package model

type SmsCodeRecord struct {
	BasicField
	SmsContent string `json:"smsContent" gorm:"column:sms_content;type:varchar(256);default:'';not null;comment:短信内容" `
	PhoneNum   string `json:"phoneNum" gorm:"column:phone_num;type:varchar(256);default:'';not null;comment:发送号码" `
	//Channel    uint64 `json:"channel" gorm:"column:channel;type:varchar(256);default:'';not null;comment:发送短信渠道" `
	Channel uint64 `json:"channel" gorm:"column:channel;default:0;not null;comment:发送短信渠道"`
}

func (t SmsCodeRecord) TableName() string { return "t_sms_record" }
