package main
import (
	"time"
	/*"math/rand"*/
	"fmt"
)
var start = 0
func main() {
	c := make(chan int)
	for i := 0; i < 5; i++ {
		worker := &Worker{id: i}
		go worker.process(c)
	}
	for {
		start = start + 1
		c <- start//rand.Int()
		time.Sleep(time.Millisecond * 500)
	}
}
type Worker struct {
	id int
}
func (w *Worker) process(c chan int) {
	for {
		test := 66
		data := <-c
		fmt.Printf("worker %d got %d\n", w.id, data)
	}
}