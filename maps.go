package main

import "fmt"

func main() {
	grades := make(map[string]float32)

	grades["Nick"] = 93
	grades["Tboy"] = 97

	nicksGrade := grades["Nick"]
	fmt.Println("Nick:", nicksGrade)

	delete(grades, "Nick")

	for k, v := range grades {
		fmt.Println(k, v)
	}
}
