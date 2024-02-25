package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

type person struct {
	Name string
	Age  int
}

type user struct {
	First   string   `json:"First"`  
	Last    string   `json:"Last"`   
	Age     int64    `json:"Age"`    
	Sayings []string `json:"Sayings"`
}

type byAge []person 

type byName []person

//sort by age
func (a byAge) Len() int{
	return len(a)
}
func (a byAge) Less(i,j int) bool{
	return a[i].Age<a[j].Age
}

func (a byAge) Swap(i,j int) {
	a[i],a[j]=a[j],a[i]
}

//sort by name
func (a1 byName) Len() int{
	return len(a1)
}
func (a1 byName) Less(i,j int) bool{
	return a1[i].Name<a1[j].Name
}

func (a1 byName) Swap(i,j int) {
	a1[i],a1[j]=a1[j],a1[i]
}

func main() {

	//exercise 1 - marshal the data
	u1 := person{
		Name: "vallarasu",
		Age:  21,
	}

	u2 := person{
		"kv",
		22,
	}

	Users := []person{u1, u2}
	bs, err := json.Marshal(Users)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(bs)

	fmt.Println("-----------------------------")
	//exercise 2 - unmarshal 
	s1:=`[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`
	var people []user
	err=json.Unmarshal([]byte(s1),&people) 
	if err!=nil{
		fmt.Println(err)
	}
	// fmt.Println(people)

	for _,v:=range people{
		fmt.Printf("%s %s age is %d and his sayings:\n",v.First,v.Last,v.Age)
		
		for i,w:=range v.Sayings{
			fmt.Println(i," ",w)
		}
		fmt.Println("--------")
	}

	fmt.Println("-----------------------------")

	//exercise-3 encode and writer interface

	err=json.NewEncoder(os.Stdout).Encode(Users)
	if err!=nil{
		fmt.Println(err)
	}

	fmt.Println("-----------------------------")

	//exercise-4  sort
	xi := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
	xs := []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry", "fragmented", "moons", "across", "magenta"}

	sort.Ints(xi)
	sort.Strings(xs)
	fmt.Println(xi)
	fmt.Println(xs)

	fmt.Println("-----------------------------")

	//exercise-5 custom sort
	
	sort.Sort(byAge(Users)) 
	fmt.Println("Sort by age")
	for _,v:= range Users {
		fmt.Println(v.Age, " " ,v.Name)
	}

	fmt.Println("Sort by name")
	sort.Sort(byName(Users)) 

	for _,v:= range Users {
		fmt.Println(v.Age, " " ,v.Name)
	}









}