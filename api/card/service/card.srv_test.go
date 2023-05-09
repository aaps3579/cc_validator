package service

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/aaps3579/cc-validator/api/card/models"
	"github.com/aaps3579/cc-validator/api/card/repository/mocks"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	cardObj     = models.Card{Number: "4123567812345678", State: "pending"}
	cardJSON, _ = json.Marshal(cardObj)
)

func TestAddCard(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodPost, "/card", strings.NewReader(string(cardJSON)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	repo := &mocks.ICardRepository{}
	repo.On("Save", cardObj).Return("4123567812345678	", nil)
	repo.On("Get", "4123567812345678").Return(models.Card{}, nil)

	h := NewCardServiceHandler(repo)

	if assert.NoError(t, h.AddCard(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
	}
}

func TestListCard(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/card", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	repo := &mocks.ICardRepository{}
	repo.On("GetAll").Return([]models.Card{}, nil)

	h := NewCardServiceHandler(repo)
	if assert.NoError(t, h.ListCards(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
