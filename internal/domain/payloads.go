package domain

type RegisterUserPayload struct {
	Username  string `json:"username,omitempty" validate:"required,min=4,max=20"`
	Firstname string `json:"firstname,omitempty" validate:"required,min=3,max=20"`
	Lastname  string `json:"lastname,omitempty" validate:"required,min=3,max=20"`
	Email     string `json:"email,omitempty" validate:"required,email,max=55"`
	Phone     string `json:"phone,omitempty" validate:"required,min=9,max=15"`
	Password  string `json:"password,omitempty" validate:"required,min=8,max=25"`
	Age       *uint8 `json:"age,omitempty" validate:"omitempty,gte=14,lte=100"`
}
