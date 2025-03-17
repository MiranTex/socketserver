package models

import "gorm.io/gorm"

type Cluster struct {
	gorm.Model
	ID          int    `gorm:"column:id"` // ID deve ser exportado (maiúsculo)
	PublicID    string `gorm:"column:public_id"`
	Name        string `gorm:"column:name"`
	AccessToken string `gorm:"column:access_token"`
	IsPublic    bool   `gorm:"column:is_public"`
	Status      bool   `gorm:"column:status"`
	Owner       string `gorm:"column:owner"`
}
