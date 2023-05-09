// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	models "github.com/aaps3579/spot-money/api/card/models"
	mock "github.com/stretchr/testify/mock"
)

// ICardRepository is an autogenerated mock type for the ICardRepository type
type ICardRepository struct {
	mock.Mock
}

// Get provides a mock function with given fields: _a0
func (_m *ICardRepository) Get(_a0 string) (models.Card, error) {
	ret := _m.Called(_a0)

	var r0 models.Card
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (models.Card, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(string) models.Card); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(models.Card)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields:
func (_m *ICardRepository) GetAll() ([]models.Card, error) {
	ret := _m.Called()

	var r0 []models.Card
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]models.Card, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []models.Card); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Card)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTop5 provides a mock function with given fields:
func (_m *ICardRepository) GetTop5() ([]models.Card, error) {
	ret := _m.Called()

	var r0 []models.Card
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]models.Card, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []models.Card); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Card)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Remove provides a mock function with given fields: _a0
func (_m *ICardRepository) Remove(_a0 models.Card) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(models.Card) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Save provides a mock function with given fields: _a0
func (_m *ICardRepository) Save(_a0 models.Card) (string, error) {
	ret := _m.Called(_a0)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(models.Card) (string, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(models.Card) string); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(models.Card) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewICardRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewICardRepository creates a new instance of ICardRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewICardRepository(t mockConstructorTestingTNewICardRepository) *ICardRepository {
	mock := &ICardRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
