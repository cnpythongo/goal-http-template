package model

import "time"

type BaseModel struct {
	ID          int       `gorm:"primaryKey;column:id;type:int(11);not null"`
	IsDeleted   bool      `gorm:"column:is_deleted;type:tinyint(1)"`          // 是否已删除
	CreateAt    time.Time `gorm:"column:create_at;type:datetime;not null"`    // 创建时间
	UpdateAt    time.Time `gorm:"column:update_at;type:datetime;not null"`    // 更新时间
	DeleteAt    time.Time `gorm:"column:delete_at;type:datetime"`             // 删除时间
}
