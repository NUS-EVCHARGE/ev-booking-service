package main

import (
	"encoding/json"
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

type Database struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

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

	var database Database
	secret := os.Getenv("MYSQL_PASSWORD")
	// Parse the JSON data into the struct
	if err := json.Unmarshal([]byte(secret), &database); err != nil {
		logrus.WithField("decodeSecretManager", database).Error("failed to decode value from secret manager")
		return
	}

	if database.Password != "" {
		hostname = "admin:" + database.Password + "@tcp(ev-charger-mysql-db.cdklkqeyoz4a.ap-southeast-1.rds.amazonaws.com:3306)/evc?parseTime=true&charset=utf8mb4"
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
