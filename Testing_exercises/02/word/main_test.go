package word

import (
	"fmt"
	"testing"
	"tutorial/Testing_exercises/02/quote"
)

func TestCount(t *testing.T) {
	n := Count(quote.SunAlso)
	if n != 322 {
		t.Error("Want ", 322, "got ", n)
	}
}

func TestUseCount(t *testing.T) {
	m := UseCount("one two three")
	for k, v := range m {
		switch k {
		case "one":
			if v != 1 {
				t.Error("got", v, "want", 1)
			}
		case "two":
			if v != 1 {
				t.Error("got", v, "want", 2)
			}

		case "three":
			if v != 1 {
				t.Error("got", v, "want", 3)
			}

		}
	}
}

func ExampleCount() {
	fmt.Println(Count(quote.SunAlso))
	//Output:
	//322
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(quote.SunAlso)
	}
}

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount(quote.SunAlso)
	}
}
