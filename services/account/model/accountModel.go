package model

type AccountInfo struct {
	ID      string
	UserID  string
	Type    string
	Balance float64
	Pin     string
}

type CreateAccount struct {
	Type    string  `json:"type"`
	Balance float64 `json:"balance"`
	Pin     string  `json:"pin"`
}

type AccountVerify struct {
	ID  string `json:"id"`
	Pin string `json:"pin"`
}

type AccountBalance struct {
	ID      string  `json:"id"`
	Balance float64 `json:"balance"`
	Type    string  `json:"type"`
}

type ChangePin struct {
	ID     string `json:"id"`
	OldPin string `json:"oldPin"`
	NewPin string `json:"newPin"`
}

type DeleteAccount struct {
	ID  string `json:"id"`
	Pin string `json:"pin"`
}

type TopUp struct {
	ID     string  `json:"id"`
	Amount float64 `json:"amount"`
}
