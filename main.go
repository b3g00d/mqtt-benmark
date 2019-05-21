package main

import (
	"fmt"
	//import the Paho Go MQTT library
	"flag"
	"strconv"
	"strings"
	"sync"
)

//define a function for the default message handler
func main() {
	//create a ClientOptions struct setting the broker address, clientid, turn
	//off trace output and set the default message handler
	var (
		broker          = flag.String("broker", "tcp://localhost:1883", "MQTT broker endpoint as scheme://host:port")
		client_type     = flag.String("client_type", "sub", "Client Type sub or pub")
		client_id_range = flag.String("client_id_range", "0,10000", "client range")
		qos      = flag.Int("qos", 1, "QoS for published messages")
	)
	flag.Parse()
	fmt.Println(*broker)
	fmt.Println(*client_type)
	fmt.Println(*client_id_range)
	s := strings.Split(*client_id_range, ",")
	from, _ := strconv.Atoi(s[0])
	to, _ := strconv.Atoi(s[1])

	var waitgroup sync.WaitGroup
	if *client_type == "sub" {
		for i := from; i < to; i++ {
			waitgroup.Add(1)
			c := &SubClient{
				ClientID: i,
				Broker:   *broker,
				Qos: *qos,
			}
			go c.Run()
		}

	}
	waitgroup.Wait()
}
