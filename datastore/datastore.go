package datastore

import (
	"database/sql"
	"go-project/model"

	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"
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
		return &model.CarGarage{}, errors.DB{Err: err}
	}

	return &response, nil

}

func (c *carInfo) Update(ctx *gofr.Context, carInfo *model.CarGarage) (*model.CarGarage, error) {
	_, err := ctx.DB().ExecContext(ctx, "UPDATE car-garage SET license_plate=$1, make=$2, model=$3, color=$4, entry_time=$5, repair_status=$6, log_id=$7, log_time=$8 WHERE id=$9",
		carInfo.LicensePlate, carInfo.Make, carInfo.Model, carInfo.Color, carInfo.EntryTime, carInfo.RepairStatus, carInfo.LogID, carInfo.LogTime, carInfo.ID)

	if err != nil {
		return &model.carGarage{}, errors.DB{Err: err}
	}

	return carInfo, nil
}

func (c *carInfo) Delete(ctx *gofr.Context, id int) error {
	_, err := ctx.DB().ExecContext(ctx, "DELETE FROM car-garage WHERE id=$1", id)

	if err != nil {
		return erros.DB{Err: err}
	}
	return nil
}
