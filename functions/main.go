package main

import (
	// "bytes"
	"fmt"
	"io"
	"log"
	"os"

	// "os"
	"strconv"
)

type book struct {
	title string
}

//writeOut method
func (b book) writeOut(w io.Writer) error {
	_, err := w.Write([]byte(b.title))
	return err
}

func (b book) String() string { //method which accepts person type and a stringer interface
	return fmt.Sprint("The title ", b.title)
}

type count int

func (c count) String() string {
	return fmt.Sprint("The count is ", strconv.Itoa(int(c))) //formats strings and return
}

//wrapper function
func logInfo(s fmt.Stringer) {
	log.Println("Log is ", s.String())
}

//function returning a function
func bar() func() int {
	return func() int {
		return 43
	}
}

//readFile
func readFile(fileName string) ([]byte, error) {
	xb, err := os.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("Error %s ", err) //function that wraps error function
	}
	return xb, nil
}

func main() {
	b := book{
		title: "harry potter",
	}
	n := 43

	fmt.Println(b)
	fmt.Println(n)
	logInfo(b)

	//create a file
	// f,err:=os.Create("Output.txt")
	// if err!=nil{
	// 	log.Fatalf("error %s",err)
	// }
	// defer f.Close()

	// s:=[]byte("hello kv!! Long time no see")

	// _,err=f.Write(s)

	// if err!=nil{
	// 	log.Fatalf("error %s",err)
	// }

	// //write in a file
	// b1:=book{
	// 	title: "naruto",
	// }

	// b1.writeOut(f)

	// //byter buffer writer
	// by:=bytes.NewBufferString("hello")
	// fmt.Println(by)
	// by.WriteString("Gophers")
	// fmt.Println(by)

	//function returning a function
	y := bar()
	fmt.Println(y())

	//wrapper function
	xb, err := readFile("output.txt")
	if err != nil {
		log.Fatalf("Error  %s ", err)
	}
	fmt.Println(xb)
	fmt.Println(string(xb))

}
