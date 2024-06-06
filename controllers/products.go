package controllers

import (
	"net/http"
	"pkg/models"
	"strconv"
	"text/template"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	allProducts := models.GetAllProducts()
	temp.ExecuteTemplate(w, "Index", allProducts)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		convertedPrice, err := strconv.ParseFloat(price, 64)
		if err != nil {
			panic(err.Error())
		}

		convertedQuantity, err := strconv.Atoi(quantity)
		if err != nil {
			panic(err.Error())
		}

		models.CreateNewProduct(name, description, convertedPrice, convertedQuantity)
	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	productdId := r.URL.Query().Get("id")
	models.DeleteProduct(productdId)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")
	product := models.EditProduct(productId)
	temp.ExecuteTemplate(w, "Edit", product)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		intId, err := strconv.Atoi(id)
		if err != nil {
			panic(err.Error())
		}

		floatPrice, err := strconv.ParseFloat(price, 64)
		if err != nil {
			panic(err.Error())
		}

		intQuantity, err := strconv.Atoi(quantity)
		if err != nil {
			panic(err.Error())
		}

		models.UpdateProduct(intId, name, description, floatPrice, intQuantity)
	}
	http.Redirect(w, r, "/", 301)
}
