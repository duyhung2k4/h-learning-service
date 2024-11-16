package queue

import (
	"app/config"
	"app/constant"
	"encoding/json"
	"log"

	"github.com/rabbitmq/amqp091-go"
)

func QueueMp4Quantity() {
	conn := config.GetRabbitmq()
	ch, err := conn.Channel()
	if err != nil {
		log.Println("error chanel: ", err)
		return
	}

	q, err := ch.QueueDeclare(
		string(constant.QUEUE_MP4_360_P),
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
	log.Printf("start %s", string(constant.QUEUE_MP4_360_P))

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
		var payload QueueMp4QuantityPayload
		err := json.Unmarshal(d.Body, &payload)
		if err != nil {
			log.Println("error msg: ", err)
			d.Reject(false)
			continue
		}

		log.Println(payload)
		d.Ack(false)
	}
}

type QueueMp4QuantityPayload struct {
	Path     string `json:"path"`
	IpServer string `json:"ipServer"`
}
