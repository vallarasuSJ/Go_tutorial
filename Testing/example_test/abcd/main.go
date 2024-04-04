//Package abcd asked if you are ready to rock
package abcd

//Sum adds an unlimited number of values of type int
func Sum(n ...int) int {
	s := 0
	for _, v := range n {
		s += v
	}
	return s
}
