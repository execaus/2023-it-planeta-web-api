package models

type RegistrationAccountInput struct {
	FirstName string `json:"firstName,omitempty" binding:"required,excludesall=' ',printascii"`
	LastName  string `json:"lastName,omitempty" binding:"required,excludesall=' ',printascii"`
	Email     string `json:"email,omitempty" binding:"required,email,excludesall=' ',printascii"`
	Password  string `json:"password,omitempty" binding:"required,excludesall=' ',printascii"`
}

type RegistrationAccountOutput struct {
	Id        int64  `json:"id,omitempty"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	Email     string `json:"email,omitempty"`
}

type GetAccountOutput struct {
	Id        int64  `json:"id,omitempty"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	Email     string `json:"email,omitempty"`
}

type GetAccountsInput struct {
	FirstName string `form:"firstName"`
	LastName  string `form:"lastName"`
	Email     string `form:"email"`
	From      *int   `form:"from"`
	Size      *int   `form:"size"`
}

type GetAccountsOutput struct {
	Id        int64  `json:"id,omitempty"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	Email     string `json:"email,omitempty"`
}
