package main

import "fmt"

func main() {
	var x = "hello "
	x += "world"
	fmt.Println(x)
	assisgn()
	constants()
	learnLoop()
}

func assisgn() {
	y := 5
	fmt.Println(y)
}

func constants() {
	const c = 5
	fmt.Println(c)
}

func doubleInput() {
	fmt.Print("Enter number")
	var input float64
	fmt.Scanf("%f", &input)
	output := 2 * input
	fmt.Println(output)
}

func learnLoop() {
	for i := 1; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println("even")
		} else {
			fmt.Println("odd")
		}
	}
	learnArray()
}

func learnArray() {
	var x [5]int
	x[1] = 33
	fmt.Println(x)
	learnMap()
}

func learnMap() {
	x := make(map[string]string)
	x["h"] = "hello"
	x["r"] = "richa"
	if name, ok := x["hi"]; ok {
		fmt.Println(name, ok)
	}

}
