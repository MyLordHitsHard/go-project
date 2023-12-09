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

	err := ctx.DB().QueryRowContext(ctx, "SELECT * FROM carGarage WHERE id=$1", id).
		Scan(&response.ID, &response.License_plate, &response.Make, &response.Model, &response.Color, &response.Entry_time, &response.Repair_status)

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
	err := ctx.DB().QueryRowContext(ctx, "INSERT INTO carGarage (license_plate, make, model, color, entry_time, repair_status) VALUES($1, $2, $3, $4, $5, $6)"+
		"RETURNING id, license_plate, make, model, color, entry_time, repair_status", carInfo.License_plate, carInfo.Make, carInfo.Model, carInfo.Color, carInfo.Entry_time, carInfo.Repair_status).Scan(
		&response.ID, &response.License_plate, &response.Make, &response.Model, &response.Color, &response.Entry_time, &response.Repair_status)

	if err != nil {
		return &model.CarGarage{}, errors.DB{Err: err}
	}

	return &response, nil

}

func (c *carInfo) Update(ctx *gofr.Context, carInfo *model.CarGarage) (*model.CarGarage, error) {
	_, err := ctx.DB().ExecContext(ctx, "UPDATE carGarage SET license_plate=$1, make=$2, model=$3, color=$4, entry_time=$5, repair_status=$6 WHERE id=$7",
		carInfo.License_plate, carInfo.Make, carInfo.Model, carInfo.Color, carInfo.Entry_time, carInfo.Repair_status, carInfo.ID)

	if err != nil {
		return &model.CarGarage{}, errors.DB{Err: err}
	}

	return carInfo, nil
}

func (c *carInfo) Delete(ctx *gofr.Context, id int) error {
	_, err := ctx.DB().ExecContext(ctx, "DELETE FROM carGarage WHERE id=$1", id)

	if err != nil {
		return errors.DB{Err: err}
	}
	return nil
}
