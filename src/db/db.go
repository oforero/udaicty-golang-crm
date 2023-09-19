package db

import (
	"context"
	"udacity_crm/model"

	"github.com/uptrace/bun"
)

func GetAllCustomers(db *bun.DB, ctx context.Context) []model.Customer {
	var customers []model.Customer
	err := db.NewSelect().
		Model(&customers).
		OrderExpr("id ASC").
		Scan(ctx)
	if err != nil {
		panic(err)
	}
	return customers
}

func GetCustomerById(db *bun.DB, ctx context.Context, id int64) model.Customer {
	// customer := new(model.Customer)
	var customers []model.Customer
	err := db.NewSelect().
		Model(&customers).
		// Model(customer).
		// Limit(1).
		// Where("id = ?", id).
		Scan(ctx)
	if err != nil {
		panic(err)
	}
	return customers[0]
}

func AddCustomer(db *bun.DB, ctx context.Context, customer *model.Customer) *model.Customer {
	sqlr, err := db.NewInsert().Model(customer).Exec(ctx)
	println(sqlr)
	if err != nil {
		panic(err)
	}
	return customer
}
