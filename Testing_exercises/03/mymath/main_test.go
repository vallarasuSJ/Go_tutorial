package mymath

import (
	"fmt"
	"testing"
)

func TestCenterAvg(t *testing.T) {
	type test struct {
		data   []int
		answer float64
	}

	tests := []test{
		test{[]int{3, 5, 34, 34}, 19.5},
		test{[]int{5, 5, 45, 45, 33}, 27.666666},
	}

	for _, v := range tests {
		f := CenterAvg(v.data)
		if f != v.answer {
			t.Error("got ", f, " want ", v.answer)
		}
	}

}

func ExampleCenterAvg() {
	fmt.Println([]int{1, 2, 4, 6, 9})
	//Output:
	//4
}

func BenchmarkCenterAvg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CenterAvg([]int{4, 5534, 54, 5, 43})
	}
}
