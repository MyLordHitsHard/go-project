package datastore

import (
	"database/sql"
	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"
	"go-project/model"
)

type carInfo struct{}

func New() *carInfo {
	return &carInfo{}
}

func (c *carInfo) GetByID(ctx *gofr.Context, id string) (*model.CarGarage, error) {
	var response model.CarGarage

	err := ctx.DB().QueryRowContext(ctx, "SELECT * FROM car-garage WHERE id=$1", id).
		Scan(&response.ID, &response.license_plate, &response.make, &response.model, &response.color, &response.entry_time, &response.repair_status)

	switch err {
	case sql.ErrNoRows:
		return &model.CarGarage{}, errors.EntityNotFound{Entity: "car", ID: id}
	case nil:
		return &response, nil
	default:
		return &model.CarGarage{}, err
	}
}

func (c *carInfo) Create(ctx *gofr.Context, carInfo *model.CarGarage) (*model.CarGarage, error) {
	var response model.CarGarage
	err := ctx.DB().QueryRowContext(ctx, "INSERT INTO car-garage (license_plate, make, model, color, entry_time, repair_status, log_id, log_time) VALUES($1, $2, $3, $4, $5, $6, $7, $8)"+
		"RETURNING id, license_plate, make, model, color, entry_time, repair_status, log_id, log_time", carInfo.LicensePlate, carInfo.Make, carInfo.Model, carInfo.Color, carInfo.EntryTime, carInfo.RepairStatus, carInfo.LogID, carInfo.LogTime).Scan(
		&response.ID, &response.license_plate, &response.make, &response.model, &response.color, &response.entry_time, &response.repair_status, &response.log_id, &response.log_time)

	if err != nil {
		return &model.CarGarage{}, err
	}

	return &response, nil

}
