package model

type BallEGWhiteList struct {
	BasicField
	ClientIp string `json:"clientIp" gorm:"column:client_ip;type:varchar(64);default:'';not null;comment:请求ip"`
}

func (t BallEGWhiteList) TableName() string { return "t_ball_eg_white_list" }
