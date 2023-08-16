package main

import (
	"fmt"
	"os"
	"sync"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

func Subscribe(subscriber MqttParam, subscribe MqttTopic, wg *sync.WaitGroup) {
	opts := MQTT.NewClientOptions()
	opts.AddBroker(subscriber.Broker)
	opts.SetClientID(subscriber.Id)
	opts.SetUsername(subscriber.User)
	opts.SetPassword(subscriber.Password)
	opts.SetCleanSession(false)

	receiveCount := 0
	choke := make(chan [2]string)

	opts.SetDefaultPublishHandler(func(client MQTT.Client, msg MQTT.Message) {
		choke <- [2]string{msg.Topic(), string(msg.Payload())}
	})

	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	token := client.Subscribe(subscribe.Topic, 0, nil)
	token.Wait()
	if token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	for receiveCount < 25 {
		select {
		case incoming := <-choke:
			topic := incoming[0]
			payload := string(incoming[1])
			fmt.Printf("RECEIVED TOPIC: %s MESSAGE: %s\n", topic, payload)
		case <-time.After(5 * time.Second):
			client.Disconnect(250)
			receiveCount++
		}
		if !client.IsConnected() {
			break
		}
	}

	fmt.Println("Sample Subscriber Disconnected")
	wg.Done()
}
