package controller

import (
	"ev-booking-service/dao"
	"ev-booking-service/dto"
	"fmt"
	evu "github.com/NUS-EVCHARGE/ev-user-service/dto"
	"time"
)

type BookingController interface {
	CreateBooking(booking dto.Booking) error
	GetBookingInfo(user evu.User) ([]dto.Booking, error)
	UpdateBooking(booking dto.Booking) error
	DeleteBooking(bookingId uint, email string) error
}

type BookingControllerImpl struct {
}

func (b* BookingControllerImpl) CreateBooking(booking dto.Booking) error {
	//validation
	if booking.StartTime.Unix() < time.Now().Unix() {
		return fmt.Errorf("start time cannot be before current time")
	}
	if booking.StartTime.Unix() > booking.EndTime.Unix() {
		return fmt.Errorf("start time cannot be after end time")
	}
	return dao.Db.CreateBookingEntry(booking)
}

func (b* BookingControllerImpl) GetBookingInfo(user evu.User) ([]dto.Booking, error) {
	return dao.Db.GetAllBookingEntry(user.Email)
}

func (b* BookingControllerImpl) UpdateBooking(booking dto.Booking) error {
	if booking.StartTime.Unix() < time.Now().Unix() {
		return fmt.Errorf("start time cannot be before current time")
	}
	if booking.StartTime.Unix() > booking.EndTime.Unix() {
		return fmt.Errorf("start time cannot be after end time")
	}
	return dao.Db.UpdateBookingEntry(booking)
}

func (b* BookingControllerImpl) DeleteBooking(bookingId uint, email string) error {
	return dao.Db.DeleteBookingEntry(dto.Booking{ID: bookingId, Email: email})
}

var (
	BookingControllerObj BookingController
)

func NewBookingController() {
	BookingControllerObj = &BookingControllerImpl{}
}
