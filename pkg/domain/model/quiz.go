package model

import "gorm.io/gorm"

type Quiz struct {
	gorm.Model

	ID      uint   `gorm:"primary_key"`
	Problem string `gorm:"type:varchar(100);not null;unique_index"`
	Answer  string `gorm:"type:varchar(100);not null;unique_index"`
}
