package handler

import (
	"ev-booking-service/config"
	"ev-booking-service/controller"
	"ev-booking-service/dto"
	"ev-booking-service/helper"
	"fmt"
	userDto "github.com/NUS-EVCHARGE/ev-user-service/dto"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

// @Summary		Health Check
// @Description 	perform health check status
// @Tags 			Health Check
// @Accept 		json
// @Produce 		json
// @Success 		200	{object}	map[string]interface{}	"returns a welcome message"
func GetBookingServiceHandler(c *gin.Context) {
	c.JSON(http.StatusOK, CreateResponse("Welcome to ev-booking-service"))
	return
}

// @Summary		Create Booking by user
// @Description	create booking by user
// @Tags			booking
// @Accept			json
// @Produce		json
// @Success		200	{object}	dto.Booking	"returns a booking object"
// @Router			/booking [post]
// @Param			authentication	header	string	yes	"jwtToken of the user"
func CreateBookingHandler(c *gin.Context) {
	var (
		user    userDto.User
		booking dto.Booking
	)
	tokenStr := c.GetHeader("Authentication")

	// Get User information
	user, err := helper.GetUser(config.GetUserUrl, tokenStr)
	if err != nil {
		// todo: change to common library
		logrus.WithField("err", err).Error("error getting user")
		c.JSON(http.StatusBadRequest, CreateResponse(fmt.Sprintf("%v", err)))
		return
	}

	err = c.BindJSON(&booking)
	if err != nil {
		// todo: change to common library
		logrus.WithField("err", err).Error("error params")
		c.JSON(http.StatusBadRequest, CreateResponse(fmt.Sprintf("%v", err)))
		return
	}
	booking.Email = user.Email
	//todo: convert to eum
	booking.Status = "waiting"

	err = controller.BookingControllerObj.CreateBooking(booking, user)
	if err != nil {
		// todo: change to common library
		logrus.WithField("err", err).Error("error creating booking")
		c.JSON(http.StatusBadRequest, CreateResponse(fmt.Sprintf("%v", err)))
		return
	}
	c.JSON(http.StatusOK, booking)
	return
}

// @Summary		Get Booking by user
// @Description	get booking by user
// @Tags			booking
// @Accept			json
// @Produce		json
// @Success		200	{object}	dto.Booking	"returns a booking object"
// @Router			/booking [get]
// @Param			authentication	header	string	yes	"jwtToken of the user"
func GetBookingHandler(c *gin.Context) {
	var (
		user        userDto.User
		bookingList []dto.Booking
	)
	tokenStr := c.GetHeader("Authentication")

	// Get User information
	user, err := helper.GetUser(config.GetUserUrl, tokenStr)
	if err != nil {
		// todo: change to common library
		logrus.WithField("err", err).Error("error getting user")
		c.JSON(http.StatusBadRequest, CreateResponse(fmt.Sprintf("%v", err)))
		return
	}

	bookingList, err = controller.BookingControllerObj.GetBookingInfo(user)
	if err != nil {
		// todo: change to common library
		logrus.WithField("err", err).Error("error getting booking")
		c.JSON(http.StatusBadRequest, CreateResponse(fmt.Sprintf("%v", err)))
		return
	}
	c.JSON(http.StatusOK, bookingList)
	return
}

// @Summary		Get Booking info by id
// @Description	get booking info by id
// @Tags			booking
// @Accept			json
// @Produce		json
// @Success		200	{object}	dto.Booking	"returns a booking object"
// @Router			/booking [get]
// @Param			authentication	header	string	yes	"jwtToken of the user"
func GetBookingIdHandler(c *gin.Context) {
	var (
		//user    userDto.User
		booking dto.Booking
	)
	tokenStr := c.GetHeader("Authentication")

	// Get User information
	_, err := helper.GetUser(config.GetUserUrl, tokenStr)
	if err != nil {
		// todo: change to common library
		logrus.WithField("err", err).Error("error getting user")
		c.JSON(http.StatusBadRequest, CreateResponse(fmt.Sprintf("%v", err)))
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, CreateResponse("id but be an integer"))
		return
	}

	booking, err = controller.BookingControllerObj.GetBookingIdInfo(uint(id))
	if err != nil {
		logrus.WithField("err", err).Error("error getting booking")
		c.JSON(http.StatusBadRequest, CreateResponse(fmt.Sprintf("%v", err)))
		return
	}

	c.JSON(http.StatusOK, booking)
	return
}

// @Summary		Create Booking by user
// @Description	create booking by user
// @Tags			booking
// @Accept			json
// @Produce		json
// @Success		200	{object}	dto.Booking	"returns a booking object"
// @Router			/booking [patch]
// @Param			authentication	header	string	yes	"jwtToken of the user"
func UpdateBookingHandler(c *gin.Context) {
	var (
		user    userDto.User
		booking dto.Booking
	)
	tokenStr := c.GetHeader("Authentication")

	// Get User information
	user, err := helper.GetUser(config.GetUserUrl, tokenStr)
	if err != nil {
		// todo: change to common library
		logrus.WithField("err", err).Error("error getting user")
		c.JSON(http.StatusBadRequest, CreateResponse(fmt.Sprintf("%v", err)))
		return
	}
	err = c.BindJSON(&booking)
	if err != nil {
		// todo: change to common library
		logrus.WithField("err", err).Error("error params")
		c.JSON(http.StatusBadRequest, CreateResponse(fmt.Sprintf("%v", err)))
		return
	}
	booking.Email = user.Email
	err = controller.BookingControllerObj.UpdateBooking(booking)
	if err != nil {
		// todo: change to common library
		logrus.WithField("booking", booking).WithField("err", err).Error("error update booking")
		c.JSON(http.StatusBadRequest, CreateResponse(fmt.Sprintf("%v", err)))
		return
	}
	c.JSON(http.StatusOK, booking)
	return
}

// @Summary		Create Booking by user
// @Description	create booking by user
// @Tags			booking
// @Accept			json
// @Produce		json
// @Success		200	{object}	dto.Booking	"returns a booking object"
// @Router			/booking [delete]
// @Param			authentication	header	string	yes		"jwtToken of the user"
// @Param			id				path	int		true	"booking id"
func DeleteBookingHandler(c *gin.Context) {
	var (
		user    userDto.User
		booking dto.Booking
	)
	tokenStr := c.GetHeader("Authentication")

	// Get User information
	user, err := helper.GetUser(config.GetUserUrl, tokenStr)
	if err != nil {
		// todo: change to common library
		logrus.WithField("err", err).Error("error getting user")
		c.JSON(http.StatusBadRequest, CreateResponse(fmt.Sprintf("%v", err)))
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, CreateResponse("id but be an integer"))
	}
	err = controller.BookingControllerObj.DeleteBooking(uint(id), user.Email)
	if err != nil {
		// todo: change to common library
		logrus.WithField("booking", booking).WithField("err", err).Error("error update booking")
		c.JSON(http.StatusBadRequest, CreateResponse(fmt.Sprintf("%v", err)))
		return
	}
	c.JSON(http.StatusOK, CreateResponse("booking deletion success"))
	return
}

func CreateResponse(message string) map[string]interface{} {
	return map[string]interface{}{
		"message": message,
	}
}
