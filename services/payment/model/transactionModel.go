package model

type PaymentInfo struct {
	ID              string
	SourceAccountID string
	ReferenceCode   string
	Amount          float64
}

type CreatePayment struct {
	SourceAccountID string
	ReferenceCode   string
	Amount          float64
}
