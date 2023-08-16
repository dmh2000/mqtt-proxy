package main

import (
	"os"
	"sync"
	"time"
)

func main() {
	var publisher = MqttParam{os.Getenv("MQTT_BROKER"), "pub", "xyzzy", "Execute1"}
	var publish = MqttTopic{"test", "Hello World"}
	var subscriber = MqttParam{os.Getenv("MQTT_BROKER"), "sub", "xyzzy", "Execute1"}
	var subscribe = MqttTopic{"test", ""}

	var wg sync.WaitGroup
	wg.Add(2)
	go Publish(publisher, publish, &wg)
	time.Sleep(5 * time.Second)
	go Subscribe(subscriber, subscribe, &wg)
	wg.Wait()
	print("done")
}
