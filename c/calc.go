package c

import "github.com/MortalHappiness/testgomodule/a"

// Double returns x*2.
func Double(x int) int {
	return x * 2
}

func HelloFromC() string {
	aHello := a.Hello()
	return "Hello from C: " + aHello
}
