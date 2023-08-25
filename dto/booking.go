package dto

import (
	"gorm.io/gorm"
	"time"
)

type Booking struct {
	gorm.Model
	ID        uint          `gorm:"primaryKey"`
	ChargerId uint          `gorm:"column:charger_id" json:"charger_id"`
	Email     string        `gorm:"column:email"`
	StartTime time.Duration `gorm:"column:start_time" json:"start_time"`
	EndTime   time.Duration `gorm:"column:end_time" json:"end_time"`
	Status    string        `gorm:"column:status"`
}
