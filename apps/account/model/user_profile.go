package model

import "github.com/cnpythongo/goal/pkg/basic"

type UserProfile struct {
	basic.BaseModel
	UserID   int    `gorm:"index:userprofile_user_id;column:user_id;type:int(11);unique;not null"` // 用户ID
	RealName string `gorm:"column:real_name;type:varchar(50);not null"`                            // 真实姓名
	IDNumber string `gorm:"column:id_number;type:varchar(50);not null"`                            // 身份证号
}

func (p *UserProfile) TableName() string {
	return "account_user_profile"
}

func NewUserProfile() *UserProfile {
	return &UserProfile{}
}

func NewUserProfileList() []*UserProfile {
	return make([]*UserProfile, 0)
}
