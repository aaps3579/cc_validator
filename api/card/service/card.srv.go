package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/aaps3579/spot-money/api/card/repository"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"golang.org/x/exp/slices"

	"github.com/aaps3579/spot-money/api/card/models"
	"gorm.io/gorm"
)

type CardServiceImpl interface {
	// list all card
	// (GET /card)
	ListCards(ctx echo.Context) error
	// adds a credit card
	// (POST /card)
	AddCard(ctx echo.Context) error
	// get card by id
	// (GET /card/{id})
	GetCard(ctx echo.Context, id string) error

	// process cards from db for country validation
	ProcessCard([]string)
}

func NewCardServiceHandler(srv repository.ICardRepository) CardServiceImpl {
	return &cardServiceHandler{srv: srv}
}

type cardServiceHandler struct {
	srv repository.ICardRepository
}

func (c cardServiceHandler) ListCards(ctx echo.Context) error {
	cards, err := c.srv.GetAll()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(echo.ErrNotFound.Code)
		}
		return err
	}
	return ctx.JSON(200, cards)
}

var validate *validator.Validate

func (c cardServiceHandler) AddCard(ctx echo.Context) error {
	var crd models.Card
	if err := ctx.Bind(&crd); err != nil {
		return echo.NewHTTPError(echo.ErrBadRequest.Code)
	}
	validate = validator.New()
	err := validate.Struct(crd)
	if err != nil {
		return echo.NewHTTPError(echo.ErrBadRequest.Code, err.Error())
	}
	cFind, _ := c.srv.Get(crd.Number)
	if cFind.Number != "" {
		return echo.NewHTTPError(echo.ErrBadRequest.Code, "card already exist in system")
	}
	crd.State = "pending"
	_, err = c.srv.Save(crd)
	if err != nil {
		return echo.NewHTTPError(echo.ErrBadGateway.Code, err.Error())
	}
	return ctx.JSON(201, crd)
}

func (c cardServiceHandler) GetCard(ctx echo.Context, id string) error {
	card, err := c.srv.Get(ctx.Param("id"))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(echo.ErrNotFound.Code)
		}
		return err
	}
	return ctx.JSON(200, card)

}

func (c cardServiceHandler) ProcessCard(bannedCountries []string) {
	cards, err := c.srv.GetTop5()
	if err != nil {
		log.Info("skip card processing")
		return
	}

	for _, crd := range cards {
		crd.State = "lookup failed"
		resp, err := http.Get(fmt.Sprintf("https://lookup.binlist.net/%s", crd.Number))
		if err != nil {
			log.Info("unable to look for card country")
		} else {
			defer resp.Body.Close()
			b, _ := ioutil.ReadAll(resp.Body)
			var lookup models.LookupBIN
			err = json.Unmarshal(b, &lookup)

			if err != nil || validate.Struct(lookup) != nil {
				log.Info("unable to look for card country")
			} else {
				if len(bannedCountries) == 0 || slices.Index(bannedCountries, lookup.Country.Alpha2) == -1 {
					crd.State = "valid"
				} else {
					err = c.srv.Remove(crd)
					if err != nil {
						log.Info("unable to delete card")
					}
					return
				}
			}
		}
		_, err = c.srv.Save(crd)
		if err != nil {
			log.Info("unable to update card")
		}
	}
}
