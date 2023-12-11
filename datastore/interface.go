package datastore

import (
	"go-project/model"

	"gofr.dev/pkg/gofr"
)

type CarInfo interface {
	GetAll(ctx *gofr.Context) (*[]model.CarGarage, error)
	GetByID(ctx *gofr.Context, id int) (*model.CarGarage, error)
	GetByPlateNumber(ctx *gofr.Context, plateNumber string) (*model.CarGarage, error)
	Create(ctx *gofr.Context, carInfo *model.CarGarage) (*model.CarGarage, error)
	Update(ctx *gofr.Context, carInfo *model.CarGarage) (*model.CarGarage, error)
	Delete(ctx *gofr.Context, id int) error
}
