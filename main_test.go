package main

import (
	"bytes"
	"net/http"
	"testing"
	"time"

	"gofr.dev/pkg/gofr/request"
)

func TestIntegration(t *testing.T) {
	go main()
	time.Sleep(5 * time.Second)

	tests := []struct {
		desc       string
		method     string
		endpoint   string
		statusCode int
		body       []byte
	}{
		// {"register success", http.MethodPost, "signup", http.StatusCreated, []byte(`{"username": "test2","password": "1234567890"}`)},
		// {"login success", http.MethodPost, "login", http.StatusCreated, []byte(`{"username": "test2","password": "1234567890"}`)},
		{"get success", http.MethodGet, "car-garage", http.StatusOK, nil},
		{"create success", http.MethodPost, "car-garage", http.StatusCreated, []byte(`{"license_plate": "HPA123","make": "BMW","model": "X0","color": "black","entry_time": "2021-09-01 00:00:00","repair_status": "Not started" }`)},
		{"get success", http.MethodGet, "car-garage/HPA123", http.StatusOK, nil},
		{"update success", http.MethodPut, "car-garage/HPA123", http.StatusOK, []byte(`{"Color": "white"}`)},
		{"delete success", http.MethodDelete, "car-garage/HPA123", http.StatusNoContent, nil},
		{"get unknown endpoint", http.MethodGet, "car", http.StatusNotFound, nil},
		{"get unknown car", http.MethodGet, "car-garage/BMW555", http.StatusNotFound, nil},
	}

	for i, tc := range tests {
		req, _ := request.NewMock(tc.method, "http://localhost:9000/"+tc.endpoint, bytes.NewBuffer(tc.body))
		c := http.Client{}

		resp, err := c.Do(req)
		if err != nil {
			t.Errorf("TEST[%v] Failed.\tHTTP request encountered Err: %v\n%s", i, err, tc.desc)
			continue // move to next test
		}

		if resp.StatusCode != tc.statusCode {
			t.Errorf("TEST[%v] Failed.\tExpected %v\tGot %v\n%s", i, tc.statusCode, resp.StatusCode, tc.desc)
		}

		_ = resp.Body.Close()
	}
}
