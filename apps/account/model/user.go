package model

import (
	"database/sql"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"strings"

	"github.com/cnpythongo/goal/pkg/basic"
)

type User struct {
	basic.BaseModel
	Uuid        string       `json:"uuid" gorm:"column:uuid;type:varchar(64);not null;unique;comment:唯一ID"`
	Username    string       `json:"username" gorm:"column:username;type:varchar(256);unique;not null;comment:'用户名'"`
	Password    string       `json:"password" gorm:"column:password;type:varchar(200);not null;comment:'密码'"`
	Salt        string       `json:"salt" gorm:"column:salt;type:varchar(20);not null;comment:'密码加盐'"`
	Email       string       `json:"email" gorm:"column:email;type:varchar(200);default:'';comment:'邮箱'"`
	Avatar      string       `json:"avatar" gorm:"column:avatar;type:varchar(200);default:'';comment:'用户头像'"`
	LastLoginAt sql.NullTime `json:"last_login_at" gorm:"column:last_login_at;type:datetime;comment:'最后登录时间'"`
}

func NewUser() *User {
	return &User{}
}

func NewUsers() []*User {
	return make([]*User, 0)
}

func (u *User) TableName() string {
	return "account_user"
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	us := uuid.NewV4().String()
	u.Uuid = strings.ReplaceAll(us, "-", "")
	return nil
}
