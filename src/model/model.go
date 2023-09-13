package model

type Customer struct {
	ID        int64  `bun:"id,pk,autoincrement"`
	Name      string `bun:"name,notnull"`
	Role      string
	Email     string
	Phone     string
	Contacted bool
}

var initialCustomerData []Customer = []Customer{
	{
		Name:      "Pedro Picapiedra",
		Role:      "Foreman",
		Email:     "pedro@flinstones.com",
		Phone:     "+491522111111",
		Contacted: false,
	},
	{
		Name:      "Hari Seldon",
		Role:      "Psycho historian",
		Email:     "hari@foundation.com",
		Phone:     "+491522000000",
		Contacted: false,
	},
	{
		Name:      "Werner Heisenberg",
		Role:      "Physicist",
		Email:     "heisenberg@nuclearwinter.com",
		Phone:     "+491522333333",
		Contacted: false,
	},
}
