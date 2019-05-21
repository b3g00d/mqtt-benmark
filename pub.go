package main

import (
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"time"
)

type PubClient struct {
	Broker   string
	ClientID int
	Qos      int
}

var f MQTT.MessageHandler = func(client MQTT.Client, message MQTT.Message) {
	fmt.Printf("Published message on topic: %s\nMessage: %s\n", message.Topic(), message.Payload())

}

func (s *PubClient) Run() {
	topic := fmt.Sprintf("topic/test/%d", s.ClientID)
	client_id := fmt.Sprintf("client_pub_%d", s.ClientID)
	opts := MQTT.NewClientOptions().AddBroker(s.Broker)
	opts.SetClientID(client_id)
	opts.SetDefaultPublishHandler(f)
	c := MQTT.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())

	} else {
		fmt.Printf("Connected to %s with ID %s to Topic %s\n", s.Broker, client_id, topic)

	}
	for {
		for i := 0; i < 4; i++ {
			text := fmt.Sprintf("this is msg #%d!", i)
			token := c.Publish(topic, byte(s.Qos), false, text)
			token.Wait()
		}
		time.Sleep(1 * time.Second)
	}

}
