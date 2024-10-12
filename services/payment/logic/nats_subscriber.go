package logic

import (
	"encoding/json"
	"log"
	"microservice/entity"

	"github.com/nats-io/nats.go"
)

func (ps *PaymentService) SubscribeToUserCreated() error {
	// Create a durable subscription to JetStream
	subscription, err := ps.js.Subscribe("user.created", func(msg *nats.Msg) {
		var newUser entity.User
		if err := json.Unmarshal(msg.Data, &newUser); err != nil {
			log.Printf("Error unmarshalling user data: %v", err)
			msg.Nak() // Acknowledge that the message could not be processed
			return
		}

		// Store the new user in the Account service's database
		if err := ps.db.Create(&newUser).Error; err != nil {
			log.Printf("Error saving user to database: %v", err)
			msg.Nak() // Acknowledge that the message could not be processed
		} else {
			log.Printf("User %s created in Payment service database.", newUser.Name)
			msg.Ack() // Acknowledge successful processing
		}
	}, nats.Durable("payment_service_durable"))

	if err != nil {
		return err
	}

	log.Printf("Subscribed to user.created events with durable subscription: %s", subscription.Subject)
	return nil
}

// SubscribeToAccountCreated subscribes to account.created events using JetStream
func (ps *PaymentService) SubscribeToAccountCreated() error {
	// Create a durable subscription to JetStream
	subscription, err := ps.js.Subscribe("account.created", func(msg *nats.Msg) {
		var newAccount entity.Account
		if err := json.Unmarshal(msg.Data, &newAccount); err != nil {
			log.Printf("Error unmarshalling account data: %v", err)
			msg.Nak() // Acknowledge that the message could not be processed
			return
		}

		// Store the new account in the Payment service's database
		if err := ps.db.Create(&newAccount).Error; err != nil {
			log.Printf("Error saving account to database: %v", err)
			msg.Nak() // Acknowledge that the message could not be processed
		} else {
			log.Printf("Account created in Payment service database.")
			msg.Ack() // Acknowledge successful processing
		}
	}, nats.Durable("payment_service_durable"))

	if err != nil {
		return err
	}

	log.Printf("Subscribed to account.created events with durable subscription: %s", subscription.Subject)
	return nil
}

func (ps *PaymentService) SubscribeToAccountTopUp() error {
	// Create a durable subscription to JetStream
	subscription, err := ps.js.Subscribe("account.topup", func(msg *nats.Msg) {
		var updatedAccount entity.Account
		if err := json.Unmarshal(msg.Data, &updatedAccount); err != nil {
			log.Printf("Error unmarshalling account data: %v", err)
			msg.Nak() // Acknowledge that the message could not be processed
			return
		}

		// Update the account in the Payment service's database
		if err := ps.db.Model(&entity.Account{}).Where("id = ?", updatedAccount.ID).Update("balance", updatedAccount.Balance).Error; err != nil {
			log.Printf("Error updating account in database: %v", err)
			msg.Nak() // Acknowledge that the message could not be processed
		} else {
			log.Printf("Account updated in Payment service database.")
			msg.Ack() // Acknowledge successful processing
		}
	}, nats.Durable("payment_service_update_durable"))

	if err != nil {
		return err
	}

	log.Printf("Subscribed to account.updated events with durable subscription: %s", subscription.Subject)
	return nil
}

func (ps *PaymentService) SubscribeToAccountDeleted() error {
	// Create a durable subscription to JetStream
	subscription, err := ps.js.Subscribe("account.deleted", func(msg *nats.Msg) {
		var deletedAccount entity.Account
		if err := json.Unmarshal(msg.Data, &deletedAccount); err != nil {
			log.Printf("Error unmarshalling account data: %v", err)
			msg.Nak() // Acknowledge that the message could not be processed
			return
		}

		// Delete the account from the Payment service's database
		if err := ps.db.Delete(&deletedAccount).Error; err != nil {
			log.Printf("Error deleting account from database: %v", err)
			msg.Nak() // Acknowledge that the message could not be processed
		} else {
			log.Printf("Account %s deleted from Payment service database.", deletedAccount.ID)
			msg.Ack() // Acknowledge successful processing
		}
	}, nats.Durable("payment_service_delete_durable"))

	if err != nil {
		return err
	}

	log.Printf("Subscribed to account.deleted events with durable subscription: %s", subscription.Subject)
	return nil
}

func (ps *PaymentService) SubscribeToAccountChangedPin() error {
	// Create a durable subscription to JetStream for account.changedPin event
	subscription, err := ps.js.Subscribe("account.changedPin", func(msg *nats.Msg) {
		var accountPinChange entity.Account
		if err := json.Unmarshal(msg.Data, &accountPinChange); err != nil {
			log.Printf("Error unmarshalling account PIN change data: %v", err)
			msg.Nak() // Acknowledge that the message could not be processed
			return
		}

		// Assuming account PIN change involves updating the account's PIN in the Payment service database
		if err := ps.db.Model(&entity.Account{}).Where("id = ?", accountPinChange.ID).Update("pin", accountPinChange.Pin).Error; err != nil {
			log.Printf("Error updating account PIN in database: %v", err)
			msg.Nak() // Acknowledge that the message could not be processed
		} else {
			log.Printf("Account PIN updated for account %s in Payment service database.", accountPinChange.ID)
			msg.Ack() // Acknowledge successful processing
		}
	}, nats.Durable("payment_service_changed_pin_durable"))

	if err != nil {
		return err
	}

	log.Printf("Subscribed to account.changedPin events with durable subscription: %s", subscription.Subject)
	return nil
}
