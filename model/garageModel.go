package model

type CarGarage struct {
	ID            int    `'json:"id"`
	License_plate string `'json:"license_plate"`
	Make          string `'json:"make"`
	Model         string `'json:"model"`
	Color         string `'json:"color"`
	Entry_time    string `'json:"entry_time"`
	Repair_status string `'json:"repair_status"`
}
