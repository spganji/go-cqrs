package main

import (
	"encoding/json"
	"log"
    //"fmt"
	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		//fmt.Println(err)
		log.Fatalf("%s: %s", msg, err)
	}
}

func rabbitmqSend(p Student, q1 *gin.Context) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	log.Println("RabbitMQ Producer...")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"StudentQueue", // name
		false,            // durable
		false,            // delete when unused
		false,            // exclusive
		false,            // no-wait
		nil,              // arguments
	)
	failOnError(err, "Failed to declare a queue")
	log.Println("RabbitMQ StudentQueue...")

	studentJson, err := json.Marshal(p)
	body := studentJson
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			// ContentType: "text/plain",
			ContentType: "application/json",
			Body:        []byte(body),
		})
	//log.Printf(" [x] Sent %s", body)
	failOnError(err, "Failed to publish a message")
	log.Println("RabbitMQ sending message:", string(body))
}
