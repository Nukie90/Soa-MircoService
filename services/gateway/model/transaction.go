package model

type TransactionInfo struct {
	ID                   string
	SourceAccountID      string
	DestinationAccountID string
	Amount               float64
}

type CreateTransaction struct {
	SourceAccountID      string
	DestinationAccountID string
	Amount               float64
}
