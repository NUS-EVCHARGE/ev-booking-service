package dto

import (
	"fmt"
	"time"
)

type Booking struct {
	//gorm.Model
	ID        uint      `gorm:"primaryKey" json:"id"`
	ChargerId uint      `gorm:"column:charger_id" json:"charger_id"`
	Email     string    `gorm:"column:email"`
	StartTime time.Time `gorm:"column:start_time" json:"start_time"`
	EndTime   time.Time `gorm:"column:end_time" json:"end_time"`
	Status    string    `gorm:"column:status"`
}

func (Booking) TableName() string {
	return "booking_tab"
}

func (b *Booking) Validate() error {
	if b.StartTime.Unix() < time.Now().Unix() {
		return fmt.Errorf("start time cannot be before current time")
	}
	if b.EndTime.Unix() < time.Now().Unix() {
		return fmt.Errorf("end time cannot be before current time")
	}
	if b.StartTime.Unix() > b.EndTime.Unix() {
		return fmt.Errorf("start time cannot be after end time")
	}
	return nil
}
