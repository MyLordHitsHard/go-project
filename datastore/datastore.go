package datastore

import (
	"database/sql"
	"go-project/model"
	"strconv"

	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"
)

type carInfo struct{}

func New() *carInfo {
	return &carInfo{}
}

func (c *carInfo) GetAll(ctx *gofr.Context) (*[]model.CarGarage, error) {
	var response []model.CarGarage

	rows, err := ctx.DB().QueryContext(ctx, "SELECT * FROM carGarage")
	if err != nil {
		return nil, errors.DB{Err: err}
	}
	defer rows.Close()
	for rows.Next() {
		var car model.CarGarage
		err := rows.Scan(&car.ID, &car.License_plate, &car.Make, &car.Model, &car.Color, &car.Entry_time, &car.Repair_status)
		if err != nil {
			return nil, errors.DB{Err: err}
		}
		response = append(response, car)
	}
	return &response, nil

}

func (c *carInfo) GetByID(ctx *gofr.Context, id int) (*model.CarGarage, error) {
	var response model.CarGarage

	err := ctx.DB().QueryRowContext(ctx, "SELECT * FROM carGarage WHERE id=?", id).
		Scan(&response.ID, &response.License_plate, &response.Make, &response.Model, &response.Color, &response.Entry_time, &response.Repair_status)

	switch err {
	case sql.ErrNoRows:
		idString := strconv.Itoa(id)
		return &model.CarGarage{}, errors.EntityNotFound{Entity: "car", ID: idString}
	case nil:
		return &response, nil
	default:
		return &model.CarGarage{}, err
	}
}

func (c *carInfo) GetByPlateNumber(ctx *gofr.Context, plateNumber string) (*model.CarGarage, error) {
	var response model.CarGarage

	err := ctx.DB().QueryRowContext(ctx, "SELECT * FROM carGarage WHERE license_plate=?", plateNumber).
		Scan(&response.ID, &response.License_plate, &response.Make, &response.Model, &response.Color, &response.Entry_time, &response.Repair_status)

	switch err {
	case sql.ErrNoRows:
		return &model.CarGarage{}, errors.EntityNotFound{Entity: "car", ID: plateNumber}
	case nil:
		return &response, nil
	default:
		return &model.CarGarage{}, err
	}
}

// func (c *carInfo) Create(ctx *gofr.Context, carInfo *model.CarGarage) (*model.CarGarage, error) {
// 	var response model.CarGarage
// 	err := ctx.DB().QueryRowContext(ctx, "INSERT INTO carGarage (license_plate, make, model, color, entry_time, repair_status) VALUES($1, $2, $3, $4, $5, $6)"+
// 		" RETURNING id, license_plate, make, model, color, entry_time, repair_status", carInfo.License_plate, carInfo.Make, carInfo.Model, carInfo.Color, carInfo.Entry_time, carInfo.Repair_status).Scan(
// 		&response.ID, &response.License_plate, &response.Make, &response.Model, &response.Color, &response.Entry_time, &response.Repair_status)

// 	if err != nil {
// 		return &model.CarGarage{}, errors.DB{Err: err}
// 	}

// 	return &response, nil

// }

// func (c *carInfo) Create(ctx *gofr.Context, carInfo *model.CarGarage) (*model.CarGarage, error) {
// 	var response model.CarGarage
// 	err := ctx.DB().QueryRowContext(ctx, "INSERT INTO carGarage (license_plate, make, model, color, entry_time, repair_status) VALUES($1, $2, $3, $4, $5, $6)", carInfo.License_plate, carInfo.Make, carInfo.Model, carInfo.Color, carInfo.Entry_time, carInfo.Repair_status)

// 	if err != nil {
// 		return &model.CarGarage{}, errors.DB{Err: err}
// 	}

// 	row := ctx.DB().QueryRowContext(ctx, "SELECT LAST_INSERT_ID()")
// 	err = row.Scan(&response.ID)
// 	if err != nil {
// 		return &model.CarGarage{}, errors.DB{Err: err}
// 	}

// 	return &response, nil
// }

// func (c *carInfo) Create(ctx *gofr.Context, carInfo *model.CarGarage) (*model.CarGarage, error) {
// 	var response model.CarGarage
// 	err := ctx.DB().QueryRowContext(ctx, "INSERT INTO carGarage (license_plate, make, model, color, entry_time, repair_status) VALUES($1, $2, $3, $4, $5, $6)"+
// 		" SELECT * from carGarage WHERE id=LAST_INSERT_ID()", carInfo.License_plate, carInfo.Make, carInfo.Model, carInfo.Color, carInfo.Entry_time, carInfo.Repair_status).Scan(
// 		&response.ID, &response.License_plate, &response.Make, &response.Model, &response.Color, &response.Entry_time, &response.Repair_status)
// 	if err != nil {
// 		return &model.CarGarage{}, errors.DB{Err: err}
// 	}
// 	return &response, nil
// }

func (c *carInfo) Create(ctx *gofr.Context, carInfo *model.CarGarage) (*model.CarGarage, error) {
	var response model.CarGarage

	res, err := ctx.DB().ExecContext(ctx, "INSERT INTO carGarage (license_plate, make, model, color, entry_time, repair_status) VALUES(?, ?, ?, ?, ?, ?)", carInfo.License_plate, carInfo.Make, carInfo.Model, carInfo.Color, carInfo.Entry_time, carInfo.Repair_status)
	if err != nil {
		return &model.CarGarage{}, errors.DB{Err: err}
	}
	id, err := res.LastInsertId()
	if err != nil {
		return &model.CarGarage{}, errors.DB{Err: err}
	}
	err = ctx.DB().QueryRowContext(ctx, "SELECT * from carGarage WHERE id=?", id).Scan(&response.ID, &response.License_plate, &response.Make, &response.Model, &response.Color, &response.Entry_time, &response.Repair_status)
	if err != nil {
		return &model.CarGarage{}, errors.DB{Err: err}
	}
	return &response, nil

}

func (c *carInfo) Update(ctx *gofr.Context, carInfo *model.CarGarage) (*model.CarGarage, error) {
	var response model.CarGarage
	err := ctx.DB().QueryRowContext(ctx, "SELECT * FROM carGarage WHERE license_plate=?", carInfo.License_plate).Scan(&response.ID, &response.License_plate, &response.Make, &response.Model, &response.Color, &response.Entry_time, &response.Repair_status)
	if err != nil {
		return &model.CarGarage{}, errors.DB{Err: err}
	}
	// if carInfo.License_plate == "" {
	// 	carInfo.License_plate = response.License_plate
	// }
	if carInfo.Make == "" {
		carInfo.Make = response.Make
	}
	if carInfo.Model == "" {
		carInfo.Model = response.Model
	}
	if carInfo.Color == "" {
		carInfo.Color = response.Color
	}
	if carInfo.Entry_time == "" {
		carInfo.Entry_time = response.Entry_time
	}
	if carInfo.Repair_status == "" {
		carInfo.Repair_status = response.Repair_status
	}

	_, err = ctx.DB().ExecContext(ctx, "UPDATE carGarage SET license_plate=?, make=?, model=?, color=?, entry_time=?, repair_status=? WHERE id=?",
		carInfo.License_plate, carInfo.Make, carInfo.Model, carInfo.Color, carInfo.Entry_time, carInfo.Repair_status, carInfo.ID)

	if err != nil {
		return &model.CarGarage{}, errors.DB{Err: err}
	}

	return carInfo, nil
}

func (c *carInfo) Delete(ctx *gofr.Context, plateNumber string) error {
	_, err := ctx.DB().ExecContext(ctx, "DELETE FROM carGarage WHERE license_plate=?", plateNumber)

	if err != nil {
		return errors.DB{Err: err}
	}
	return nil
}
