package logic

import (
	"encoding/json"
	"fmt"
	"log"
	"microservice/entity"

	"github.com/nats-io/nats.go"
	"gorm.io/gorm"
)

func (ts *TransactionService) SubscribeToUserCreated() error {
	// Create a durable subscription to JetStream
	subscription, err := ts.js.Subscribe("user.created", func(msg *nats.Msg) {
		var newUser entity.User
		if err := json.Unmarshal(msg.Data, &newUser); err != nil {
			log.Printf("Error unmarshalling user data: %v", err)
			msg.Nak() // Acknowledge that the message could not be processed
			return
		}

		// Store the new user in the Transaction service's database
		if err := ts.db.Create(&newUser).Error; err != nil {
			log.Printf("Error saving user to database: %v", err)
			msg.Nak() // Acknowledge that the message could not be processed
		} else {
			log.Printf("User %s created in Transaction service database.", newUser.Name)
			msg.Ack() // Acknowledge successful processing
		}
	}, nats.Durable("transaction_service_durable"))

	if err != nil {
		return err
	}

	log.Printf("Subscribed to user.created events with durable subscription: %s", subscription.Subject)
	return nil
}

func (ts *TransactionService) SubscribeToAccountCreated() error {
	// Create a durable subscription to JetStream
	subscription, err := ts.js.Subscribe("account.created", func(msg *nats.Msg) {
		var newAccount entity.Account
		if err := json.Unmarshal(msg.Data, &newAccount); err != nil {
			log.Printf("Error unmarshalling account data: %v", err)
			msg.Nak() // Acknowledge that the message could not be processed
			return
		}

		// Store the new account in the Transaction service's database
		if err := ts.db.Create(&newAccount).Error; err != nil {
			log.Printf("Error saving account to database: %v", err)
			msg.Nak() // Acknowledge that the message could not be processed
		} else {
			log.Printf("Account created in Transaction service database.")
			msg.Ack() // Acknowledge successful processing
		}
	}, nats.Durable("transaction_service_durable"))

	if err != nil {
		return err
	}

	log.Printf("Subscribed to account.created events with durable subscription: %s", subscription.Subject)
	return nil
}

func (ts *TransactionService) SubscribeToAccountTopUp() error {
	// Create a durable subscription to JetStream
	subscription, err := ts.js.Subscribe("account.topup", func(msg *nats.Msg) {
		var updatedAccount entity.Account
		if err := json.Unmarshal(msg.Data, &updatedAccount); err != nil {
			log.Printf("Error unmarshalling account data: %v", err)
			msg.Nak() // Acknowledge that the message could not be processed
			return
		}

		// Update the account in the Transaction service's database
		if err := ts.db.Model(&entity.Account{}).Where("id = ?", updatedAccount.ID).Update("balance", updatedAccount.Balance).Error; err != nil {
			log.Printf("Error updating account in database: %v", err)
			msg.Nak() // Acknowledge that the message could not be processed
		} else {
			log.Printf("Account updated in Transaction service database.")
			msg.Ack() // Acknowledge successful processing
		}
	}, nats.Durable("transaction_service_update_durable"))

	if err != nil {
		fmt.Println("Error")
		return err
	}

	log.Printf("Subscribed to account.updated events with durable subscription: %s", subscription.Subject)
	return nil
}

func (ts *TransactionService) SubscribeToAccountDeleted() error {
	// Create a durable subscription to JetStream
	subscription, err := ts.js.Subscribe("account.deleted", func(msg *nats.Msg) {
		var deletedAccount entity.Account
		if err := json.Unmarshal(msg.Data, &deletedAccount); err != nil {
			log.Printf("Error unmarshalling account data: %v", err)
			msg.Nak() // Acknowledge that the message could not be processed
			return
		}

		// Delete the account from the Transaction service's database
		if err := ts.db.Delete(&deletedAccount).Error; err != nil {
			log.Printf("Error deleting account from database: %v", err)
			msg.Nak() // Acknowledge that the message could not be processed
		} else {
			log.Printf("Account %d deleted from Transaction service database.", deletedAccount.ID)
			msg.Ack() // Acknowledge successful processing
		}
	}, nats.Durable("transaction_service_delete_durable"))

	if err != nil {
		return err
	}

	log.Printf("Subscribed to account.deleted events with durable subscription: %s", subscription.Subject)
	return nil
}

func (ts *TransactionService) SubscribeToAccountChangedPin() error {
	// Create a durable subscription to JetStream for account.changedPin event
	subscription, err := ts.js.Subscribe("account.changedPin", func(msg *nats.Msg) {
		var accountPinChange entity.Account
		if err := json.Unmarshal(msg.Data, &accountPinChange); err != nil {
			log.Printf("Error unmarshalling account PIN change data: %v", err)
			msg.Nak() // Acknowledge that the message could not be processed
			return
		}

		// Assuming account PIN change involves updating the account's PIN in the Transaction service database
		if err := ts.db.Model(&entity.Account{}).Where("id = ?", accountPinChange.ID).Update("pin", accountPinChange.Pin).Error; err != nil {
			log.Printf("Error updating account PIN in database: %v", err)
			msg.Nak() // Acknowledge that the message could not be processed
		} else {
			log.Printf("Account PIN updated for account %s in Transaction service database.", accountPinChange.ID)
			msg.Ack() // Acknowledge successful processing
		}
	}, nats.Durable("transaction_service_changed_pin_durable"))

	if err != nil {
		return err
	}

	log.Printf("Subscribed to account.changedPin events with durable subscription: %s", subscription.Subject)
	return nil
}

func (ts *TransactionService) SubscribeToPaymentCreated() error {
	subscription, err := ts.js.Subscribe("payment.created", func(msg *nats.Msg) {
		var payment entity.Payment
		if err := json.Unmarshal(msg.Data, &payment); err != nil {
			log.Printf("Error unmarshalling payment data: %v", err)
			msg.Nak() // Acknowledge that the message could not be processed
			return
		}

		log.Printf("Processing payment %s", payment.ID)

		// Perform the payment within a database payment for safety
		err := ts.db.Transaction(func(tx *gorm.DB) error {
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
	}, nats.Durable("transaction_service_payment_durable"))

	if err != nil {
		return err
	}

	log.Printf("Subscribed to payment.created events with durable subscription: %s", subscription.Subject)
	return nil
}
