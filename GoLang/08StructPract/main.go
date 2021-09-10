package main

import (
	"fmt"
	"strconv"
)

func main() {

	var prod1 Product = Product{
		id:          1,
		title:       "Nike Shoes",
		description: "Running shoes by Nike",
		price:       159.99,
	}

	var prod2 Product = *CreateProduct(
		2,
		"POLO Shirt",
		"Casual Polo T-Shirt",
		99,
	)

	prod3 := getProduct()

	// fmt.Println(newProduct)
	// fmt.Println(newProdByHelper)
	// productData(&newProdByHelper)
	prod1.outputData()
	prod2.outputData()
	prod3.outputData()
	prod3.saveToFile()
}

func getProduct() *Product {
	prodId := readUserInput("Enter Product Id: ")
	prodTitle := readUserInput("Enter Product Title: ")
	prodDesc := readUserInput("Enter Product Desc: ")
	prodPrice := readUserInput("Enter Product Price: ")

	id, _ := strconv.ParseInt(prodId, 0, 64)
	price, _ := strconv.ParseFloat(prodPrice, 64)

	var newProd Product = *CreateProduct(id, prodTitle, prodDesc, price)
	return &newProd

}

func readUserInput(promptText string) (data string) {
	fmt.Print(promptText)
	fmt.Scanln(&data)
	return
}
