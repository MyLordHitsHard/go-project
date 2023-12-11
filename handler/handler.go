package handler

import (
	"strconv"

	"go-project/datastore"
	"go-project/model"

	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"
)

type handler struct {
	store datastore.CarInfo
}

func New(c datastore.CarInfo) handler {
	return handler{store: c}
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
	plateNumber:= ctx.PathParam("license_plate")
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

	if err := ctx.Bind(&carInfo); err != nil {
		ctx.Logger.Errorf("error in binding: %v", err)
		return nil, errors.InvalidParam{Param: []string{"body"}}
	}

	response, err := h.store.Create(ctx, &carInfo)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (h handler) Update(ctx *gofr.Context) (interface{}, error) {
	i := ctx.PathParam("id")

	if i == "" {
		return nil, errors.MissingParam{Param: []string{"id"}}
	}
	id, err := ValidateID(i)
	if err != nil {
		return nil, errors.InvalidParam{Param: []string{"id"}}
	}
	var carInfo model.CarGarage
	if err = ctx.Bind(&carInfo); err != nil {
		ctx.Logger.Errorf("error in binding: %v", err)
		return nil, errors.InvalidParam{Param: []string{"body"}}
	}

	carInfo.ID = id
	response, err := h.store.Update(ctx, &carInfo)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (h handler) Delete(ctx *gofr.Context) (interface{}, error) {
	i := ctx.PathParam("id")

	if i == "" {
		return nil, errors.MissingParam{Param: []string{"id"}}
	}
	id, err := ValidateID(i)
	if err != nil {
		return nil, errors.InvalidParam{Param: []string{"id"}}
	}

	if err := h.store.Delete(ctx, id); err != nil {
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
