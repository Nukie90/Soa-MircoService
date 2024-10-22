package logic

import (
	"encoding/json"
	"log"
	"microservice/entity"

	"github.com/nats-io/nats.go"
	"gorm.io/gorm"
)

func (as *AccountService) SubscribeToUserCreated() error {
	// Create a durable subscription to JetStream
	subscription, err := as.js.Subscribe("user.created", func(msg *nats.Msg) {
		var newUser entity.User
		if err := json.Unmarshal(msg.Data, &newUser); err != nil {
			log.Printf("Error unmarshalling user data: %v", err)
			msg.Nak() // Acknowledge that the message could not be processed
			return
		}

		// Store the new user in the Account service's database
		if err := as.db.Create(&newUser).Error; err != nil {
			log.Printf("Error saving user to database: %v", err)
			msg.Nak() // Acknowledge that the message could not be processed
		} else {
			log.Printf("User %s created in Account service database.", newUser.Name)
			msg.Ack() // Acknowledge successful processing
		}
	}, nats.Durable("account_service_durable"))

	if err != nil {
		return err
	}

	log.Printf("Subscribed to user.created events with durable subscription: %s", subscription.Subject)
	return nil
}

func (as *AccountService) SubscribeToTransactionCreated() error {
	subscription, err := as.js.Subscribe("transaction.created", func(msg *nats.Msg) {
		var transaction entity.Transaction
		if err := json.Unmarshal(msg.Data, &transaction); err != nil {
			log.Printf("Error unmarshalling transaction data: %v", err)
			msg.Nak() // Acknowledge that the message could not be processed
			return
		}

		log.Printf("Processing transaction %s", transaction.ID)

		// Perform the transaction within a database transaction for safety
		err := as.db.Transaction(func(tx *gorm.DB) error {
			var sourceAccount, destinationAccount entity.Account

			// Fetch source and destination accounts
			if err := tx.Where("id = ?", transaction.SourceAccountID).First(&sourceAccount).Error; err != nil {
				return err
			}
			if err := tx.Where("id = ?", transaction.DestinationAccountID).First(&destinationAccount).Error; err != nil {
				return err
			}

			// Adjust balances
			sourceAccount.Balance -= transaction.Amount
			destinationAccount.Balance += transaction.Amount

			// Save updated accounts
			if err := tx.Save(&sourceAccount).Error; err != nil {
				return err
			}
			if err := tx.Save(&destinationAccount).Error; err != nil {
				return err
			}

			return nil
		})

		if err != nil {
			log.Printf("Error processing transaction %s: %v", transaction.ID, err)
			msg.Nak() // Acknowledge failure to process the message
			return
		}

		log.Printf("Transaction %s processed successfully", transaction.ID)
		msg.Ack() // Acknowledge successful processing
	}, nats.Durable("account_service_transaction_durable"))

	if err != nil {
		return err
	}

	log.Printf("Subscribed to transaction.created events with durable subscription: %s", subscription.Subject)
	return nil
}

func (as *AccountService) SubscribeToPaymentCreated() error {
	subscription, err := as.js.Subscribe("payment.created", func(msg *nats.Msg) {
		var payment entity.Payment
		if err := json.Unmarshal(msg.Data, &payment); err != nil {
			log.Printf("Error unmarshalling payment data: %v", err)
			msg.Nak() // Acknowledge that the message could not be processed
			return
		}

		log.Printf("Processing payment %s", payment.ID)

		// Perform the payment within a database payment for safety
		err := as.db.Transaction(func(tx *gorm.DB) error {
			var sourceAccount entity.Account

			// Fetch source account
			if err := tx.Where("id = ?", payment.SourceAccountID).First(&sourceAccount).Error; err != nil {
				return err
			}

			// Adjust balance
			sourceAccount.Balance -= payment.Amount

			// Save updated account
			if err := tx.Save(&sourceAccount).Error; err != nil {
				return err
			}

			return nil
		})

		if err != nil {
			log.Printf("Error processing payment %s: %v", payment.ID, err)
			msg.Nak() // Acknowledge failure to process the message
			return
		}

		log.Printf("Payment %s processed successfully", payment.ID)
		msg.Ack() // Acknowledge successful processing
	}, nats.Durable("account_service_payment_durable"))

	if err != nil {
		return err
	}

	log.Printf("Subscribed to payment.created events with durable subscription: %s", subscription.Subject)
	return nil
}
