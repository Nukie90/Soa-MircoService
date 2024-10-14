package model

type PaymentInfo struct {
	ID              string
	SourceAccountID string
	ReferenceCode   string
	Amount          float64
	CreateAt 	  string
}

type CreatePayment struct {
	SourceAccountID string
	ReferenceCode   string
	Amount          float64
}
