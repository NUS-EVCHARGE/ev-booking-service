package dao

import (
	"ev-booking-service/dto"
	"fmt"
)

type mockDbImpl struct {
	bookingList []dto.Booking
}

func (d *mockDbImpl) UpdateBookingEntry(booking dto.Booking) error {
	if len(d.bookingList) < int(booking.ID) {
		return fmt.Errorf("booking not found")
	}
	d.bookingList[int(booking.ID)] = booking
	return nil
}

func (d *mockDbImpl) CreateBookingEntry(booking dto.Booking) error {
	d.bookingList = append(d.bookingList, booking)
	return nil
}

func (d *mockDbImpl) DeleteBookingEntry(booking dto.Booking) error {
	if len(d.bookingList) < int(booking.ID) || len(d.bookingList) == 0 {
		return fmt.Errorf("booking not found")
	}
	if len(d.bookingList) == 1 {
		d.bookingList = []dto.Booking{}
		return nil
	}
	d.bookingList = d.bookingList[:int(booking.ID)]
	d.bookingList = append(d.bookingList, d.bookingList[int(booking.ID)+1:]...)
	return nil
}

func (d *mockDbImpl) GetAllBookingEntry(email string) ([]dto.Booking, error) {
	var (
		bookingList []dto.Booking
	)
	for _, booking := range d.bookingList {
		if booking.Email == email {
			bookingList = append(bookingList, booking)
		}
	}
	return bookingList, nil
}

func (d *mockDbImpl) GetBookingIdEntry(id uint) (dto.Booking, error) {
	if len(d.bookingList) < int(id) || len(d.bookingList) == 0 {
		return dto.Booking{}, fmt.Errorf("booking not found")
	}
	if len(d.bookingList) == 1 {
		return d.bookingList[0], nil
	}

	return d.bookingList[int(id)], nil
}

func NewMockDatabase(bookingList []dto.Booking) Database {
	return &mockDbImpl{
		bookingList: bookingList,
	}
}
