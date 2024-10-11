package logic

import (
	"encoding/json"
	"log"
	"microservice/entity"

	"github.com/nats-io/nats.go"
)

func (ps *PaymentService) SubscribeToUserCreated(nc *nats.Conn) {
	nc.Subscribe("user.created", func(msg *nats.Msg) {
		var newUser entity.User
		if err := json.Unmarshal(msg.Data, &newUser); err != nil {
			log.Printf("Error unmarshalling user data: %v", err)
			return
		}

		// Store the new user in the Payment service's database
		if err := ps.db.Create(&newUser).Error; err != nil {
			log.Printf("Error saving user to database: %v", err)
		} else {
			log.Printf("User %s created in Payment service database.", newUser.Name)
		}
	})
}
