package model

// 关注表
type UserAttention struct {
	BasicField
	FollowUserID uint64 `json:"followUserID" gorm:"column:follow_user_id;type:int(11);default:0;not null;comment:关注user_id"`
	FansUserID   uint64 `json:"fansUserID" gorm:"column:fans_user_id;type:int(11);default:0;not null;comment:粉丝user_id"`
}

func (t UserAttention) TableName() string {
	return "t_user_attention"
}
