package db

import (
	"context"
	"database/sql"
	"fmt"
	"udacity_crm/db/sqlite"
	"udacity_crm/model"

	"github.com/uptrace/bun"
)

type DB bun.DB

func NewSqliteDB() *DB {
	sqlite := sqlite.NewDB()
	return (*DB)(sqlite)
}

type CustomerDB interface {
	GetAllCustomers(ctx context.Context) []model.Customer
	GetCustomerById(ctx context.Context, id int64) model.Customer
	AddCustomer(ctx context.Context, customer *model.Customer) *model.Customer
}

func (db *DB) GetAllCustomers(ctx context.Context) []model.Customer {
	bunDB := bun.DB(*db)
	var customers []model.Customer
	err := bunDB.NewSelect().
		Model(&customers).
		OrderExpr("id ASC").
		Scan(ctx)
	if err != nil {
		panic(err)
	}
	return customers
}

func (db *DB) GetCustomerById(ctx context.Context, id int64) (*model.Customer, bool) {
	bunDB := bun.DB(*db)
	var customer *model.Customer = new(model.Customer)

	err := bunDB.NewSelect().
		Model(customer).
		Where("id = ?", id).
		Scan(ctx)
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}
	if err == sql.ErrNoRows {
		return nil, false
	}
	return customer, true
}

func (db *DB) AddCustomer(ctx context.Context, customer *model.Customer) (*model.Customer, bool) {
	bunDB := bun.DB(*db)
	sqlr, err := bunDB.NewInsert().Model(customer).Exec(ctx)
	println(sqlr)
	if err != nil {
		return nil, false
	}
	return customer, true
}

func (db *DB) UpdateCustomer(ctx context.Context, customer *model.Customer) (*model.Customer, bool) {
	bunDB := bun.DB(*db)
	sqlr, err := bunDB.NewUpdate().Model(customer).WherePK().Exec(ctx)
	fmt.Printf("Updating result %v/n", sqlr)
	if err != nil {
		return nil, false
	}
	return customer, true
}

func (db *DB) DeleteCustomerById(ctx context.Context, id int64) (*model.Customer, bool) {
	bunDB := bun.DB(*db)
	var customer *model.Customer = new(model.Customer)
	customer.ID = id

	sqlr, err := bunDB.NewDelete().
		Model(customer).
		WherePK().
		Exec(ctx)

	rows, sqlErr := sqlr.RowsAffected()
	if err != nil || sqlErr != nil || rows == 0 {
		return nil, false
	}
	return customer, true

}
