package main

import (
	"io"
	"net/http"
	"sync"
	"time"
)

// send request to local server, return response body
func request(req string) (string, error) {
	resp, err := http.Get(req)
	if err != nil {
		// handle error
		return "", err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

// send message to broker
func publish(res string) error {
	// send to broker
	print(res + "\n")
	return nil
}

// wait for message from broker
func subscribe() (string, error) {

	// receive from broker
	time.Sleep(time.Second * 2)
	return "http://localhost:8000", nil
}

func subscriber(wg *sync.WaitGroup) {
	// subscribe to topic

	for {
		// receive from broker
		req, err := subscribe()
		if err != nil {
			print(err.Error())
			break
		}

		// send/receive to local server
		res, err := request(req)
		if err != nil {
			print(err.Error())
			break
		}

		// publish to broker
		err = publish(res)
		if err != nil {
			print(err.Error())
			break
		}

	}

	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	wg.Add(1)

	go subscriber(&wg)

	wg.Wait()
}
