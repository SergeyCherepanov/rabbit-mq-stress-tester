package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/streadway/amqp"
	"github.com/codegangsta/cli"
)

func Consume(uri string, doneChan chan bool, c *cli.Context) {
	log.Println("Consuming...")
	connection, err := amqp.DialConfig(uri, amqp.Config{
		Heartbeat: 10 * time.Second,
		Locale:    "en_US",
		Vhost: c.String("vhost"),
	})
	if err != nil {
		log.Fatal(err.Error())
	}
	defer connection.Close()

	channel, err := connection.Channel()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer channel.Close()

	q := MakeQueue(channel, c)

	msgs, err := channel.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	for d := range msgs {
		doneChan <- true
		var thisMessage MqMessage
		err4 := json.Unmarshal(d.Body, &thisMessage)
		if err4 != nil {
			log.Printf("Error unmarshalling! %s", err.Error())
		}
		log.Printf("Message age: %s", time.Since(thisMessage.TimeNow))

	}

	log.Println("done receiving")

}
