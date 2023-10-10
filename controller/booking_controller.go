package controller

import (
	"ev-booking-service/dao"
	"ev-booking-service/dto"
	"fmt"
	evu "github.com/NUS-EVCHARGE/ev-user-service/dto"
)

type BookingController interface {
	CreateBooking(booking dto.Booking, user evu.User) error
	GetBookingInfo(user evu.User) ([]dto.Booking, error)
	GetBookingIdInfo(bookingId uint) (dto.Booking, error)
	UpdateBooking(booking dto.Booking) error
	DeleteBooking(bookingId uint, email string) error
}

type BookingControllerImpl struct {
}

func (b *BookingControllerImpl) CreateBooking(newBooking dto.Booking, user evu.User) error {
	//validation
	if err := newBooking.Validate(); err != nil {
		return err
	}

	bookingList, err := b.GetBookingInfo(user)
	if err != nil {
		return err
	}
	for _, booking := range bookingList {
		if booking.EndTime.Unix() > newBooking.StartTime.Unix() {
			return fmt.Errorf("there are overlapping bookings")
		}
	}

	return dao.Db.CreateBookingEntry(newBooking)
}

func (b *BookingControllerImpl) GetBookingInfo(user evu.User) ([]dto.Booking, error) {
	return dao.Db.GetAllBookingEntry(user.Email)
}

func (b *BookingControllerImpl) GetBookingIdInfo(bookingId uint) (dto.Booking, error) {
	return dao.Db.GetBookingIdEntry(bookingId)
}

func (b *BookingControllerImpl) UpdateBooking(booking dto.Booking) error {
	if err := booking.Validate(); err != nil {
		return err
	}
	return dao.Db.UpdateBookingEntry(booking)
}

func (b *BookingControllerImpl) DeleteBooking(bookingId uint, email string) error {
	return dao.Db.DeleteBookingEntry(dto.Booking{ID: bookingId, Email: email})
}

var (
	BookingControllerObj BookingController
)

func NewBookingController() {
	BookingControllerObj = &BookingControllerImpl{}
}
