package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

// Receive params:
// WAIT ms of wait between sequence of requests
// NUM number of services to test
// BATCH number of tests per batch
func main() {
	wait, err := strconv.Atoi(os.Getenv("WAIT"))
	top, err := strconv.Atoi(os.Getenv("BATCH"))
	num, err := strconv.Atoi(os.Getenv("NUM"))
	checkError("Params have to be a number", err)

	for {
		for i := 0; i < top; i++ {
			key := random(1, num)
			url := fmt.Sprintf("http://ms-%04d-srv:5000/srv%d", key, key)
			healthz := fmt.Sprintf("http://ms-%04d-srv:5000/_status/healthz", key)
			response, err := http.Get(url)
			responseHealthz, err := http.Get(healthz)
			checkError("Request failed", err)

			defer response.Body.Close()
			defer responseHealthz.Body.Close()
			_, err = io.Copy(os.Stdout, response.Body)
			checkError("Cannot read Response srv", err)
			_, err = io.Copy(os.Stdout, responseHealthz.Body)
			checkError("Cannot read Response healthz", err)
		}
		time.Sleep(time.Duration(wait) * time.Millisecond)
	}
}

func random(min, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(max-min) + min
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}
