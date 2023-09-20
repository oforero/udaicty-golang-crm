package backend

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"udacity_crm/db"
	"udacity_crm/model"

	"github.com/gorilla/mux"
)

// @Summary get all customers in the database
// @Produce json
// @Success 200 {object} []model.Customer
// @Router /customers [get]
func BuildGetAllCustomersHandler(db *db.DB, ctx context.Context) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Searching the database with (NewListDataHandler): %v", r)
		data := db.GetAllCustomers(ctx)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data)
	}
}

// @Summary get the Customer identified by ID in the database
// @ID get-customer-by-id
// @Produce json
// @Success 200 {object} model.Customer
// @Router /customers [get]
func BuildGetCustomerByIdHandler(db *db.DB, ctx context.Context) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.ParseInt(vars["ID"], 0, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(nil)
		} else {
			data, found := db.GetCustomerById(ctx, id)
			fmt.Printf("Result: %v\n", data)
			if !found {
				w.WriteHeader(http.StatusNotFound)
				w.Write(nil)
			} else {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)

				json.NewEncoder(w).Encode(data)
			}

		}
	}
}

// @Summary Add a customer to the database
// @ID create-customer
// @Produce json
// @Param data body model.Customer true "customer data"
// @Success 200 {object} model.Customer
// @Failure 400
// @Router /customers [post]
func BuildAddCustomerHandler(db *db.DB, ctx context.Context) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		customer := new(model.Customer)
		// Read the HTTP request body
		reqBody, _ := io.ReadAll(r.Body)
		// Encode the request body into a Golang value so that we can work with the data
		json.Unmarshal(reqBody, &customer)
		if _, ok := db.AddCustomer(ctx, customer); ok {
			w.WriteHeader(http.StatusCreated)
		} else {
			w.WriteHeader(http.StatusConflict)
		}
		json.NewEncoder(w).Encode(customer)
	}
}

// @Summary update a customer item by ID
// @ID update-customer-by-id
// @Produce json
// @Param data body model.Customer true "customer data"
// @Success 200 {object} model.Customer
// @Failure 404
// @Router /customers/{id} [PATCH]
func BuildUpdateCustomerHandler(db *db.DB, ctx context.Context) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.ParseInt(vars["ID"], 0, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(nil)
			return
		}
		customer := new(model.Customer)
		// Read the HTTP request body
		reqBody, ioErr := io.ReadAll(r.Body)
		if ioErr != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		// Encode the request body into a Golang value so that we can work with the data
		err = json.Unmarshal(reqBody, &customer)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		customer.ID = id
		fmt.Printf("Updating: %v", customer)
		if _, ok := db.UpdateCustomer(ctx, customer); ok {
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(customer)
		} else {
			w.WriteHeader(http.StatusConflict)
			return
		}
	}
}

// @Summary delete a customer item by ID
// @ID delete-customer-by-id
// @Produce json
// @Param id path string true "customer ID"
// @Success 200 {object} model.Customer
// @Failure 404
// @Router /customers/{id} [delete]
func BuildDeleteCustomerHandler(db *db.DB, ctx context.Context) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.ParseInt(vars["ID"], 0, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(nil)
		} else {
			fmt.Printf("Deleting Customer %v\n", id)
			data, found := db.DeleteCustomerById(ctx, id)
			fmt.Printf("Result: %v\n", data)
			if !found {
				w.WriteHeader(http.StatusNotFound)
				w.Write(nil)
			} else {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)

				json.NewEncoder(w).Encode(data)
			}

		}
	}
}
