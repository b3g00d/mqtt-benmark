package main

import (
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
)

type SubClient struct {
	Broker   string
	ClientID int
	Qos      int
}

func onMessageReceived(client MQTT.Client, message MQTT.Message) {
	fmt.Printf("Received message on topic: %s\nMessage: %s\n", message.Topic(), message.Payload())

}

func (s *SubClient) Run() {
	topic := fmt.Sprintf("topic/test/%d", s.ClientID)
	client_id := fmt.Sprintf("client_sub_%d", s.ClientID)
	opts := MQTT.NewClientOptions().AddBroker(s.Broker)
	opts.SetClientID(client_id)
	opts.OnConnect = func(c MQTT.Client) {
		if token := c.Subscribe(topic, byte(s.Qos), onMessageReceived); token.Wait() && token.Error() != nil {
			panic(token.Error())

		}

	}
	c := MQTT.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())

	} else {
		fmt.Printf("Connected to %s with ID %s to Topic %s\n", s.Broker, client_id, topic)

	}

}
