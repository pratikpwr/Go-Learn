package main

import (
	"fmt"
	"os"
)

type Product struct {
	id          int64
	title       string
	description string
	price       float64
}

func CreateProduct(
	id int64,
	title string,
	description string,
	price float64,
) (product *Product) {

	product = &Product{
		id:          id,
		title:       title,
		description: description,
		price:       price,
	}
	return
}

func (product *Product) outputData() {
	fmt.Println("----- Product Details -----")
	fmt.Printf(
		"Id: %v\nTitle: %v\nDescription: %v\nPrice: %v\n",
		product.id,
		product.title,
		product.description,
		product.price,
	)
}

func (product *Product) saveToFile() {
	// create a file and pass filename as parameter
	file, _ := os.Create(product.title + ".txt")

	// string to be written in file
	content := fmt.Sprintf(
		"Id: %v\nTitle: %v\nDescription: %v\nPrice: %v\n",
		product.id,
		product.title,
		product.description,
		product.price,
	)

	// writes the string in file
	file.WriteString(content)

	// close after job is done
	file.Close()
}
