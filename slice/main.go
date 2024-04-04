package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4}
	fmt.Println(a)

	str := []string{"Almond Biscotti Café", "Banana Pudding ", "Balsamic Strawberry (GF)",
		"Bittersweet Chocolate (GF)", "Blueberry Pancake (GF)", "Bourbon Turtle (GF)",
		"Browned Butter Cookie Dough", "Chocolate Covered Black Cherry (GF)", "Chocolate Fudge Brownie",
		"Chocolate Peanut Butter (GF)", "Coffee (GF)", "Cookies & Cream", "Fresh Basil (GF)",
		"Garden Mint Chip (GF)", "Lavender Lemon Honey (GF)", "Lemon Bar",
		"Madagascar Vanilla (GF)", "Matcha (GF)", "Midnight Chocolate Sorbet (GF, V)",
		"Peanut Butter (GF, V)", "Non-Dairy Coconut Matcha (GF, V)",
		"Non-Dairy Sunbutter Cinnamon (GF, V)", "Orange Cream (GF) ", "Peanut Butter Cookie Dough",
		"Raspberry Sorbet (GF, V)", "Salty Caramel (GF)", "Slate Slate Different",
		"Strawberry Lemonade Sorbet (GF, V)", "Vanilla Caramel Blondie", "Vietnamese Cinnamon (GF)",
		"Wolverine Tracks (GF)"}
	//append value in slice
	str = append(str, "ice cream")

	//slicing  in slice
	fmt.Println(str[5:8])

	//deleting a slice
	str = append(str[:4], str[5:]...)

	fmt.Println("------------------")

	for _, v := range str {
		fmt.Println(v)
	}

	fmt.Println("------------------")

	//make function
	d := make([]int, 0, 10)
	fmt.Println(len(d))
	fmt.Println(cap(d))
	d = append(d, 3, 2, 2, 3, 32)
	fmt.Println(d)
	fmt.Println(len(d))
	fmt.Println(cap(d))

	fmt.Println("------------------")

	//multidimensional slice
	a1 := []int{1, 2, 3}
	a2 := []int{6, 4, 5}
	a3 := [][]int{a1, a2}
	fmt.Println(a3)

	fmt.Println("------------------")

	//underlying array
	b1 := []int{4, 5, 6, 7}
	b2 := make([]int, 4)
	// copy(b2,b1)
	b1[3] = 9
	fmt.Println(b1)
	fmt.Println(b2)

}
