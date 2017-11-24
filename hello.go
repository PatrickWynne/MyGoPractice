package main
 
import (
	"fmt"
	"os"	
)
 
func main() {
	name, power, test := "Goku", getPower(), "passed"
	fmt.Println(len(os.Args))
	if len(os.Args) > 2 {
		os.Exit(1)
	}
	os.Args[1] = string(power)
	fmt.Println("It's over", string(os.Args[1]), name, test)
		
	fmt.Println("Hello World!")
	_, exists := powers("goku")
	if exists == false {
		// handle this error case
		fmt.Println("false")
	}
}
func getPower() string{
	return "this many"
}
func log(message string) {
	
}
func add(a int, b int) int {
	return 10
}
func powers(name string) (int, bool) {
	return 10,false
}