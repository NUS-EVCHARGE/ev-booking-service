package handler

import (
	"ev-booking-service/config"
	"ev-booking-service/controller"
	"ev-booking-service/dto"
	"ev-booking-service/helper"
	userDto "github.com/NUS-EVCHARGE/ev-user-service/dto"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

//	@Summary		Create Booking by user
//	@Description	create booking by user
//	@Tags			booking
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	dto.Booking	"returns a user object"
//	@Router			/booking/create_booking [post]
//	@Param			authentication	header	string	yes	"jwtToken of the user"
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
		c.JSON(http.StatusBadRequest, err)
		return
	}

	err = c.BindJSON(&booking)
	if err != nil {
		// todo: change to common library
		logrus.WithField("err", err).Error("error params")
		c.JSON(http.StatusBadRequest, err)
		return
	}
	booking.Email = user.Email
	//todo: convert to eum
	booking.Status = "waiting"

	err = controller.BookingControllerObj.CreateBooking(booking)
	if err != nil {
		// todo: change to common library
		logrus.WithField("err", err).Error("error creating booking")
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, booking)
	return
}

func CreateResponse(message, body string) map[string]interface{} {
	return map[string]interface{}{
		"message": message,
		"body":    body,
	}
}
