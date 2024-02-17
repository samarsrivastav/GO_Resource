package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("New Data  Type called Slices 🤩")

	var fruitlist = []string{"apple", "banna", "jashdf"}

	fmt.Println(fruitlist)

	fruitlist = append(fruitlist, "hjh", "asdf")
	fmt.Println(fruitlist)

	// fruitlist = append(fruitlist[:3])
	// fmt.Println(fruitlist)

	// make syntax

	scores := make([]int, 4)

	scores[0] = 1
	scores[1] = 1
	scores[2] = 1

	scores = append(scores, 32, 22)
	fmt.Println(scores)
	fmt.Println(sort.IntsAreSorted(scores))
	sort.Ints(scores)
	fmt.Println(scores)
	fmt.Println(sort.IntsAreSorted(scores))

	// removing a value
	var courses = []string{"react", "js", "go", "c++"}
	courses = append(courses[:2], courses[3:]...) //deletes the 2nd index ;)
	fmt.Println(courses)
}
