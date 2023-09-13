package main

import (
	"context"
	"fmt"
	"udacity_crm/model"

	"udacity_crm/db/sqlite"
)

func main() {
	ctx := context.Background()
	db := sqlite.NewDB()

	model.InitModel(db, ctx)
	model.InitData(db, ctx)

	var customers []model.Customer
	err := db.NewSelect().Model(&customers).OrderExpr("id ASC").Limit(10).Scan(ctx)
	if err != nil {
		panic(err)
	}

	for _, c := range customers {
		fmt.Printf("%v\n", c)
	}

}
