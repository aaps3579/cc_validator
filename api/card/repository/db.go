package repository

import (
	"github.com/aaps3579/cc-validator/api/card/models"
	"gorm.io/gorm"
)

//go:generate mockery --name ICardRepository
type ICardRepository interface {
	Save(models.Card) (string, error)
	Get(string) (models.Card, error)
	GetAll() ([]models.Card, error)
	Remove(models.Card) error
	GetTop5() ([]models.Card, error)
}

func NewCardRepository(db *gorm.DB) ICardRepository {
	return &cardDb{
		db: db,
	}
}

type cardDb struct {
	db *gorm.DB
}

func (c cardDb) GetAll() ([]models.Card, error) {
	var cards []models.Card
	if tx := c.db.Find(&cards); tx.Error != nil {
		return nil, tx.Error
	}
	return cards, nil
}

func (c cardDb) Get(number string) (models.Card, error) {
	var _card models.Card
	if tx := c.db.First(&_card, number); tx.Error != nil {
		return models.Card{}, tx.Error
	}
	return _card, nil
}

func (c cardDb) Save(crd models.Card) (string, error) {
	if tx := c.db.Save(&crd); tx.Error != nil {
		return "", tx.Error
	}
	return crd.Number, nil
}

func (c cardDb) Remove(crd models.Card) error {
	if tx := c.db.Delete(&crd); tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (c cardDb) GetTop5() ([]models.Card, error) {
	var cards []models.Card
	if tx := c.db.Find(&cards).Where("state = ?", "pending").Limit(5); tx.Error != nil {
		return nil, tx.Error
	}
	return cards, nil
}
