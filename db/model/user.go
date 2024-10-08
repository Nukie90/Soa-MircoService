package model

type Signup struct {
	Name      string `json:"name"`
	Address   string `json:"address"`
	Password  string `json:"password"`
	BirthDate string `json:"birth_date"`
}
