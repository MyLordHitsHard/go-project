package handler

import (
	"encoding/json"
	"fmt"
	"go-project/datastore"
	"go-project/model"
	"strconv"

	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"
)

type handler struct {
	store datastore.CarInfo
}

func New(c datastore.CarInfo) handler {
	return handler{store: c}
}

type EntityAlreadyExists struct {
	Entity string
	ID     string
}

func (e EntityAlreadyExists) Error() string {
	return fmt.Sprintf("%s with ID %s already exists", e.Entity, e.ID)
}

func isValidJSON(car model.CarGarage) error {
	byteCar, err := json.Marshal(car)
	valid := json.Valid(byteCar)

	if !valid || err != nil {
		return &errors.Response{Reason: "Invalid JSON"}
	}
	return nil
}

func (h handler) GetAll(ctx *gofr.Context) (interface{}, error) {
	response, err := h.store.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return response, nil
}
func (h handler) GetByID(ctx *gofr.Context) (interface{}, error) {
	id := ctx.PathParam("id")
	if id == "" {
		return nil, errors.MissingParam{Param: []string{"id"}}
	}
	if _, err := ValidateID(id); err != nil {
		return nil, errors.InvalidParam{Param: []string{"id"}}
	}
	// idInt, err := ValidateID(id);
	// if err != nil {
	// 	return nil, errors.InvalidParam{Param: []string{"id"}}
	// }
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return nil, errors.InvalidParam{Param: []string{"id"}}
	}
	response, err := h.store.GetByID(ctx, idInt)

	if err != nil {
		return nil, errors.EntityNotFound{
			Entity: "car",
			ID:     id,
		}
	}
	return response, nil
}

func (h handler) GetByPlateNumber(ctx *gofr.Context) (interface{}, error) {
	plateNumber := ctx.PathParam("license_plate")
	if plateNumber == "" {
		return nil, errors.MissingParam{Param: []string{"license_plate"}}
	}

	response, err := h.store.GetByPlateNumber(ctx, plateNumber)

	if err != nil {
		return nil, errors.EntityNotFound{
			Entity: "car",
			ID:     plateNumber,
		}
	}
	return response, nil
}

func (h handler) Create(ctx *gofr.Context) (interface{}, error) {
	var carInfo model.CarGarage

	// if err != nil {
	// 	return nil, err
	// }

	if err := ctx.Bind(&carInfo); err != nil {
		ctx.Logger.Errorf("error in binding: %v", err)
		return nil, errors.InvalidParam{Param: []string{"body"}}
	}

	if err := isValidJSON(carInfo); err != nil {
		return nil, err
	}

	response, err := h.store.GetByPlateNumber(ctx, carInfo.License_plate)

	if response.License_plate == carInfo.License_plate {
		return nil, EntityAlreadyExists{
			Entity: "car",
			ID:     carInfo.License_plate,
		}
	}

	response, err = h.store.Create(ctx, &carInfo)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (h handler) Update(ctx *gofr.Context) (interface{}, error) {
	plateNumber := ctx.PathParam("license_plate")

	if plateNumber == "" {
		return nil, errors.MissingParam{Param: []string{"license_plate"}}
	}
	// id, err := ValidateID(i)
	// if err != nil {
	// 	return nil, errors.InvalidParam{Param: []string{"id"}}
	// }
	var carInfo model.CarGarage
	if err := ctx.Bind(&carInfo); err != nil {
		ctx.Logger.Errorf("error in binding: %v", err)
		return nil, errors.InvalidParam{Param: []string{"body"}}
	}

	carInfo.License_plate = plateNumber
	response, err := h.store.Update(ctx, &carInfo)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (h handler) Delete(ctx *gofr.Context) (interface{}, error) {
	plateNumber := ctx.PathParam("license_plate")

	if plateNumber == "" {
		return nil, errors.MissingParam{Param: []string{"id"}}
	}
	// id, err := ValidateID(i)
	// if err != nil {
	// 	return nil, errors.InvalidParam{Param: []string{"id"}}
	// }
	response, err := h.store.GetByPlateNumber(ctx, plateNumber)
	if len(response.License_plate) == 0 && err != nil {
		return nil, errors.EntityNotFound{
			Entity: "car",
			ID:     plateNumber,
		}
	}
	if err := h.store.Delete(ctx, plateNumber); err != nil {
		return nil, err
	}

	return "Deleted successfully", nil
}

func ValidateID(id string) (int, error) {
	response, err := strconv.Atoi(id)
	if err != nil {
		return 0, err
	}

	return response, err
}
