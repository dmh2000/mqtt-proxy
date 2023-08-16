package main

type MqttParam struct {
	Broker   string // tls://localhost:8883
	Id       string // client id (every client needs a unique one)
	User     string // user Id at broker
	Password string // password at broker
}

type MqttTopic struct {
	Topic   string
	Message string
}
