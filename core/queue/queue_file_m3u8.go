package queue

import (
	"app/config"
	"app/constant"
	"encoding/json"
	"log"

	"github.com/rabbitmq/amqp091-go"
)

func QueueFileM3U8() {
	conn := config.GetRabbitmq()
	ch, err := conn.Channel()
	if err != nil {
		log.Println("error chanel: ", err)
		return
	}

	q, err := ch.QueueDeclare(
		string(constant.QUEUE_FILE_M3U8),
		true,
		false,
		false,
		false,
		amqp091.Table{},
	)
	if err != nil {
		log.Println("error queue declare: ", err)
		return
	}
	log.Printf("start %s", string(constant.QUEUE_FILE_M3U8))

	msgs, err := ch.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Println("error consumer: ", err)
		return
	}

	for d := range msgs {
		var payload QueueFileM3U8Payload
		err := json.Unmarshal(d.Body, &payload)
		if err != nil {
			log.Println("error msg: ", err)
			d.Reject(true)
			continue
		}

		log.Println(payload)
		d.Ack(false)
	}
}

type QueueFileM3U8Payload struct {
	Path     string `json:"path"`
	IpServer string `json:"ipServer"`
	Uuid     string `json:"uuid"`
	Quantity string `json:"quantity"`
}
