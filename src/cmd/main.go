package main

import (
	"context"
	"fmt"
	"udacity_crm/backend"
	"udacity_crm/db"
	"udacity_crm/model"

	"encoding/json"
	"net/http"
	"udacity_crm/db/sqlite"

	"github.com/gorilla/mux"
	"github.com/uptrace/bun"
)

type customer model.Customer
type customers []model.Customer

func (data customers) WriteResponse(w http.ResponseWriter) {
	json.NewEncoder(w).Encode(data)
}

func (data customer) WriteResponse(w http.ResponseWriter) {
	json.NewEncoder(w).Encode(data)
}

var getCustomers backend.Route = backend.Route{
	Url:    "/customers",
	Method: "GET",
}

var getCustomerById backend.Route = backend.Route{
	Url:    "/customer/{ID}",
	Method: "GET",
}

var addCustomer backend.Route = backend.Route{
	Url:    "/customer",
	Method: "POST",
}

var deleteCustomer backend.Route
var getCustomer backend.Route

func setupRoutesWithHandlers(dbServer *bun.DB, ctx context.Context) *mux.Router {
	fromDbGetCustomers := func() backend.ResultData {
		var data customers = db.GetAllCustomers(dbServer, ctx)
		return data
	}
	getCustomers.Handler = backend.NewListDataHandler(fromDbGetCustomers)

	fromDbGetCustomerById := func(id int64) backend.ResultData {
		r := db.GetCustomerById(dbServer, ctx, id)
		fmt.Printf("Result from query: %v\n", r)
		c := customer(r)
		return c
	}
	getCustomerById.Handler = backend.NewElemenbByIdHandler(fromDbGetCustomerById)

	fromDbAddCustomer := func(customer *model.Customer) backend.ResultData {
		r := db.AddCustomer(dbServer, ctx, customer)
		fmt.Printf("Result from query: %v\n", r)
		c := customer(r)
		return c
	}
	addCustomer.Handler = backend.NewElemenbByIdHandler(fromDbAddCustomer)

	routes := []backend.Route{getCustomers, getCustomerById}

	return backend.BuildRouter(routes)
}

func main() {
	ctx := context.Background()
	dbServer := sqlite.NewDB()

	model.InitModel(dbServer, ctx)
	model.InitData(dbServer, ctx)

	router := setupRoutesWithHandlers(dbServer, ctx)
	fmt.Println(getCustomer)
	http.ListenAndServe(":3000", router)
}
