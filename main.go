package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

var seats = 20
var mutex = &sync.Mutex{}

func book(wg *sync.WaitGroup) {
	defer wg.Done()
	mutex.Lock()
	if seats > 0 {
		seats--
		fmt.Println("Booked a seat")
	} else {
		fmt.Println("No seats available")
	}
	mutex.Unlock()
}

func bookSeat(w http.ResponseWriter, r *http.Request) {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go book(wg)
	wg.Wait()

	response := struct {
		Message string `json:"message"`
	}{
		Message: "Seat booked successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/book", bookSeat)
	fmt.Printf("Starting server at port 8080 \n")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
