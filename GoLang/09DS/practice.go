package main

import "fmt"

type Product struct {
	id    int
	title string
	price int
}

func practice() {

	// 1
	var hobbies [3]string = [3]string{"Gaming", "Anime", "Coding"}
	fmt.Println(hobbies)

	// 2
	h1 := hobbies[0]
	fmt.Println(h1)
	h2 := hobbies[1:3]
	fmt.Println(h2)

	// 3
	h3 := hobbies[0:2]
	fmt.Println(h3)

	h4 := h1[:2]
	fmt.Println(h4)

	// 4
	h5 := h4[1:3]
	fmt.Println(h5)

	// 5
	var goals []string = []string{"GoBasics", "GoApi"}
	fmt.Println(goals)

	// 6
	goals[1] = "GoBackend"
	goals = append(goals, "GoJob")
	fmt.Println(goals)

	// 7
	var p1 Product = Product{1, "Shoes", 1299}
	var p2 Product = Product{2, "T-Shirt", 899}

	var products []Product = []Product{p1, p2}
	fmt.Println(products)

	p3 := Product{3, "Jeans", 2500}
	products = append(products, p3)

	fmt.Println(products)

}
