package main

import "fmt" 

var  (
	a,b,c *string               //package level scope
	d *int
)

func init(){
	//exercise-2 (dereference an address) 
	p:="Drop by drop, the bucket"
	q:="Persistently , patiently"                         //function level scope
	r:="the meaning of ...."
	n:=32
	a=&p
	b=&q
	c=&r
	d=&n
} 

type dog struct{
	name string
}

func (d dog) walk(){
	fmt.Println("My name is ",d.name, "and I'm walking")
}

func (d *dog) run(){
	fmt.Println("My name is ",d.name, "and I'm running")
} 

type youngin interface{
	walk()
	run()
} 

func youngRun(y youngin){
	y.walk()
	y.run()
} 

func (d *dog) changeName(s string) *dog{
	d.name=s
	return d
} 

type person struct{
	name string
}

func valueSemantics(p person,s string) person{
	p.name=s
	return p

}

func pointerSemantics(p *person, s string){
	p.name=s
} 

func main() {

	//exercise-1 
	x:=42 
	fmt.Printf("%v, %v\n",x,&x) 
	y:=&x
	fmt.Printf("%v, %v\n",y,*y)
	fmt.Println(*&x) 
	*y=45
	fmt.Println(*y) 
	fmt.Println("----------------------------------------")

	//exercise-2   

	fmt.Printf("%v\t%T\n",a,a)
	fmt.Printf("%v\t%T\n",b,b)
	fmt.Printf("%v\t%T\n",c,c)
	fmt.Printf("%v\t%T\n",d,d)

	fmt.Println(*a)
	fmt.Println(*b)
	fmt.Println(*c)
	fmt.Println(*d)
	fmt.Println("----------------------------------------")

	//exercise-3
	d1:=dog{
		name: "Puppy",
	}
	d1.walk()
	d1.run()


	d2:=&dog{
		name:"Tiger",
	}
	youngRun(d2)

	d2=d2.changeName("Rover")
	youngRun(d2)

	fmt.Println("----------------------------------------")

	//exercise-4  
	p:=person{"shinchan"}
	fmt.Println(p)
	p=valueSemantics(p,"Naruto")
	fmt.Println(p)

	pointerSemantics(&p,"Luffy")
	fmt.Println(p)







}