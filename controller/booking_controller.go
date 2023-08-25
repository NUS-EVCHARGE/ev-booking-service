package controller

import (
	"ev-booking-service/dto"
	evu "github.com/NUS-EVCHARGE/ev-user-service/dto"
	"github.com/sirupsen/logrus"
)

type BookingController interface {
	CreateBooking(booking dto.Booking) error
	GetBookingInfo(user evu.User) (dto.Booking, error)
	UpdateBooking(booking dto.Booking) error
	DeleteBooking(booking dto.Booking) error
}

type BookingControllerImpl struct {
}

func (b BookingControllerImpl) CreateBooking(booking dto.Booking) error {
	logrus.Info("create_booking")
	//panic("implement me")
	return nil
}

func (b BookingControllerImpl) GetBookingInfo(user evu.User) (dto.Booking, error) {
	panic("implement me")
}

func (b BookingControllerImpl) UpdateBooking(booking dto.Booking) error {
	panic("implement me")
}

func (b BookingControllerImpl) DeleteBooking(booking dto.Booking) error {
	panic("implement me")
}

var (
	BookingControllerObj BookingController
)

func NewBookingController() {
	BookingControllerObj = &BookingControllerImpl{}
}
