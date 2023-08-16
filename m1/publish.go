package main

import (
	"fmt"
	"sync"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

func Publish(publisher MqttParam, publish MqttTopic, wg *sync.WaitGroup) {

	opts := MQTT.NewClientOptions()
	opts.AddBroker(publisher.Broker)
	opts.SetClientID(publisher.Id)
	opts.SetUsername(publisher.User)
	opts.SetPassword(publisher.Password)
	opts.SetCleanSession(false)

	client := MQTT.NewClient(opts)
	token := client.Connect()
	if token.Wait() && token.Error() != nil {
		print(token.Error())
		panic(token.Error())
	}
	fmt.Println("Sample Publisher Started")
	for i := 0; i < 25; i++ {
		fmt.Printf("PUBLISHED TOPIC: %s MESSAGE: %s\n", publish.Topic, publish.Message)
		token := client.Publish(publish.Topic, 0, false, []byte(publish.Message))
		token.Wait()
		time.Sleep(1 * time.Second)
	}

	client.Disconnect(250)
	fmt.Println("Sample Publisher Disconnected")
	wg.Done()
}
