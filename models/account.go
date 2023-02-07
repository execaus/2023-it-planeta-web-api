package models

type RegistrationAccountInput struct {
	FirstName string `json:"firstName,omitempty" binding:"required,no_whitespace_only"`
	LastName  string `json:"lastName,omitempty" binding:"required,no_whitespace_only"`
	Email     string `json:"email,omitempty" binding:"required,email,no_whitespace_only"`
	Password  string `json:"password,omitempty" binding:"required,no_whitespace_only"`
}

type RegistrationAccountOutput struct {
	Id        int32  `json:"id,omitempty"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	Email     string `json:"email,omitempty"`
}
