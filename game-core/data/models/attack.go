package models

import "gorm.io/gorm"

type Attacks struct {
	gorm.Model
	AttackType string
}
