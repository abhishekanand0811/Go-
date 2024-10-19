package main
import("fmt")

func main(){
	myslices1 := []int{}
	myslices2 := [5]int{1,2,3,4,5}
	myslices3 := []string{"Go", "String", "Are", "Powerful"}
	fmt.Println(myslices1)
	fmt.Println(myslices2)
	fmt.Println(len(myslices1))
	fmt.Println(cap(myslices1))
	fmt.Println(len(myslices3))
	fmt.Println(cap(myslices3))




}