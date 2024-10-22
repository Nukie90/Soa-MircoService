package model

type SignUp struct {
	ID        string `json:"id" example:"111"`
	Name      string `json:"name" example:"John Doe"`
	Address   string `json:"address" example:"Bangkok"`
	Password  string `json:"password" example:"password"`
	BirthDate string `json:"birth_date" example:"2004-01-02"`
}

type Login struct {
	ID       string `json:"id" example:"111"`
	Password string `json:"password" example:"password"`
}
