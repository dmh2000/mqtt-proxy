package mqttlib

import (
	MQTT "github.com/eclipse/paho.mqtt.golang"
)

type MqttTopic struct {
	Topic   string
	Message string
}

func Connect(broker string, id string, user string, password string) (MQTT.Client, error) {
	var token MQTT.Token
	opts := MQTT.NewClientOptions()
	opts.AddBroker(broker)
	opts.SetClientID(id)
	opts.SetUsername(user)
	opts.SetPassword(password)
	opts.SetCleanSession(false)

	// create new client
	client := MQTT.NewClient(opts)
	token = client.Connect()
	token.Wait()
	if token.Error() != nil {
		print(token.Error())
		return nil, token.Error()
	}

	return client, nil
}

func Disconnect(client MQTT.Client) {
	client.Disconnect(250)
}

func Publish(client MQTT.Client, topic string, message string) error {

	token := client.Publish(topic, 0, false, []byte(message))
	token.Wait()
	if token.Error() != nil {
		print(token.Error())
		client.Disconnect(250)
		return token.Error()
	}

	return nil
}

func Subscribe(client MQTT.Client, subscribe MqttTopic, callback MQTT.MessageHandler) error {

	token := client.Subscribe(subscribe.Topic, 0, callback)
	token.Wait()
	if token.Error() != nil {
		print(token.Error())
		client.Disconnect(250)
		return token.Error()
	}

	return nil
}
