package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	url := "https://jsonplaceholder.typicode.com/users"
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("get eror %d", err)
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(response.Body)
		if err != nil {
			println("error.")
		}
		data := string(bodyBytes)
		println(data)
	}

}

// type Person struct {
// 	Name    string `json:"name"`
// 	Age     int    `json:"age,omitempty"`
// 	IsAdult string `json:"-"`
// }

// // func main() {
// // 	var person Person
// // 	rawMessage := json.RawMessage(`{
// // 		"name": "mike",
// // 		"age": 25
// // 	}`)
// // 	fmt.Println(rawMessage)
// // 	err := json.Unmarshal(rawMessage, &person)
// // 	if err != nil {
// // 		println("error in unmarshall")
// // 		fmt.Println(err)
// // 		return
// // 	}
// // 	fmt.Printf("%+v\n", person)
// // }

// Channels
// func reader(id int, ch <-chan int, wg *sync.WaitGroup) {
// 	defer wg.Done()

// 	for {
// 		val, ok := <-ch
// 		if !ok {
// 			fmt.Println("Channel Closed")
// 			return
// 		}
// 		fmt.Printf("Reader %d Recieved %d\n", id, val)
// 	}

// }

// func main() {

// 	var wg sync.WaitGroup
// 	ch := make(chan int)

// 	wg.Add(4)

// 	go reader(1, ch, &wg)
// 	go reader(2, ch, &wg)
// 	go reader(3, ch, &wg)
// 	go reader(4, ch, &wg)

// 	for i := 0; i < 100; i++ {
// 		ch <- i
// 	}

// 	close(ch)

// 	wg.Wait()
// 	println("main...")
// }

// func taskA() {
// 	println("start task A ...")
// 	time.Sleep(3 * time.Second)
// 	println("end task A ...")
// }

// func taskB() {
// 	println("start task B ...")
// 	time.Sleep(3 * time.Second)
// 	println("end task B ...")
// }
// func main() {
// 	println(" start main...")
// 	go taskA()
// 	go taskB()
// 	time.Sleep(10 * time.Second)
// 	println(" end main...")
// }
