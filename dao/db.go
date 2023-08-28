package dao

import (
	"ev-booking-service/dto"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database interface {
	CreateBookingEntry(booking dto.Booking) error
	UpdateBookingEntry(booking dto.Booking) error
	DeleteBookingEntry(booking dto.Booking) error
	GetAllBookingEntry(email string) ([]dto.Booking, error)
}
var (
	Db Database
)

type dbImpl struct {
	Dsn string
	DbController *gorm.DB
}

func (d *dbImpl) UpdateBookingEntry(booking dto.Booking) error {
	results := d.DbController.Model(booking).Updates(booking)
	if results.RowsAffected == 0 {
		return fmt.Errorf("booking not found")
	}
	return results.Error
}

func (d*dbImpl) CreateBookingEntry(booking dto.Booking) error {
	result := d.DbController.Create(&booking)
	return result.Error
}

func (d*dbImpl) DeleteBookingEntry(booking dto.Booking) error {
	results := d.DbController.Delete(&booking)
	if results.RowsAffected == 0 {
		return fmt.Errorf("booking not found")
	}
	return results.Error
}

func (d*dbImpl) GetAllBookingEntry(email string) ([]dto.Booking, error) {
	var existingBooking []dto.Booking

	results := d.DbController.Find(&existingBooking, "email = ?", email)
	return existingBooking, results.Error
}

func InitDB(dsn string) error {
	if dbObj, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		return err
	} else {
		Db = NewDatabase(dsn, dbObj)
		return nil
	}
}

func NewDatabase(dsn string, dbObj *gorm.DB) Database {
	return &dbImpl{
		Dsn:          dsn,
		DbController: dbObj,
	}
}
