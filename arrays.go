package main
import ("fmt")

func main(){

	//arrays with defined length
	var arr1 = [3]int{1,2,3}
	arr2 := [5]int{3,4,5,6,7}
	fmt.Println(arr1)
	fmt.Println(arr2)

	//arrays with undefined length
	var arr3 = [...]int{1,2,3}
	arr4 := [...]int{4,5,6,7,8}
	fmt.Println(arr3)
	fmt.Println(arr4)

	//change elemetns of an array
	prices :=[3]int{10,20,30}
	prices[2]=50
	fmt.Println(prices)

	//array initialization
	arr5 :=[5]int{}//not initialized
	arr6 :=[5]int{1,2}//partially initialized
	arr7 :=[5]int{1,2,3,4,5}//fully initialized
	arr8 :=[5]int{1:10,2:40}
	fmt.Println(arr5)
	fmt.Println(arr6)
	fmt.Println(arr7)
	fmt.Println(arr8)

}