package model

type Garage struct {
	ID            int    `'json:"id"`
	license_plate string `'json:"license_plate"`
	make          string `'json:"make"`
	model         string `'json:"model"`
	color         string `'json:"color"`
	entry_time    string `'json:"entry_time"`
	repair_status string `'json:"repair_status"`
	log_id   int    `'json:"log_id"`
	log_time string `'json:"log_time"`
}
