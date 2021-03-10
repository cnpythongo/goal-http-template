package model

import "time"

type User struct {
	BaseModel
	UUID        string    `gorm:"column:uuid;type:varchar(64);not null"`      // 用户uuid
	Phone       string    `gorm:"column:phone;type:varchar(50);not null"`     // 手机号
	Username    string    `gorm:"column:username;type:varchar(200);not null"` // 用户名
	Email       string    `gorm:"column:email;type:varchar(200);not null"`    // 邮箱
	Password    string    `gorm:"column:password;type:varchar(200);not null"` // 密码
	Salt        string    `gorm:"column:salt;type:varchar(20);not null"`      // 密码加盐
	Avatar      string    `gorm:"column:avatar;type:varchar(200);not null"`   // 用户头像
	LastLoginAt time.Time `gorm:"column:last_login_at;type:datetime"`         // 最后登录时间
}

func (m *User) TableName() string {
	return "account_user"
}

func NewUser() *User {
	return &User{}
}

func NewUsers() *[]User {
	return &[]User{}
}

type UserProfile struct {
	BaseModel
	UserID   int    `gorm:"index:userprofile_user_id;column:user_id;type:int(11);not null"` // 用户ID
	RealName string `gorm:"column:real_name;type:varchar(50);not null"`                     // 真实姓名
	IDNumber string `gorm:"column:id_number;type:varchar(50);not null"`                     // 身份证号
}

func (m *UserProfile) TableName() string {
	return "account_user_profile"
}

func NewUserProfile() *UserProfile {
	return &UserProfile{}
}

func NewUserProfiles() *[]UserProfile {
	return &[]UserProfile{}
}

type LoginHistory struct {
	BaseModel
	UserID int `gorm:"index:loginhistory_user_id;column:user_id;type:int(11);not null"` // 用户ID
}

func (m *LoginHistory) TableName() string {
	return "account_login_history"
}

func NewLoginHistory() *LoginHistory {
	return &LoginHistory{}
}

func NewLoginHistories() *[]LoginHistory {
	return &[]LoginHistory{}
}
