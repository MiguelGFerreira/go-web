package models

import "pkg/db"

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Quantity    int
}

func GetAllProducts() []Product {
	db := db.ConnectDatabase()

	selectAll, err := db.Query("select * from produtos")

	if err != nil {
		panic(err.Error())
	}

	p := Product{}
	products := []Product{}

	for selectAll.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err := selectAll.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}

		p.Name = name
		p.Description = description
		p.Price = price
		p.Quantity = quantity

		products = append(products, p)
	}
	defer db.Close()
	return products
}
