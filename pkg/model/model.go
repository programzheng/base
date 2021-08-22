package model

type Tabler interface {
	TableName() string
}

type Orm interface {
	init()
	Add() (interface{}, error)
}

type Model struct {
	Tabler
	Orm
}

func (m Model) TableName() string {
	return ""
}

func (m Model) init() {
	if !DB.HasTable(&m) {
		DB.CreateTable(&m)
	}
}

func (m Model) Add() (Model, error) {
	if err := DB.Create(&m).Error; err != nil {
		return Model{}, err
	}

	return m, nil
}
