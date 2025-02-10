package common

import (
	"time"

	"gorm.io/gorm"
)

// 自增ID主键
type ID struct {
    ID int `json:"id" gorm:"primaryKey"`
}

// 创建、更新时间
type Timestamps struct {
    CreatedAt time.Time `json:"created_at" gorm:"column:created_at;autoCreateTime"`
    UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;autoUpdateTime"`
}

// 软删除
type SoftDeletes struct {
    DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index;column:deleted_at"`
}