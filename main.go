package main

import (
	"log"
	"os"
	"strings"

	"github.com/aaps3579/spot-money/api/card"
	"github.com/aaps3579/spot-money/api/card/models"
	"github.com/aaps3579/spot-money/api/card/repository"
	"github.com/aaps3579/spot-money/api/card/service"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/robfig/cron/v3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("spot_money.sqlite"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&models.Card{})
	if err != nil {
		panic("unable to migrate")
	}

	e := echo.New()
	// implement service generated from openapi spec
	var myApi service.CardServiceImpl = service.NewCardServiceHandler(repository.NewCardRepository(db))
	card.RegisterHandlers(e, myApi)

	c := cron.New()
	//
	c.AddFunc("*/5 * * * *", func() {
		// get list of banned countries
		err := godotenv.Load()
		if err != nil {
			return
		}
		s := os.Getenv("BANNED")
		if len(s) == 0 {
			return
		}
		myApi.ProcessCard(strings.Split(s, ","))
	})
	c.Start()
	if err := e.Start(":8080"); err != nil {
		log.Fatal(err)
	}
}
