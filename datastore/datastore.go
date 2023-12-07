package datastore

import (
	"database/sql"
	"gofr.dev/pkg/gofr"
	"gofr.dev/pkg/errors"
	"go-project/model"
)

type carInfo struct {}

func New() *carInfo {
	return &carInfo{}
}

func (c *carInfo) GetByID (ctx *gofr.Context, id string) (*model.CarGarage, error) {
	var response model.CarGarage

	err := ctx.DB().QueryRowContext(ctx, "SELECT * FROM car-garage WHERE id=$1", id).
		Scan(&response.ID, &response.LicensePlate, &response.Make, &response.Model, &response.Color, &response.EntryTime, &response.RepairStatus)
	
	switch err {
	case sql.ErrNoRows:
		return &model.CarGarage{}, errors.EntityNotFound{Entity: "car", ID: id}
	case nil:
		return &response, nil
	default:
		return &model.CarGarage{}, err
	}
}