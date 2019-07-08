package main

import (
	"encoding/json"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
)

/*
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}*/

func ConsumeStudents(q1 *gin.Context) {
	var p Student
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	log.Println("RabbitMQ Consumer...")
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

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			//p1 := json.Unmarshal(p, &d.Body)
			json.Unmarshal(d.Body, &p)
			
			//log.Printf("Received a message: %s", d.Body)
			log.Println("RabbitMQ received message", string(d.Body))
			db_elastic_CreateStudent(p, q1)
		}
	}()
	
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
