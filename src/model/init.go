package model

import (
	"context"

	"github.com/uptrace/bun"
)

func InitModel(db *bun.DB, ctx context.Context) {
	_, err := db.NewCreateTable().
		IfNotExists().
		Model((*Customer)(nil)).
		Exec(ctx)

	if err != nil {
		panic(err)
	}
}

func InitData(db *bun.DB, ctx context.Context) {
	rc, err := db.NewSelect().Model((*Customer)(nil)).Count(ctx)

	if err != nil {
		panic(err)
	}

	if rc == 0 {
		for _, customer := range initialCustomerData {
			_, err := db.NewInsert().Model(&customer).Exec(ctx)
			if err != nil {
				panic(err)
			}
		}
	}
}
