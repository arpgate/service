package arpgate

// option https://github.com/yosssi/gmq

import (
	"fmt"
	MQTT "git.eclipse.org/gitroot/paho/org.eclipse.paho.mqtt.golang.git"
)

func ListenMqtt() {

	fmt.Println("Mqtt listener starts...")

	opts := MQTT.NewClientOptions()
	opts.AddBroker("tcp://mqtt.arpgate.com:1883")
	//opts.SetProtocolVersion(4)
	opts.SetClientID("someguid")

	opts.SetDefaultPublishHandler(func(client *MQTT.Client, msg MQTT.Message) {

		fmt.Println("Topic=", msg.Topic(), "Payload=", string(msg.Payload()))
	})

	opts.SetOnConnectHandler(func(onConn *MQTT.Client) {
		fmt.Println(" connected")

		onConn.Subscribe("hello/world", 0, func(client *MQTT.Client, msg MQTT.Message) {

			fmt.Println("Topic=", msg.Topic(), "Payload=", string(msg.Payload()))
		})
		fmt.Println("Mqtt listening...")
	})

	client := MQTT.NewClient(opts)

	client.Connect()
	// fmt.Println(token)

}
