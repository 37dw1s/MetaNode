package main

import "fmt"

type Custom struct {
	i int
}

var carr [5]*Custom = [5]*Custom{
	{6},
	{7},
	{8},
	{9},
	{10},
}

func main() {
	a := [5]int{5, 4, 3, 2, 1}
	fmt.Println("before all, a = ", a)
	for i := range carr {
		fmt.Printf("in main func, carr[%d] = %p, value = %v \n", i, &carr[i], *carr[i])
	}
	printFuncParamPointer(carr)

	receiveArray(a)
	fmt.Println("after receiveArray, a = ", a)

	fmt.Println(a)

	receiveArrayPointer(&a)
	fmt.Println("after receiveArrayPointer, a = ", a)
}

func receiveArray(param [5]int) {
	fmt.Println("in receiveArray func, before modify, param = ", param)
	param[1] = -5
	v := param
	fmt.Println(v)
	fmt.Println("in receiveArray func, after modify, param = ", param)
}

func receiveArrayPointer(param *[5]int) {
	fmt.Println("in receiveArrayPointer func, before modify, param = ", param)
	param[1] = -5
	fmt.Println("in receiveArrayPointer func, after modify, param = ", param)
}

func printFuncParamPointer(param [5]*Custom) {
	for i := range param {
		fmt.Printf("in printFuncParamPointer func, param[%d] = %p, value = %v \n", i, &param[i], *param[i])
	}
}
