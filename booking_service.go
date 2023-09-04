package main

import (
	"ev-booking-service/config"
	"ev-booking-service/controller"
	"ev-booking-service/dao"
	_ "ev-booking-service/docs"
	"ev-booking-service/handler"
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"os"
)

var (
	r *gin.Engine
)

func main() {
	var (
		configFile string
	)
	flag.StringVar(&configFile, "config", "config.yaml", "configuration file of this service")
	flag.Parse()

	// init configurations
	configObj, err := config.ParseConfig(configFile)
	if err != nil {
		logrus.WithField("error", err).WithField("filename", configFile).Error("failed to init configurations")
		return
	}

	var hostname string
	dbHost := os.Getenv("DATABASE_HOST")
	if dbHost != "" {
		hostname = dbHost
	} else {
		hostname = configObj.Dsn // localhost
	}

	err = dao.InitDB(hostname)
	if err != nil {
		logrus.WithField("config", configObj).Error("failed to connect to database")
		return
	}
	controller.NewBookingController()
	InitHttpServer(configObj.HttpAddress)
}

func InitHttpServer(httpAddress string) {
	r = gin.Default()
	registerHandler()

	if err := r.Run(httpAddress); err != nil {
		logrus.WithField("error", err).Errorf("http server failed to start")
	}
}

func registerHandler() {
	// use to generate swagger ui
	//	@BasePath	/api/v1
	//	@title		Booking Service API
	//	@version	1.0
	//	@schemes	http
	r.GET("/booking/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.GET("/booking/home", handler.GetBookingServiceHandler)

	// api versioning
	v1 := r.Group("/api/v1")
	// get user info handler
	v1.POST("/booking", handler.CreateBookingHandler)
	v1.GET("/booking", handler.GetBookingHandler)
	v1.PATCH("/booking/", handler.UpdateBookingHandler)
	v1.DELETE("/booking/:id", handler.DeleteBookingHandler)
}
