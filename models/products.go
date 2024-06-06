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

	selectAll, err := db.Query("select * from produtos order by id")

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

		p.Id = id
		p.Name = name
		p.Description = description
		p.Price = price
		p.Quantity = quantity

		products = append(products, p)
	}
	defer db.Close()
	return products
}

func CreateNewProduct(name, description string, price float64, quantity int) {
	db := db.ConnectDatabase()

	insertData, err := db.Prepare("insert into produtos(nome, descricao, preco, quantidade) values ($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	insertData.Exec(name, description, price, quantity)
	defer db.Close()
}

func DeleteProduct(id string) {
	db := db.ConnectDatabase()

	productDelete, err := db.Prepare("delete from produtos where id=$1")
	if err != nil {
		panic(err.Error())
	}

	productDelete.Exec(id)
	defer db.Close()
}

func EditProduct(id string) Product {
	db := db.ConnectDatabase()

	productRetrieved, err := db.Query("select * from produtos where id = $1", id)
	if err != nil {
		panic(err.Error())
	}

	actualProduct := Product{}

	for productRetrieved.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = productRetrieved.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}

		actualProduct.Id = id
		actualProduct.Name = name
		actualProduct.Description = description
		actualProduct.Price = price
		actualProduct.Quantity = quantity
	}
	defer db.Close()
	return actualProduct
}

func UpdateProduct(id int, name, description string, price float64, quantity int) {
	db := db.ConnectDatabase()

	UpdateProduct, err := db.Prepare("update produtos set name=$1, descicao=$2, preco=$3, quantidade=$4 where id=$5")
	if err != nil {
		panic(err.Error())
	}

	UpdateProduct.Exec(name, description, price, quantity, id)
	defer db.Close()
}
