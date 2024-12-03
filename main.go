package main

import "fmt"

func intFn() {
	var v int = 100

	var pt *int = &v

	var ppt **int = &pt

	fmt.Println("v", v)
	fmt.Println("pt", pt, *pt)
	fmt.Println("ppt", ppt, **ppt)

	v = 200
	fmt.Println("->", **ppt)
	fmt.Println("->", **ppt)
}

func arrayFn() {
	arr := []int{1, 2, 3}
	arrPt := &arr
	arrPpt := &arrPt

	fmt.Println("-> arr", arr)
	fmt.Println("-> arrPt", *arrPt)
	fmt.Println("-> arrPpt", **arrPpt)

	*arrPpt = &[]int{3, 4, 5}
	fmt.Println("---> arrPpt", **arrPpt)

}

func main() {
	arrayFn()
}
