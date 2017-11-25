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
	param1, exists := powers("goku")
	if exists == false {
		// handle this error case
		fmt.Println(exists)
		fmt.Println(param1)
	}
	gokus := SetupGoku()
	fmt.Println(gokus.Name)
	taco := Saiyan{Name: "taco man"}
	fmt.Println(taco.Name)

}

func getPower() string{
	return "this many"
}
func SetupGoku() Saiyan{
	goku := Saiyan{
		Name: "Goku",
		Power: 9000,
	}
	return goku
}
func add(a int, b int) int {
	return 10
}
func powers(name string) (int, bool) {
	return 10,false
}
type Saiyan struct {
	Name string
	Power int
}