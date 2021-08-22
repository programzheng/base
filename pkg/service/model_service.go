package service

import (
	"base/pkg/model/billing"

	"github.com/jinzhu/copier"
)

type ModelService struct {
}

func (ms ModelService) Add(m billing.Billing) (ModelService, error) {
	copier.Copy(&ms, &ms)
	result, err := m.Add()

	if err != nil {
		return ms, err
	}
	modelService := ms
	copier.Copy(&modelService, &result)

	return modelService, nil
}
