package datastore

import (
	"go-project/model"
	"gofr.dev/pkg/gofr"
)


type CarInfo interface {
	GetByID(ctx *gofr.Context, id string) (*model.CarGarage, error)
	Create(ctx *gofr.Context, carInfo *model.CarGarage) (*model.CarGarage, error)
	Update(ctx *gofr.Context, carInfo *model.CarGarage) (*model.CarGarage, error)
	Delete(ctx *gofr.Context, id int) error
}