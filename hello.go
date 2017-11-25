package main
import (
	//"bytes"
    "encoding/json"
    "fmt"
    "io/ioutil"
	"net/http"
	"strings"
)
var start = 0
func main() {
	fmt.Println("Starting the application...")
    response, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
    if err != nil {
        fmt.Printf("The HTTP request failed with error %s\n", err)
    } else {
        data, _ := ioutil.ReadAll(response.Body)
		result := string(data)
		dec := json.NewDecoder(strings.NewReader(result))
		
		type Message struct {
			UserId int
			Id int
			Title string
			Body string
		}
		var m Message
		dec.Decode(&m)
		fmt.Printf("%s: %d\n", m.Title, m.Id)
		//dec := json.NewDecoder(strings.NewReader(data))
	}
	

/*
	fmt.Printf("posting to rest api")
    jsonData := map[string]string{"firstname": "Nic", "lastname": "Raboy"}
    jsonValue, _ := json.Marshal(jsonData)
    response, err = http.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(jsonValue))
    if err != nil {
        fmt.Printf("The HTTP request failed with error %s\n", err)
    } else {
        data, _ := ioutil.ReadAll(response.Body)
        fmt.Println(string(data))
    }
	fmt.Println("Terminating the application...")*/
}