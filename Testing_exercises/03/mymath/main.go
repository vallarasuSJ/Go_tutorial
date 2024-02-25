//Package mymath provides math solutions relevant to out company
package mymath

import "sort"


//CenterAverage computes the average of list of numbers  
//after removing largest and smallest values
func CenterAvg(x []int) float64 {
	sort.Ints(x) 
	x=x[1:(len(x)-1)]  

	n:=0 

	for _,v:=range x{
		n+=v
	}

	f:=float64(n)/float64(len(x))
	return f

	
	

}