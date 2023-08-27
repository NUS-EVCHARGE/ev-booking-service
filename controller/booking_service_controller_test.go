package controller

import (
	"ev-booking-service/dao"
	"ev-booking-service/dto"
	"fmt"
	userDto "github.com/NUS-EVCHARGE/ev-user-service/dto"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func setup() {
	NewBookingController()
}

func TestCreateBookingSuccess(t *testing.T) {
	setup()
	var (
		actualBooking = dto.Booking{
			ChargerId: 1,
			Email:     "example@example.com",
			StartTime: time.Now(),
			EndTime:   time.Now().Add(15*time.Minute),
			Status:    "",
		}
		user = userDto.User{
			User:  "example",
			Email: "example@example.com",
		}
	)
	dao.Db = dao.NewMockDatabase([]dto.Booking{})
	fmt.Println(actualBooking)
	err := BookingControllerObj.CreateBooking(actualBooking, user)
	assert.Nil(t, err)

	expectedBooking, err := BookingControllerObj.GetBookingInfo(user)
	assert.Nil(t, err)
	assert.Equal(t, actualBooking, expectedBooking[0])
}

func TestCreateBookingThatOverlaps(t *testing.T) {
	setup()
	var (
		actualBooking = dto.Booking{
			ChargerId: 1,
			Email:     "example@example.com",
			StartTime: time.Now(),
			EndTime:   time.Now().Add(15*time.Minute),
			Status:    "",
		}
		user = userDto.User{
			User:  "example",
			Email: "example@example.com",
		}
	)
	dao.Db = dao.NewMockDatabase([]dto.Booking{actualBooking})
	err := BookingControllerObj.CreateBooking(actualBooking, user)
	assert.Equal(t, err, fmt.Errorf("there are overlapping bookings"))
}

func TestCreateBookingWhereStartTimeAfterEndTime(t *testing.T) {
	setup()
	var (
		actualBooking = dto.Booking{
			ChargerId: 1,
			Email:     "example@example.com",
			StartTime: time.Now().Add(15*time.Minute),
			EndTime:   time.Now(),
			Status:    "",
		}
		user = userDto.User{
			User:  "example",
			Email: "example@example.com",
		}
	)
	dao.Db = dao.NewMockDatabase([]dto.Booking{actualBooking})
	err := BookingControllerObj.CreateBooking(actualBooking, user)
	assert.Equal(t, err, fmt.Errorf("start time cannot be after end time"))
}

func TestCreateBookingWhereStartTimeBeforeCurrentTime(t *testing.T) {
	setup()
	var (
		actualBooking = dto.Booking{
			ChargerId: 1,
			Email:     "example@example.com",
			StartTime: time.Date(2023,8,27,10,50,0, 0,time.Local),
			EndTime:   time.Now().Add(15*time.Minute),
			Status:    "",
		}
		user = userDto.User{
			User:  "example",
			Email: "example@example.com",
		}
	)
	dao.Db = dao.NewMockDatabase([]dto.Booking{actualBooking})
	time.Sleep(time.Second)
	err := BookingControllerObj.CreateBooking(actualBooking, user)
	assert.Equal(t, err, fmt.Errorf("start time cannot be before current time"))
}

func TestGetBookingWhereUserDoesNotHaveBooking(t *testing.T) {
	setup()
	var (
		actualBooking = dto.Booking{
			ChargerId: 1,
			Email:     "example@example.com",
			StartTime: time.Date(2023,8,27,10,50,0, 0,time.Local),
			EndTime:   time.Now().Add(15*time.Minute),
			Status:    "",
		}
		user = userDto.User{
			User:  "example",
			Email: "example1@example.com",
		}
	)
	dao.Db = dao.NewMockDatabase([]dto.Booking{actualBooking})
	time.Sleep(time.Second)
	bookingList, err := BookingControllerObj.GetBookingInfo(user)
	assert.Nil(t, err)
	assert.Equal(t, len(bookingList), 0)
}

func TestUpdateBookingSuccess(t *testing.T) {
	setup()
	var (
		actualBooking = dto.Booking{
			ID: 0,
			ChargerId: 1,
			Email:     "example@example.com",
			StartTime: time.Now(),
			EndTime:   time.Now().Add(15*time.Minute),
			Status:    "",
		}
		user = userDto.User{
			User:  "example",
			Email: "example@example.com",
		}
	)
	dao.Db = dao.NewMockDatabase([]dto.Booking{actualBooking})
	time.Sleep(time.Second)
	actualBooking.StartTime = time.Now().Add(15*time.Minute)
	actualBooking.EndTime = time.Now().Add(30*time.Minute)
	err := BookingControllerObj.UpdateBooking(actualBooking)
	assert.Nil(t, err)
	expectedBooking, err := BookingControllerObj.GetBookingInfo(user)
	assert.Nil(t, err)
	assert.Equal(t, actualBooking, expectedBooking[0])
}

func TestUpdateBookingWhereStartTimeBeforeCurrentTime(t *testing.T) {
	setup()
	var (
		actualBooking = dto.Booking{
			ID: 0,
			ChargerId: 1,
			Email:     "example@example.com",
			StartTime: time.Date(2023,8,27,10,50,0, 0,time.Local),
			EndTime:   time.Now().Add(15*time.Minute),
			Status:    "",
		}
	)
	dao.Db = dao.NewMockDatabase([]dto.Booking{actualBooking})
	time.Sleep(time.Second)

	actualBooking.StartTime = time.Date(2023,8,27,10,50,0, 0,time.Local)
	actualBooking.EndTime = time.Now().Add(30*time.Minute)
	err := BookingControllerObj.UpdateBooking(actualBooking)
	assert.Equal(t, err, fmt.Errorf("start time cannot be before current time"))
}

func TestUpdateBookingWhereEndTimeBeforeStartTime(t *testing.T) {
	setup()
	var (
		actualBooking = dto.Booking{
			ID: 0,
			ChargerId: 1,
			Email:     "example@example.com",
			StartTime: time.Date(2023,8,27,10,50,0, 0,time.Local),
			EndTime:   time.Now().Add(15*time.Minute),
			Status:    "",
		}
	)
	dao.Db = dao.NewMockDatabase([]dto.Booking{actualBooking})
	time.Sleep(time.Second)

	actualBooking.StartTime = time.Now().Add(30*time.Minute)
	actualBooking.EndTime = time.Now()
	err := BookingControllerObj.UpdateBooking(actualBooking)
	assert.Equal(t, err, fmt.Errorf("start time cannot be after end time"))
}

func TestUpdateBookingWhereEndTimeBeforeCurrentTime(t *testing.T) {
	setup()
	var (
		actualBooking = dto.Booking{
			ID: 0,
			ChargerId: 1,
			Email:     "example@example.com",
			StartTime: time.Date(2023,8,27,10,50,0, 0,time.Local),
			EndTime:   time.Now().Add(15*time.Minute),
			Status:    "",
		}
	)
	dao.Db = dao.NewMockDatabase([]dto.Booking{actualBooking})
	time.Sleep(time.Second)

	actualBooking.StartTime = time.Now()
	actualBooking.EndTime = time.Date(2023,8,27,10,50,0, 0,time.Local)
	err := BookingControllerObj.UpdateBooking(actualBooking)
	assert.Equal(t, err, fmt.Errorf("end time cannot be before current time"))
}

func TestDeleteBookingSuccess(t *testing.T) {
	setup()
	var (
		actualBooking = dto.Booking{
			ID: 0,
			ChargerId: 1,
			Email:     "example@example.com",
			StartTime: time.Now(),
			EndTime:   time.Now().Add(15*time.Minute),
			Status:    "",
		}
		user = userDto.User{
			User:  "example",
			Email: "example@example.com",
		}
	)
	dao.Db = dao.NewMockDatabase([]dto.Booking{actualBooking})
	time.Sleep(time.Second)
	err := BookingControllerObj.DeleteBooking(0, user.Email)
	assert.Nil(t, err)

	expectedBooking, err := BookingControllerObj.GetBookingInfo(user)
	assert.Nil(t, err)
	assert.Equal(t, len(expectedBooking), 0)
}