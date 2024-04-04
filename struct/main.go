package main

import (
	"fmt"
)

type person struct {
	name        string
	age         int
	favIcecream []string
}
type secret struct {
	person
	flag bool
}

type engine struct {
	electric bool
}

type vehicle struct {
	engine
	make  string
	model string
	doors int
	color string
}

func main() {
	p1 := person{
		name:        "kv",
		age:         21,
		favIcecream: []string{"cone", "chocolate"},
	}
	fmt.Println(p1, " ", p1.name, " ", p1.age)

	p2 := secret{
		person: person{
			name:        "kv",
			age:         34,
			favIcecream: []string{"choc", "vanilla"},
		},
		flag: true,
	}
	fmt.Printf("%#v\n", p2)

	for _, v := range p2.favIcecream {
		fmt.Println(v)
	}

	//map struct
	m := map[string]person{
		p1.name: p1,
	}

	fmt.Println(m)

	fmt.Println("--------------------------")

	v1 := vehicle{
		engine: engine{electric: true},
		make:   "ford",
		model:  "mustang",
		doors:  4,
		color:  "black",
	}

	fmt.Println(v1)
	fmt.Println(v1.electric)

	fmt.Println("--------------------------")

	a1 := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first:   "kv",
		friends: map[string]int{"naruto": 43, "luffy": 32},
		favDrinks: []string{
			"fanta",
			"pepsi",
		},
	}

	for k, v := range a1.friends {
		fmt.Println(k, " ", v)
	}
	for _, v := range a1.favDrinks {
		fmt.Println(v)
	}

}
