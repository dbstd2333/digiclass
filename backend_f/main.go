package main

import (
	"log"
	"time"

	"github.com/ansrivas/fiberprometheus/v2"
	_ "github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	_ "gorm.io/plugin/prometheus"
)

type UpdateVersion struct { //app Update
	Uuid      int64  `gorm:"primaryKey"`
	NewestVer string `gorm:"size:20"`
	Platefrom string `gorm:"size:20"`
	Info      string `gorm:"size:50"`
	Time      time.Time
}
type rever struct {
	URL       string    `json:"url"`
	Version   string    `json:"version"`
	Notes     string    `json:"notes"`
	PubDate   time.Time `json:"pub_date"`
	Signature string    `json:"signature"`
}

func main() {
	app := fiber.New()
	prometheus := fiberprometheus.New("my-service-name")
	prometheus.RegisterAt(app, "/metrics")
	dsn := "digiclass:nhZhfep3pK3txs2c@tcp(119.23.69.180:3306)/digiclass?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai"
	Mysql, err := gorm.Open(mysql.Open(dsn), &gorm.Config{PrepareStmt: true})
	sqlDB, err2 := Mysql.DB()
	if err != nil {
		println(err.Error())
	}
	if err2 != nil {
		println(err2)
	}
	app.Use(prometheus.Middleware)
	app.Get("/api/desktop/update", hello(Mysql))

	log.Fatal(app.Listen(":10001"))
}

func hello(c *fiber.Ctx, Mysql *gorm.DB) error {
	var ver UpdateVersion
	Mysql.Table("update_version").Last(ver)
	return c.SendString("I made a ☕ for you!")
}
