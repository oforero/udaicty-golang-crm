package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
	"udacity_crm/db"
	"udacity_crm/model"

	"github.com/gorilla/mux"
	"github.com/uptrace/bun"
)

var runOnce sync.Once
var router *mux.Router

func initDb() {
	runOnce.Do(
		func() {
			ctx := context.Background()
			dbServer := db.NewSqliteDB()

			bunDB := bun.DB(*dbServer)
			model.InitModel(&bunDB, ctx)
			model.InitData(&bunDB, ctx)

			router = setupRoutesWithHandlers(dbServer, ctx)
		},
	)
}

// Tests happy path of submitting a well-formed GET /customers request
func TestGetCustomersHandler(t *testing.T) {
	initDb()
	handler := getCustomers
	req, err := http.NewRequest(handler.Method, handler.Url, nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	// Checks for 200 status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("getCustomers returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Checks for JSON response
	if ctype := rr.Header().Get("Content-Type"); ctype != "application/json" {
		t.Errorf("Content-Type does not match: got %v want %v",
			ctype, "application/json")
	}
}

// Tests happy path of submitting a well-formed POST /customers request
func TestAddCustomerHandler(t *testing.T) {
	initDb()
	handler := addCustomer

	requestBody := strings.NewReader(`
		{
			"name": "Example Name",
			"role": "Example Role",
			"email": "Example Email",
			"phone": 5550199,
			"contacted": true
		}
	`)

	req, err := http.NewRequest(handler.Method, handler.Url, requestBody)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	// Checks for 201 status code
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("addCustomer returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	// Checks for JSON response
	if ctype := rr.Header().Get("Content-Type"); ctype != "application/json" {
		t.Errorf("Content-Type does not match: got %v want %v",
			ctype, "application/json")
	}
}

// Tests unhappy path of deleting a user that doesn't exist
func TestDeleteInexistingCustomerHandler(t *testing.T) {
	initDb()
	req, err := http.NewRequest("DELETE", "/customers/8", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	// Checks for 404 status code
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("deleteCustomer returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}
}

// Tests unhappy path of deleting a user that doesn't exist
func TestDeleteExistingCustomerHandler(t *testing.T) {
	initDb()
	req, err := http.NewRequest("DELETE", "/customers/2", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	// Checks for 200 status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("deleteCustomer returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

// Tests unhappy path of getting a user that doesn't exist
func TestGetCustomerHandler(t *testing.T) {
	initDb()
	req, err := http.NewRequest("GET", "/customers/8", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	// Checks for 404 status code
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("getCustomer returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}
}
