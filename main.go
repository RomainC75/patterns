package main

import (
	"design-patterns/finder"
	"fmt"
	"log"
)

var min int64 = 1_000_000
var max int64 = 10_000_000

func main() {
	f, err := finder.NewFinder[int64](finder.Prime)
	if err != nil {
		log.Fatal(err.Error())
	}

	res := f.RunParallel(min, max)
	fmt.Println("res : ", res)

}
