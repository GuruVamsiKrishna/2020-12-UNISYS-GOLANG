package main

import "fmt"

func main() {

	// var scores = map[string]int{"james": 79}
	scores := make(map[string]int)

	scores["john"] = 92
	scores["miller"] = 87
	scores["jacob"] = 95
	scores["mike"] = 0

	name := "john"
	s1, ok := scores[name]
	if ok {
		fmt.Printf("%s's score is %d\n", name, s1)
	} else {
		fmt.Println("No one with name -", name)
	}

	fmt.Println("before delete:", scores)
	delete(scores, "mike")
	fmt.Println("after delete:", scores)

}
