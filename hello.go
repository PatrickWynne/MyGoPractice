package main
 
import (
	"fmt"
	"os"	
	"time"
	"sync"
)

var (
	counter = 0
	lock sync.Mutex
)
 
func main() {
	//check the length of args
	if len(os.Args) > 2 {
		os.Exit(1)
	}

	runProcess()
	go func() {
		fmt.Println("processing 2")
	}()

	for i := 0; i < 20; i++ {
		go incr()
	}
	time.Sleep(time.Millisecond * 10)
		
/*
	//decalare 3 variables in 1 line
	name, power, test := "Goku", getPower(), "passed"
	//display the results
	fmt.Println("It's over", string(os.Args[1]), name, test, power)

	param1, exists := powers("goku")
	if exists == false {
		// handle this error case
		fmt.Println(exists)
		fmt.Println(param1)
	}
	gokus := SetupGoku()
	goku3 := &Saiyan{"Goku", 9000}
	Super(gokus)
	fmt.Println(goku3.Power)
	
	fmt.Println(gokus.Name)
	taco := Saiyan{Name: "taco man"}
	fmt.Println(taco.Name)
	fmt.Println(gokus.Name)
	vegita := NewSaiyan("vegita", 10000)
	fmt.Println(vegita.Name)
	*/

}

func getPower() int{
	return 9000
}
func Super(s Saiyan) Saiyan{
	s.Power += 10000
	return s
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
func NewSaiyan(name string, power int) *Saiyan {
	return &Saiyan{
		Name: name,
		Power: power,
	}
}
func runProcess(){
	fmt.Println("start")
	go process()
	time.Sleep(time.Millisecond * 10) // this is bad, don't do this!
	fmt.Println("done")
}

func process() {
	fmt.Println("processing")
}
func incr() {
	counter++
	fmt.Println(counter)
}